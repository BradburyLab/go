package log

type BackendKind uint8

const (
	BACKEND_KIND_SEELOG BackendKind = iota
	BACKEND_KIND_VANILA
	BACKEND_KIND_LOGRUS
)

type Backend interface {
	ProduceLogger(*ConfigLogger) Logger
}
