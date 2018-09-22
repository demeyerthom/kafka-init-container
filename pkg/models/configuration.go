package models

type Settings struct {
	Zookeeper []string `yaml:"brokers"`
	Topics    []Topic  `yaml:"topics"`
}

type Topic struct {
	Name              string      `yaml:"name"`
	ReplicationFactor string      `yaml:"replication-factor"`
	Partitions        string      `yaml:"partitions"`
	Configuration     TopicConfig `yaml:"config"`
}

type TopicConfig map[string]string
