package log

type Config struct {
	Path    string `yaml:"path"`
	Loggers map[string]*ConfigLogger

	BackendKind BackendKind `yaml:"-"`
	backend     Backend
}

type ConfigLogger struct {
	Prefix string `yaml:"prefix"`
	Path   struct {
		Access string `yaml:"access"`
		Error  string `yaml:"error"`
	} `yaml:"path"`
	Files struct {
		Access string `yaml:"access"`
		Error  string `yaml:"error"`
	} `yaml:"files"`
	Level struct {
		Min LogLevel `yaml:"min"`
		Max LogLevel `yaml:"max"`
	} `yaml:"level"`
	Format struct {
		File    string `yaml:"file"`
		Console string `yaml:"console"`
	} `yaml:"format"`
}

func (it *Config) RegisterBackend(v Backend)                 { it.backend = v }
func (it *Config) ProduceLogger(config *ConfigLogger) Logger { return it.backend.ProduceLogger(config) }
