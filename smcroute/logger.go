/* https://github.com/cihub/seelog/wiki/Writing-libraries-with-Seelog */

package smcroute

import (
	"fmt"
	"io"
	"sync"

	seelog "github.com/cihub/seelog"
)

var (
	loggerMutex sync.RWMutex
	logger      seelog.LoggerInterface
)

func init() {
	// Disable logger by default.
	DisableLog()
}

func log() seelog.LoggerInterface {
	loggerMutex.RLock()
	defer loggerMutex.RUnlock()

	return logger
}

// DisableLog disables all library log output.
func DisableLog() { UseLogger(seelog.Disabled) }

// UseLogger uses a specified seelog.LoggerInterface to output library log.
// Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()
	logger = newLogger
}

// SetLogWriter uses a specified io.Writer to output library log.
// Use this func if you are not using Seelog logging system in your app.
func SetLogWriter(writer io.Writer) error {
	if writer == nil {
		return fmt.Errorf("provided writer is nil")
	}

	newLogger, err := seelog.LoggerFromWriterWithMinLevel(writer, seelog.TraceLvl)
	if err != nil {
		return err
	}

	UseLogger(newLogger)
	return nil
}

// Call this before app shutdown
func FlushLog() {
	logger.Flush()
}
