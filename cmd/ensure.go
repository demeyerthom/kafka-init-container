package main

import (
	"bytes"
	"fmt"
	"github.com/demeyerthom/kafka-init-container/pkg/configurator"
	"github.com/demeyerthom/kafka-init-container/pkg/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
)

const (
	KafkaDir = "configuration"
)

func main() {

	files, err := ioutil.ReadDir(KafkaDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		configuration := models.Settings{}

		yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", KafkaDir, f.Name()))
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(yamlFile, &configuration)
		if err != nil {
			log.Fatal(err)
		}

		pingZookeeper(configuration.Zookeeper)

		topicConfigurator := configurator.NewTopicConfigurator(configuration.Zookeeper)
		updateTopics(topicConfigurator, configuration.Topics)
	}
}

func pingZookeeper(zookeeperAddr string) {
	var continuePing = true

	for continuePing {
		_, err := net.Dial("tcp", zookeeperAddr)

		if err == nil {
			log.Print("Online")
			continuePing = false
			continue
		}

		log.Print("Connection error:", err)
	}
}

func updateTopics(configurator configurator.TopicConfigurator, topics []models.Topic) {
	for _, topic := range topics {
		cmd := configurator.CreateTopicCommand(topic)
		cmdOutput := &bytes.Buffer{}
		cmd.Stdout = cmdOutput
		err := cmd.Run()
		if err != nil {
			log.Fatalf("An error occured while updating: %s %s", err, string(cmdOutput.Bytes()))
		}

		log.Print(string(cmdOutput.Bytes()))
		log.Print(topic)
	}
}
