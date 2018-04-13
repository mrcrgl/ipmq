package api

type Kind uint8

type VersionKind struct {
	MinVersion uint8
	MaxVersion uint8
	Kind       Kind
}

func (vk *VersionKind) VersionKind() VersionKind {
	return VersionKind{vk.MinVersion, vk.MaxVersion, vk.Kind}
}

type versionKind interface {
	VersionKind() VersionKind
}
