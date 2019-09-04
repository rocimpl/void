package parser

import (
    "errors"
    "github.com/rocimpl/void/pkg/types"
)

var (
    ErrUnknownParser = errors.New("void: Unknown parser type")
)

type Parser interface {
    Parse(sequence []byte) (types.LogFormat, error)
}
