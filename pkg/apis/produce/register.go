package produce

import (
	"github.com/mrcrgl/ipmq/pkg/api"
	"github.com/mrcrgl/ipmq/pkg/api/scheme"
)

func init() {
	api.Scheme.AddFunction(
		api.KindPushBlockRequest,
		scheme.Version{},
		PushBlockHandler,
	)
}
