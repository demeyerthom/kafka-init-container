# Kafka init container

A container to initialize kafka topics on an arbitrary cluster

## Config file example
```yaml
---
brokers:
- localhost:2181

topics:
- name: some-topic
  replication-factor: 1
  partitions: 4
  config:
    cleanup.policy: compact
    compression.type: producer
    retention.ms: -1
```

Other topic configs can be provided by adding them under `config`. See
[the documentation](http://kafka.apache.org/documentation/#topicconfigs)
for all possible options. Multiple configuration files can be provided
in the same folder.

## Usage

Load the container and mount the configuration files on the container.
See below for all the options:

| Environment variable | Description |
| :-------------: | :-------------: |
| `CONFIGURATION_DIR` | The configuration directory to read from. Defaults to `./configuration` |

## Todo
- Tests (because I hate writing them)
- Have a nice docker container somewhere
- Updating topics with new configurations when files have changed
- Rollback changes on error
- Set/Reset consumer positions (handy for testing etc?)