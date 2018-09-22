package main

import (
	"fmt"
	"github.com/alecthomas/kingpin"
	"github.com/demeyerthom/kafka-init-container/pkg/models"
	"github.com/demeyerthom/kafka-init-container/pkg/topic"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	kafkaDir = kingpin.Flag("configuration-dir", "the directory containing all the configuration files").
		Default("configuration").Envar("CONFIGURATION_DIR").String()
)

func init() {
	kingpin.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	files, err := ioutil.ReadDir(*kafkaDir)
	if err != nil {
		log.WithError(err).Fatalf("An error occurred while reading files: %s", err)
	}

	for _, f := range files {
		configuration := models.Settings{}

		file := fmt.Sprintf("%s/%s", *kafkaDir, f.Name())
		yamlFile, err := ioutil.ReadFile(file)
		if err != nil {
			log.WithError(err).WithField("fileName", file).Fatalf("An error occurred while reading file %s: %s", file, err)
		}

		err = yaml.Unmarshal(yamlFile, &configuration)
		if err != nil {
			log.WithError(err).WithField("fileName", file).Fatalf("An error occurred while unmarshalling file %s: %s", file, err)
		}

		lister := topic.NewTopicLister(configuration.Zookeeper)

		list := lister.GetTopicList()

		topicCreator := topic.NewTopicCreator(configuration.Zookeeper, true)

		for _, configTopic := range configuration.Topics {
			if contains(list, configTopic.Name) {
				log.Warnf("Topic %s should be updated; not implemented yet", configTopic.Name)
			} else {
				topicCreator.CreateTopic(configTopic)
			}
		}
	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
