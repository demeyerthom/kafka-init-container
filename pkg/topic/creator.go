package topic

import (
	"bytes"
	"github.com/demeyerthom/kafka-init-container/pkg"
	"github.com/demeyerthom/kafka-init-container/pkg/models"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type Creator struct {
	*Handler
	strict bool
}

func NewTopicCreator(zookeeper []string, strict bool) Creator {
	handler := Handler{zookeeper: zookeeper}

	return Creator{&handler, strict}
}

func (c *Creator) generateCommonTopicArgs(topic models.Topic) (args []string) {
	args = append(args, "/opt/kafka/bin/kafka-topics.sh")

	if topic.Name == "" {
		log.Fatal("No name provided")
	}
	args = append(args, pkg.CreateFlag("topic"))
	args = append(args, topic.Name)

	for _, broker := range c.zookeeper {
		args = append(args, pkg.CreateFlag("zookeeper"))
		args = append(args, broker)
	}

	if topic.Partitions == "" {
		topic.Partitions = "1"
	}
	args = append(args, pkg.CreateFlag("partitions"))
	args = append(args, topic.Partitions)

	if topic.ReplicationFactor == "" {
		topic.Partitions = "1"
	}
	args = append(args, pkg.CreateFlag("replication-factor"))
	args = append(args, topic.ReplicationFactor)

	for key, value := range topic.Configuration {
		args = append(args, pkg.CreateFlag("config"))
		args = append(args, pkg.CreateConfig(key, value))
	}

	if false == c.strict {
		args = append(args, pkg.CreateFlag("if-not-exists"))
	}

	return args
}

func (c *Creator) CreateTopic(topic models.Topic) {
	args := c.generateCommonTopicArgs(topic)
	args = append(args, pkg.CreateFlag("create"))
	cmd := exec.Command("sh", args...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	log.Debugf("Executing create command for %s", topic.Name)
	err := cmd.Run()
	if err != nil {
		log.WithError(err).Fatalf("An error occurred while configuring topics: %s", errb.String())
	}

	log.WithField("topic", topic).WithField("Stdout", outb.String()).Infof("Successfully created topic %s", topic.Name)
}
