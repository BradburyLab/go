package main

import (
	"github.com/BradburyLab/go/log"
	"gopkg.in/yaml.v2"
)

const configData string = `
path: "."
loggers:
  app:
    prefix: "[---]"
    files: { access: access.log, error: error.log }
    level: { min: trace, max: critical }
  api:
    prefix: "[API]"
    files: { access: access.log, error: error.log }
    level: { min: trace, max: critical }
  fs:
    prefix: "[FS]"
    files: { access: access.log, error: error.log }
    level: { min: trace, max: critical }
  ws:
    prefix: "[WS]"
    files: { access: access.log, error: error.log }
    level: { min: trace, max: critical }
  billing:
    prefix: "[BILLING]"
    files: { access: access.log, error: error.log }
    level: { min: trace, max: critical }
  monitoring:
    prefix: "[MON]"
    files: { access: access.log, error: zabbix.log }
    level: { min: trace, max: critical }
    format:
      file: "%%Date(2006-Jan-2 15:04:05) %%Msg%%n"`

func main() {
	config := log.Config{}
	config.RegisterBackend(&log.BackendSeelog{})

	err := yaml.Unmarshal([]byte(configData), &config)
	if err != nil {
		panic(err)
	}

	for name, configLogger := range config.Loggers {
		logger := config.ProduceLogger(configLogger)

		logger.Tracef("%s ~ trace", name)
		logger.Debugf("%s ~ debug", name)
		logger.Infof("%s ~ info", name)
		logger.Warnf("%s ~ warn", name)
		logger.Errorf("%s ~ error", name)
		logger.Criticalf("%s ~ critical", name)
		logger.Flush()
	}
}
