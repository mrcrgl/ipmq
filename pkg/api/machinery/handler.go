package machinery

import "github.com/mrcrgl/ipmq/pkg/protocol"

type Handler func(vk protocol.VersionKind, decoder bool) (protocol.ByteEncoder, error)
