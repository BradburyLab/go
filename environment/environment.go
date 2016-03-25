package l

import (
	"fmt"
	"strings"
)

type Environment uint8

const (
	ENV_UNKNOWN Environment = 0

	ENV_DEV Environment = iota
	ENV_TEST
	ENV_DEMO
	ENV_PROD
)

var environmentText = map[Environment]string{
	ENV_UNKNOWN: "unknown",
	ENV_DEV:     "development",
	ENV_TEST:    "test",
	ENV_DEMO:    "demo",
	ENV_PROD:    "production",
}

func EnvironmentValidList() (validList []string) {
	for m, name := range environmentText {
		if m.IsUnknown() {
			continue
		}
		validList = append(validList, name)
	}

	return
}

func EnvironmentValidListAsString() string { return strings.Join(EnvironmentValidList(), " | ") }

func NewEnvironment(v string) Environment {
	v = strings.ToLower(v)
	v = strings.Replace(v, " ", "", -1)
	v = strings.Replace(v, "_", "-", -1)

	switch v {
	case "d", "dev", "development":
		return ENV_DEV
	case "t", "tst", "test":
		return ENV_TEST
	case "demo":
		return ENV_DEMO
	case "p", "prod", "production":
		return ENV_PROD
	}

	return ENV_UNKNOWN
}

func (self Environment) Is(v Environment) bool { return self == v }

func (self Environment) IsDev() bool     { return self.Is(ENV_DEV) }
func (self Environment) IsTest() bool    { return self.Is(ENV_TEST) }
func (self Environment) IsNotTest() bool { return !self.Is(ENV_TEST) }
func (self Environment) IsDemo() bool    { return self.Is(ENV_DEMO) }
func (self Environment) IsProd() bool    { return self.Is(ENV_PROD) }
func (self Environment) IsUnknown() bool { return self.Is(ENV_UNKNOWN) }
func (self Environment) IsValid() bool   { return !self.IsUnknown() }

func (self Environment) String() string                  { return environmentText[self] }
func (self Environment) ValidList() (validList []string) { return EnvironmentValidList() }
func (self Environment) ValidListAsString() string       { return EnvironmentValidListAsString() }
func (self Environment) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", self.String())), nil
}

func (self *Environment) UnmarshalJSON(data []byte) error {
	v := strings.Replace(string(data), "\"", "", -1)
	*self = NewEnvironment(v)
	return nil
}

func (self *Environment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	v := ""
	unmarshal(&v)
	*self = NewEnvironment(v)
	if self.IsUnknown() {
		return fmt.Errorf("got unknown environment '%s', possible environments are number of: %s;", v, EnvironmentValidListAsString())
	}
	return nil
}
