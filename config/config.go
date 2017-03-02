// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period time.Duration			  time.Duration     `config:"period"`
	Port                 			  int               `config:"port"`
	Addr                 			  string
	EnableJsonValidation              bool              `config:"enable_json_validation"`
	PublishFailedJsonSchemaValidation bool              `config:"publish_failed_json_schema_validation"`
	PublishFailedJsonInvalid          bool              `config:"publish_failed_json_invalid"`
	JsonDocumentTypeSchema            map[string]string `config:"json_document_type_schema"`
}

var DefaultConfig = Config{
	Period:							   1 * time.Second,
	Port:							   24224,
	EnableJsonValidation:              false,
	PublishFailedJsonSchemaValidation: false,
	PublishFailedJsonInvalid:          false,
}
