package configurator

import (
	"github.com/demeyerthom/kafka-init-container/pkg/models"
	"log"
	"os/exec"
)

type TopicConfigurator struct {
	zookeeper string
}

func NewTopicConfigurator(zookeeper string) TopicConfigurator {
	return TopicConfigurator{zookeeper: zookeeper}
}

func (c *TopicConfigurator) CreateTopicCommand(topic models.Topic) *exec.Cmd {

	args := c.createTopicsArgs(topic)

	cmd := exec.Command("/opt/kafka/bin/kafka-topics.sh", args...)

	return cmd
}

func (c *TopicConfigurator) createTopicsArgs(topic models.Topic) []string {
	var args []string

	if topic.Name == "" {
		log.Fatal("No name provided")
	}
	args = append(args, CreateFlag("topic", topic.Name))
	args = append(args, CreateFlag("zookeeper", c.zookeeper))

	if topic.Partitions == 0 {
		topic.Partitions = 1
	}
	args = append(args, CreateFlag("partitions", topic.Partitions))

	if topic.ReplicationFactor == 0 {
		topic.Partitions = 1
	}
	args = append(args, CreateFlag("replication-factor", topic.ReplicationFactor))

	if topic.Configuration.CleanupPolicy == "" {
		topic.Configuration.CleanupPolicy = "delete"
	}
	args = append(args, CreateConfigFlag("cleanup-policy", topic.Configuration.CleanupPolicy))

	return args
}
