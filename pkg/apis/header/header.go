package header

import "github.com/mrcrgl/ipmq/pkg/api"

type Request struct {
	api.VersionKind
	Length        uint16
	CorrelationID uint16
}

type Response struct {
	api.VersionKind
	Length        uint16
	CorrelationID uint16
}
