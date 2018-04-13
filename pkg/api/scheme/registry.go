package scheme

import (
	"fmt"

	"errors"

	"github.com/mrcrgl/ipmq/pkg/api"
	"github.com/mrcrgl/ipmq/pkg/api/machinery"
)

var (
	ErrFuncNotImplemented = errors.New("api function not implemented")
	ErrFuncIncompatible   = errors.New("api function does not support requested version")

	blankFunc = APIFunction{}
)

type Registry struct {
	reg map[api.Kind]APIFunction
}

func (ar Registry) Lookup(kind, version uint8) (APIFunction, error) {
	a, ok := ar.reg[kind]
	if !ok {
		return blankFunc, ErrFuncNotImplemented
	}

	if !a.Version.Matches(version) {
		return blankFunc, ErrFuncIncompatible
	}

	return a, nil
}

func (ar Registry) AddFunction(kind api.Kind, version Version, handler machinery.Handler) {
	ar.Add(kind, NewFunction(kind, version, handler))
}

func (ar Registry) Add(kind api.Kind, a APIFunction) {
	if _, exists := ar.reg[kind]; exists {
		panic(fmt.Sprintf("handler of kind %s already registered", kind))
	}
	ar.reg[kind] = a
}

func (ar Registry) Len() int {
	return len(ar.reg)
}

func (ar Registry) Map() map[int16]APIFunction {
	return ar.reg[:]
}
