package produce

import (
	"github.com/mrcrgl/ipmq/pkg/apis/header"
	"github.com/mrcrgl/ipmq/pkg/protocol"
)

type PushBlockRequest struct {
	header.Request
}

type PushBlockResponse struct {
	header.Response
}

func PushBlockHandler(vk protocol.VersionKind, decoder bool) (protocol.ByteEncoder, error) {

}
