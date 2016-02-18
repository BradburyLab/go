package metrics

type Metric interface {
	Collect() chan Result
	Len() int
}

type Result interface {
	V() interface{}
	Kind() Kind
	Name() string

	OK() bool
	Err() error
}
