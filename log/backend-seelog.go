package log

import (
	"strings"

	"github.com/cihub/seelog"
)

const (
	FORMAT_FILE     string = `[%Date(2/Jan/2006 15:04:05)] {{.Prefix}} [%l] %Msg%n`
	FORMAT_TRACE    string = `%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(36){{.Prefix}}%EscM(0)%EscM(36;1) %EscM(36;1)[%l]%EscM(0)%EscM(36) %Msg%n%EscM(0)`
	FORMAT_DEBUG    string = `%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(36){{.Prefix}}%EscM(0)%EscM(36;1) %EscM(34)[%l]%EscM(0)%EscM(34;1) %Msg%n%EscM(0)`
	FORMAT_INFO     string = `%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(36){{.Prefix}}%EscM(0)%EscM(36;1) %EscM(32;1)[%l]%EscM(0)%EscM(32) %Msg%n%EscM(0)`
	FORMAT_WARN     string = `%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(36){{.Prefix}}%EscM(0)%EscM(36;1) %EscM(33;1)[%l]%EscM(0)%EscM(33) %Msg%n%EscM(0)`
	FORMAT_ERROR    string = `%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(36){{.Prefix}}%EscM(0)%EscM(36;1) %EscM(35;1)[%l]%EscM(0)%EscM(35) %Msg%n%EscM(0)`
	FORMAT_CRITICAL string = `%EscM(0)[%Date(2/Jan/2006 15:04:05)] %EscM(36){{.Prefix}}%EscM(0)%EscM(36;1) %EscM(31;1)[%l]%EscM(0)%EscM(31) %Msg%n%EscM(0)`
)

type BackendSeelog struct{}

func (it *BackendSeelog) ProduceLogger(config *ConfigLogger) Logger {
	consoleWriter, e := seelog.NewConsoleWriter()
	if e != nil {
		panic(e)
	}

	// logrotate writer open here
	fileAccessWriter, e := seelog.NewFileWriter("./access.log")
	if e != nil {
		panic(e)
	}

	// logrotate writer open here
	fileErrorWriter, e := seelog.NewFileWriter("./error.log")
	if e != nil {
		panic(e)
	}

	formatFile := strings.Replace(FORMAT_FILE, "{{.Prefix}}", config.Prefix, -1)
	formatterFile, e := seelog.NewFormatter(formatFile)
	if e != nil {
		panic(e)
	}

	formatTrace := strings.Replace(FORMAT_TRACE, "{{.Prefix}}", config.Prefix, -1)
	formatterTrace, e := seelog.NewFormatter(formatTrace)
	if e != nil {
		panic(e)
	}

	formatDebug := strings.Replace(FORMAT_DEBUG, "{{.Prefix}}", config.Prefix, -1)
	formatterDebug, e := seelog.NewFormatter(formatDebug)
	if e != nil {
		panic(e)
	}

	formatInfo := strings.Replace(FORMAT_INFO, "{{.Prefix}}", config.Prefix, -1)
	formatterInfo, e := seelog.NewFormatter(formatInfo)
	if e != nil {
		panic(e)
	}

	formatWarn := strings.Replace(FORMAT_WARN, "{{.Prefix}}", config.Prefix, -1)
	formatterWarn, e := seelog.NewFormatter(formatWarn)
	if e != nil {
		panic(e)
	}

	formatError := strings.Replace(FORMAT_ERROR, "{{.Prefix}}", config.Prefix, -1)
	formatterError, e := seelog.NewFormatter(formatError)
	if e != nil {
		panic(e)
	}

	formatCritical := strings.Replace(FORMAT_CRITICAL, "{{.Prefix}}", config.Prefix, -1)
	formatterCritical, e := seelog.NewFormatter(formatCritical)
	if e != nil {
		panic(e)
	}

	constraints, _ := seelog.NewMinMaxConstraints(seelog.TraceLvl, seelog.CriticalLvl)

	dispatcherFileAccess, _ := seelog.NewFilterDispatcher(formatterFile, []interface{}{fileAccessWriter},
		seelog.TraceLvl, seelog.DebugLvl, seelog.InfoLvl,
		seelog.WarnLvl, seelog.ErrorLvl, seelog.CriticalLvl,
	)
	dispatcherFileError, _ := seelog.NewFilterDispatcher(formatterFile, []interface{}{fileErrorWriter}, seelog.WarnLvl, seelog.ErrorLvl, seelog.CriticalLvl)

	dispatcherTrace, _ := seelog.NewFilterDispatcher(formatterTrace, []interface{}{consoleWriter}, seelog.TraceLvl)
	dispatcherDebug, _ := seelog.NewFilterDispatcher(formatterDebug, []interface{}{consoleWriter}, seelog.DebugLvl)
	dispatcherInfo, _ := seelog.NewFilterDispatcher(formatterInfo, []interface{}{consoleWriter}, seelog.InfoLvl)
	dispatcherWarn, _ := seelog.NewFilterDispatcher(formatterWarn, []interface{}{consoleWriter}, seelog.WarnLvl)
	dispatcherError, _ := seelog.NewFilterDispatcher(formatterError, []interface{}{consoleWriter}, seelog.ErrorLvl)
	dispatcherCritical, _ := seelog.NewFilterDispatcher(formatterCritical, []interface{}{consoleWriter}, seelog.CriticalLvl)

	root, _ := seelog.NewSplitDispatcher(formatterTrace,
		[]interface{}{
			dispatcherFileAccess,
			dispatcherFileError,
			dispatcherTrace,
			dispatcherDebug,
			dispatcherInfo,
			dispatcherWarn,
			dispatcherError,
			dispatcherCritical,
		},
	)

	logger := seelog.NewAsyncLoopLogger(seelog.NewLoggerConfig(constraints, nil, root))
	return logger
}
