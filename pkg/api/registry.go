package api

import (
	"fmt"
)

func newRegistry() *registry {
	return &registry{
		m: make(map[VersionKind]versionKind),
	}
}

type registry struct {
	m map[VersionKind]versionKind
}

func (r *registry) MustRegister(vks ...versionKind) {
	for _, vk := range vks {
		// TODO ensure is pointer

		if _, ok := r.m[vk.VersionKind()]; ok {
			panic(fmt.Sprintf("VersionKind %v already registered!", vk.VersionKind()))
		}

		r.m[vk.VersionKind()] = vk
	}
}
