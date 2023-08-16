package cli

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
	"github.com/toodemhard/asciify/internal/asciify-lib"
)

func Start() {
	var cmdOptions struct {
		File    string `short:"f" long:"file" description:"Image file path to source"`
		Invert  bool   `short:"i" long:"invert" description:"Invert the values of the image"`
		CharSet string `short:"c" long:"charset" description:"Set of characters to use in output" default:"simple"`
		Scale   int    `short:"s" long:"scale" description:"Width of output in number of characters" default:"20"`
	}

	_, err := flags.Parse(&cmdOptions)
	if err != nil {
		log.Fatal(err)
	}

	img, err := asciify.DecodeImageFile(cmdOptions.File)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmdOptions)

	options := asciify.Options{
		CharSetName: cmdOptions.CharSet,
		Invert:      cmdOptions.Invert,
		ScaleWidth:  cmdOptions.Scale,
	}

	fmt.Print(asciify.ImageToText(img, options))
}
