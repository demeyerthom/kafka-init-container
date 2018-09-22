package topic

import (
	"bytes"
	"github.com/demeyerthom/kafka-init-container/pkg"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

type Lister struct {
	*Handler
}

func NewTopicLister(zookeeper []string) Lister {
	handler := Handler{zookeeper: zookeeper}

	return Lister{&handler}
}

func (l *Lister) GetTopicList() (list []string) {
	var args []string
	args = append(args, "/opt/kafka/bin/kafka-topics.sh")
	args = append(args, pkg.CreateFlag("list"))
	for _, broker := range l.zookeeper {
		args = append(args, pkg.CreateFlag("zookeeper"))
		args = append(args, broker)
	}

	cmd := exec.Command("sh", args...)

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	log.Debugf("Executing list command")
	err := cmd.Run()
	if err != nil {
		log.WithError(err).Fatalf("An error occurred while listing topics: %s", errb.String())
	}

	list = strings.Split(outb.String(), "\n")

	log.WithField("Stdout", outb.String()).Debugf("Found %d topics", len(list))

	return list
}
