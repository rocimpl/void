package disk_file_follow

import (
    "bufio"
    "errors"
    "github.com/rocimpl/void/pkg/parser"
)

var (
    ErrDiskFileFollowFileNotSet = errors.New("disk_file_follow: Param 'file' not set")
)

type DiskFileFollow struct {
    file      string
    belt      parser.Parser
    sequencer bufio.Scanner
}

func NewDiskFileFollow(params map[string]string, parser parser.Parser) (*DiskFileFollow, error) {
    file, found := params["file"]
    if !found {
        return nil, ErrDiskFileFollowFileNotSet
    }

    return &DiskFileFollow{
        file: file,
        belt: parser,
    }, nil
}

func (f *DiskFileFollow) Process(seek int64) (read []byte, err error) {
    f.sequencer.Scan()

    f.belt.Parse(f.sequencer.Bytes())
}

func (f *DiskFileFollow) Stop() {
    panic("implement me")
}
