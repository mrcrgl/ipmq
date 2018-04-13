package produce

import "github.com/mrcrgl/ipmq/pkg/api"

func init() {
	api.Scheme.MustRegister(
		&PushBlockRequest{},
		&PushBlockResponse{},
		&PushStreamRequest{},
		&PushStreamResponse{},
	)
}
