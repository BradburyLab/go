package log

type BackendLogrus struct{}

func (it *BackendLogrus) ProduceLogger(config *ConfigLogger) Logger {
	return nil
}
