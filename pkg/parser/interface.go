package parser

import "github.com/rocimpl/void/pkg/types"

type Parser interface {
    Parse(sequence []byte) error
    Snapshot() []types.LogFormat
}
