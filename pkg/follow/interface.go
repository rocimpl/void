package follow

import "github.com/rocimpl/void/pkg/parser"

type Follow interface {
    Start(parser parser.Parser) error
    Stop()
}
