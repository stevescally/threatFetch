package main

import (
	"flag"
	"github.com/pterm/pterm"
	"github.com/stevescally/threatFetch/threatview"
	"os"
)

var (
	pathName = flag.String("p", "", "Pathname location where threat feed data exists or will be downloaded."+
		"\nDefault: <current working directory>/threatview_data")
)

func main() {

	// Threatview.io will be the default feed
	// Print banner of goGetThreatFeed
	// process arguments
	//  - Debug or info logging(default)
	// Based on input argument call other functionality
	//   Create threatview.io package
	//   Have all functions that relate to:
	//     - Checking if download directory exists
	//     - What is the date/time on the current files
	//     - Download the list of available files

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

	// Call ThreatFeed.io Downloader
	threatview.Download(*pathName)
}
