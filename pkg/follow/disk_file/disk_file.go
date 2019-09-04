package disk_file_follow

import (
    "errors"
    "github.com/rocimpl/void/pkg/parser"
)

var (
    ErrDiskFileFollowFileNotSet = errors.New("disk_file_follow: Param 'file' not set")
)

type DiskFileFollow struct {
    file   string
    parser parser.Parser
}

func NewDiskFileFollow(params map[string]string, parser parser.Parser) (*DiskFileFollow, error) {
    file, found := params["file"]
    if !found {
        return nil, ErrDiskFileFollowFileNotSet
    }

    return &DiskFileFollow{
        file:   file,
        parser: parser,
    }, nil
}

func (f *DiskFileFollow) Process() error {
    panic("implement me")
}

func (f *DiskFileFollow) Stop() {
    panic("implement me")
}
