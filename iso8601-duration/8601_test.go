package iso8601duration

import (
	"fmt"
	"testing"
	"time"
)

var fixtures = []struct {
	in       time.Duration
	expected string
}{
	{1 * time.Millisecond, "PT0.1S"},
	{1*time.Minute + 6*time.Second + 762*time.Millisecond, "PT1M6.762S"},
	{20 * time.Second, "PT20S"},
	{3*time.Minute + 15*time.Second, "PT3M15S"},
	{4*time.Hour + 54*time.Minute + 18*time.Second, "PT4H54M18S"},
	{6*24*time.Hour + 4*time.Hour + 54*time.Minute + 18*time.Second, "P6DT4H54M18S"},
	{12*7*24*time.Hour + 6*24*time.Hour + 4*time.Hour + 54*time.Minute + 43*time.Second, "P12W6DT4H54M43S"},
	{2*365*24*time.Hour + 12*7*24*time.Hour + 6*24*time.Hour + 4*time.Hour + 54*time.Minute + 18*time.Second, "P2Y12W6DT4H54M18S"},

	{6 * 24 * time.Hour, "P6D"},
	{2 * 365 * 24 * time.Hour, "P2Y"},
	{3 * time.Minute, "PT3M"},
}

func TestDuration(t *testing.T) {
	for _, f := range fixtures {
		actual := String(f.in)
		if actual != f.expected {
			t.Errorf(`{"d": %s, "expected": "%s", "actual": "%s"}`,
				f.in, f.expected, actual)
		} else {
			fmt.Printf(`{"d": %s, "expected": "%s", "actual": "%s"}`+"\n",
				f.in, f.expected, actual)
		}
	}
}
