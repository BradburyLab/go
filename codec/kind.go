package codec

type Kind uint8

const (
	UNKNOWN Kind = iota
	JSON
	HTML
	CSV
	TSV
	XML
	YAML
	MSG
	PROTO
	M3U8
)

const (
	CONTENT_TYPE_UNKNOWN string = "application/x-unknown"
	CONTENT_TYPE_JSON    string = "application/json"
	CONTENT_TYPE_HTML    string = "text/html"
	CONTENT_TYPE_CSV     string = "text/csv"
	CONTENT_TYPE_TSV     string = "text/tsv"
	CONTENT_TYPE_XML     string = "text/xml"
	CONTENT_TYPE_YAML    string = "application/x-yaml"
	CONTENT_TYPE_MSG     string = "application/x-msgpack"
	CONTENT_TYPE_PROTO   string = "application/x-protobuf"

	// See section 3.1 of http://tools.ietf.org/html/draft-pantos-http-live-streaming-08#section-3.1
	CONTENT_TYPE_M3U8 string = "application/vnd.apple.mpegurl"
	// https://developer.apple.com/library/ios/documentation/NetworkingInternet/Conceptual/StreamingMediaGuide/DeployingHTTPLiveStreaming/DeployingHTTPLiveStreaming.html
	CONTENT_TYPE_M3U8_1 string = "application/x-mpegURL"
)

var kindText1 = map[Kind]string{
	UNKNOWN: "UNKNOWN",
	JSON:    "JSON",
	HTML:    "HTML",
	CSV:     "CSV",
	TSV:     "TSV",
	XML:     "XML",
	YAML:    "YAML",
	MSG:     "MSG",
	PROTO:   "ProtocolBuffers",
	M3U8:    "M3U8",
}

var kindText2 = map[Kind]string{
	UNKNOWN: "unk",
	JSON:    "json",
	HTML:    "html",
	CSV:     "csv",
	TSV:     "tsv",
	XML:     "xml",
	YAML:    "yml",
	MSG:     "msg",
	PROTO:   "proto",
	M3U8:    "m3u8",
}

var kindContentType = map[Kind]string{
	UNKNOWN: "application/x-unknown",
	JSON:    "application/json",
	HTML:    "text/html",
	CSV:     "text/csv",
	TSV:     "text/tsv",
	XML:     "text/xml",
	YAML:    "application/x-yaml",
	MSG:     "application/x-msgpack",
	PROTO:   "application/x-protobuf",

	// See section 3.1 of http://tools.ietf.org/html/draft-pantos-http-live-streaming-08#section-3.1
	M3U8: "application/vnd.apple.mpegurl",
}

func KindText(ext Kind) string        { return kindText1[ext] }
func KindContentType(ext Kind) string { return kindContentType[ext] }

func (it Kind) Is(that Kind) bool      { return it == that }
func (it Kind) IsJSON() bool           { return it.Is(JSON) }
func (it Kind) IsHTML() bool           { return it.Is(HTML) }
func (it Kind) IsCSV() bool            { return it.Is(CSV) }
func (it Kind) IsTSV() bool            { return it.Is(TSV) }
func (it Kind) IsXML() bool            { return it.Is(XML) }
func (it Kind) IsYAML() bool           { return it.Is(YAML) }
func (it Kind) IsMSG() bool            { return it.Is(MSG) }
func (it Kind) IsProto() bool          { return it.Is(PROTO) }
func (it Kind) IsM3U8() bool           { return it.Is(M3U8) }
func (it Kind) IsUnknown() bool        { return it.Is(UNKNOWN) }
func (it Kind) GetContentType() string { return KindContentType(it) }
func (it Kind) ContentType() string    { return it.GetContentType() }
func (it Kind) String() string         { return kindText2[it] }

func (it *Kind) UnmarshalYAML(unmarshal func(interface{}) error) error {
	v := ""
	unmarshal(&v)
	*it = NewKind(v)
	if it.IsUnknown() {
		return ErrUnknownKind
	}
	return nil
}

func NewKind(s string) Kind {
	switch s {
	case ".json", ".JSON", "json", "JSON", CONTENT_TYPE_JSON:
		return JSON
	case ".html", ".HTML", "html", "HTML", CONTENT_TYPE_HTML:
		return HTML
	case ".csv", ".CSV", "csv", "CSV", CONTENT_TYPE_CSV:
		return CSV
	case ".tsv", ".TSV", "tsv", "TSV", CONTENT_TYPE_TSV:
		return TSV
	case ".xml", ".XML", "xml", "XML", CONTENT_TYPE_XML:
		return XML
	case ".yaml", ".yml", "yaml", "yml", "YAML", "YML", CONTENT_TYPE_YAML:
		return YAML
	case ".msg", ".msgpack", "msg", "msgpack", "MSG", "MSGPACK", CONTENT_TYPE_MSG:
		return YAML
	case ".proto", ".PROTO", "proto", "PROTO", CONTENT_TYPE_PROTO:
		return PROTO
	case ".m3u8", ".M3U8", "m3u8", "M3U8", CONTENT_TYPE_M3U8, CONTENT_TYPE_M3U8_1:
		return M3U8
	}

	return UNKNOWN
}

func GuessKind(s string) Kind {
	it := NewKind(s)
	if it.IsUnknown() {
		it = JSON
	}
	return it
}
