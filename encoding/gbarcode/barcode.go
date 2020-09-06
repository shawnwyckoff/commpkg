package gbarcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/tuotoo/qrcode"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"
)

// surffix: "gif","jpg","jpeg","png"
func EncodeQrCode(content string, size int, surffix string, w io.Writer) error {
	if len(content) == 0 {
		return errors.Errorf("Empty input content")
	}
	if size <= 0 || size > 10000 {
		size = 200
	}

	// Create the barcode
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return err
	}

	// Scale the barcode to 200x200 pixels
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return err
	}

	// encode the barcode as png
	surffix = strings.ToLower(surffix)
	switch surffix {
	case "png":
		return png.Encode(w, qrCode)
	case "jpg":
		return jpeg.Encode(w, qrCode, nil)
	case "jpeg":
		return jpeg.Encode(w, qrCode, nil)
	case "gif":
		return gif.Encode(w, qrCode, nil)
	}
	return errors.Errorf("Unsupported output image surffix %s", surffix)
}

func DecodeQrCode(r io.Reader) (string, error) {
	matrix, err := qrcode.Decode(r)
	if err != nil {
		return "", err
	}
	return matrix.Content, nil
}
