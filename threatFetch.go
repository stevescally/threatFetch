package main

import (
	"flag"
	"github.com/pterm/pterm"
	"github.com/stevescally/threatFetch/threatview"
	"github.com/stevescally/threatFetch/greensnow"
	"os"
)

var (
	pathName = flag.String("p", "", "Pathname location where threat feed data exists or will be downloaded."+
		"\nDefault: <current working directory>/threatview_data")
)

func main() {

	// Generate BigLetters
	pterm.Println()
	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Threat Fetch")).Srender()
	pterm.Println(s)
	pterm.Description.Println("threatFetch downloads updated threat feed data.")

	// Process Options
	flag.Parse()

	if *pathName == "" {
		cdir, err := os.Getwd()

		if err != nil {
			pterm.Error.Println("Current working directory could not be determined.")
			os.Exit(1)
		}

		*pathName = cdir
	}

	// Call threat feed downloader
	threatview.Download(*pathName)
  greensnow.Download(*pathName)
}
