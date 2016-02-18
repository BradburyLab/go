package metrics

import "sync"

type Metrics struct {
	metrics []Metric
}

func (it *Metrics) Len() (result int) {
	for _, metric := range it.metrics {
		result += metric.Len()
	}
	return
}

func (it *Metrics) Register(metrics ...Metric) *Metrics {
	for _, metric := range metrics {
		it.Append(metric)
	}
	return it
}

func (it *Metrics) Append(metric Metric) *Metrics {
	it.metrics = append(it.metrics, metric)
	return it
}

func (it *Metrics) collect(kind CollectKind) (out chan chan Result) {
	out = make(chan chan Result, len(it.metrics))

	wg := new(sync.WaitGroup)
	wg.Add(cap(out))

	for _, metric := range it.metrics {
		go func(metric Metric) {
			out <- metric.Collect()
			wg.Done()
		}(metric)
	}

	switch kind {
	case COLLECT_KIND_ASYNC:
		go func() { wg.Wait(); close(out) }()
	case COLLECT_KIND_SYNC:
		wg.Wait()
		close(out)
	}

	return out
}

func (it *Metrics) CollectAsync() (out chan chan Result) { return it.collect(COLLECT_KIND_ASYNC) }
func (it *Metrics) Collect() (out chan chan Result)      { return it.collect(COLLECT_KIND_SYNC) }

func NewMetrics() *Metrics { return new(Metrics) }
func New() *Metrics        { return NewMetrics() }
