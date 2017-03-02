package beater

import "github.com/elastic/beats/libbeat/logp"

type FluentdLogger struct{}

func (FluentdLogger) Print(v ...interface{}) {
	logp.Warn("Fluentd message: %v", v...)
}

func (FluentdLogger) Printf(format string, v ...interface{}) {
	logp.Warn(format, v...)
}

func (FluentdLogger) Println(v ...interface{}) {
	logp.Warn("Fluentd message: %v", v...)
}
