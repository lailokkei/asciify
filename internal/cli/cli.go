package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	asciify "github.com/toodemhard/asciify/internal/asciify-lib"
)

func Start() {
	var hFlag = flag.Bool("h", false, "")
	var iFlag = flag.Bool("i", false, "")
	var fFlag = flag.String("f", "", "")
	var sFlag = flag.Int("s", 20, "")
	var cFlag = flag.String("c", "standard", "")
	flag.Parse()

	if *hFlag {
		fmt.Println("idk...")
		os.Exit(0)
	}

	fmt.Println(*fFlag)
	img, err := asciify.DecodeImageFile(*fFlag)
	if err != nil {
		log.Fatal(err)
	}

	options := asciify.Options{CharSetName: *cFlag, Invert: *iFlag, ScaleHeight: *sFlag}
	fmt.Print(asciify.ImageToText(img, options))
}
