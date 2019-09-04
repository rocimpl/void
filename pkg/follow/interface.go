package follow

import "errors"

var (
    ErrUnknownFollow = errors.New("void: Unknown follow type")
)

type Follow interface {
    Process(seek int64) (read []byte, err error)
    Stop()
}
