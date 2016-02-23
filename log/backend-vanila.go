package log

type BackendVanila struct{}

func (it *BackendVanila) ProduceLogger(config *ConfigLogger) Logger {
	return nil
}
