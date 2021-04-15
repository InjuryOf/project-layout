package di

type TraceMessage struct {
	TraceId  string
	TraceUrl string
}

func NewTraceMessage() TraceMessage {
	return TraceMessage{TraceId: "T20201220", TraceUrl: "baidu.com"}
}
