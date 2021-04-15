//+build wireinject
package di

import "github.com/google/wire"

func InitializeTraceMessage() TraceMessage {
	wire.Build(NewTraceMessage)
	return TraceMessage{}
}
