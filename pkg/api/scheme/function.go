package scheme

import (
	"errors"

	"github.com/mrcrgl/ipmq/pkg/api"
	"github.com/mrcrgl/ipmq/pkg/api/machinery"
	"github.com/travisjeffery/jocko"
	"github.com/travisjeffery/jocko/protocol"
)

var (
	ErrHandlerNotSetUp = errors.New("api handler function not set up")
)

func NewFunction(kind api.Kind, v Version, h machinery.Handler) APIFunction {
	return APIFunction{enabled: true, Kind: kind, Version: v, handler: h}
}

type APIFunction struct {
	enabled bool
	Kind    api.Kind
	Version Version
	handler machinery.Handler
}

func (af APIFunction) Handle(kv protocol.APIKeyVersion, decoder *protocol.ByteDecoder, broker jocko.Broker) (protocol.ByteEncoder, error) {
	if af.handler == nil {
		return nil, ErrHandlerNotSetUp
	}
	return af.handler(kv, decoder, broker)
}
