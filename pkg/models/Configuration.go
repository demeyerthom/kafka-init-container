package models

type Settings struct {
	Zookeeper string  `yaml:"zookeeper"`
	Topics    []Topic `yaml:"topics"`
}
