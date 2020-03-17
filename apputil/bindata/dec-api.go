package bindata

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/sys/gfs"
)

func Dec(fileHexString *string, output_binary_filename string) error {
	if fileHexString == nil {
		return errors.Errorf("fileHexString is nil")
	}
	if len(*fileHexString)%2 != 0 {
		return errors.Errorf("fileHexString length error")
	}
	buf, err := hexutil.Decode(*fileHexString)
	if err != nil {
		return err
	}
	return gfs.BytesToFile(buf, output_binary_filename)
}
