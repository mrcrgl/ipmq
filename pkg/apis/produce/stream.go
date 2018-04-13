package produce

import "github.com/mrcrgl/ipmq/pkg/apis/header"

type PushStreamRequest struct {
	header.Request
}

type PushStreamResponse struct {
	header.Response
}
