package scheme

type Version struct {
	min uint8
	max uint8
}

func (a Version) Matches(version uint8) bool {
	return a.min <= version && a.max >= version
}

func (a Version) Min() uint8 {
	return a.min
}

func (a Version) Max() uint8 {
	return a.max
}
