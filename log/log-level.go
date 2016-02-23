package log

import "strings"

type LogLevel uint8

const (
	LOG_LEVEL_TRACE LogLevel = iota
	LOG_LEVEL_DEBUG
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_CRITICAL
	LOG_LEVEL_OFF
)

var logLevelText = map[LogLevel]string{
	LOG_LEVEL_TRACE:    "TRACE",
	LOG_LEVEL_DEBUG:    "DEBUG",
	LOG_LEVEL_INFO:     "INFO",
	LOG_LEVEL_WARN:     "WARNING",
	LOG_LEVEL_ERROR:    "ERROR",
	LOG_LEVEL_CRITICAL: "CRITICAL",
	LOG_LEVEL_OFF:      "OFF",
}

func (it LogLevel) String() string {
	if s, ok := logLevelText[it]; ok {
		return s
	}

	return "UNKNOWN"
}

func (it *LogLevel) UnmarshalJSON(data []byte) error {
	v := strings.Replace(string(data), "\"", "", -1)
	v = strings.TrimSpace(v)
	*it = NewLogLevelFromString(v)
	return nil
}

func (it *LogLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	v := ""
	unmarshal(&v)
	*it = NewLogLevelFromString(v)
	return nil
}

func NewLogLevelFromString(v string) LogLevel {
	v = strings.TrimSpace(v)
	v = strings.ToLower(v)

	switch v {
	case "t", "trace":
		return LOG_LEVEL_TRACE
	case "d", "debug", "dbg":
		return LOG_LEVEL_DEBUG
	case "i", "info":
		return LOG_LEVEL_INFO
	case "w", "warn", "warning":
		return LOG_LEVEL_WARN
	case "e", "error":
		return LOG_LEVEL_ERROR
	case "c", "critical", "crit", "panic":
		return LOG_LEVEL_CRITICAL
	case "o", "off", "no":
		return LOG_LEVEL_CRITICAL
	}

	return LOG_LEVEL_INFO
}
