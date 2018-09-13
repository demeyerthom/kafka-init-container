package models

type Topic struct {
	Name              string        `yaml:"name"`
	ReplicationFactor int           `yaml:"replication-factor"`
	Partitions        int           `yaml:"partitions"`
	Configuration     Configuration `yaml:"config"`
}

type Configuration struct {
	CleanupPolicy   string `yaml:"cleanup.policy"`
	CompressionType string `yaml:"compression.type"`
	RetentionMs     int    `yaml:"retention.ms"`
}
