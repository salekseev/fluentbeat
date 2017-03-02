package beater

import (
	"fmt"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/fluent/fluentd-forwarder"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/xeipuuv/gojsonschema"

	"github.com/salekseev/fluentbeat/config"
)

type Fluentbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
	jsonDocumentSchema	map[string]gojsonschema.JSONLoader
	input	*fluentd_forwarder.ForwardInput
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Fluentbeat{
		done:   make(chan struct{}),
		config: config,
	}

	if bt.config.EnableJsonValidation {

		bt.jsonDocumentSchema = map[string]gojsonschema.JSONLoader{}

		for name, path := range config.JSONDocumentTypeSchema {
			logp.Info("Loading JSON schema %s from %s", name, path)
			schemaLoader := gojsonschema.NewReferenceLoader("file://" + path)
			ds := schemaLoader
			bt.jsonDocumentSchema[name] = ds
		}

	}

	bt.config.Addr = fmt.Sprintf("127.0.0.1:%d", bt.config.Port)

	return bt, nil
}

func (bt *Fluentbeat) Run(b *beat.Beat) error {
	logp.Info("fluentbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Fluentbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
