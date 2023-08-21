package asciify

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func DecodeImageFile(filepath string) (image.Image, error) {
	f, err := os.Open(filepath)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
