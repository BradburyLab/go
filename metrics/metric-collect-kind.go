package metrics

type CollectKind uint8

const (
	COLLECT_KIND_SYNC CollectKind = iota
	COLLECT_KIND_ASYNC
)
