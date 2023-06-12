package threatview

import (
	"github.com/pterm/pterm"
	"log"
	"os"
)

var (
	dataDir string = "threatview_data"
	feedURL        = map[string]string{
		"osintThreatFeed":      "https://threatview.io/Downloads/Experimental-IOC-Tweets.txt",
		"c2HuntFeed":           "https://threatview.io/Downloads/High-Confidence-CobaltStrike-C2 -Feeds.txt",
		"IPBlocklist":          "https://threatview.io/Downloads/IP-High-Confidence-Feed.txt",
		"domainBlocklist":      "https://threatview.io/Downloads/DOMAIN-High-Confidence-Feed.txt",
		"MD5HashBlocklist":     "https://threatview.io/Downloads/MD5-HASH-ALL.txt",
		"URLBlocklist":         "https://threatview.io/Downloads/URL-High-Confidence-Feed.txt",
		"bitcoinAddressIntel":  "https://threatview.io/Downloads/MALICIOUS-BITCOIN_FEED.txt",
		"SHAFileHashBlocklist": "https://threatview.io/Downloads/SHA-HASH-FEED.txt",
	}
)

func Download(pathName string) {
	pterm.Info.Println("Threatview.io: Downloading")
	pterm.Info.Println("Download Directory: ", pathName)

	// Confirm path and directory exists
	// Pass confirmed/created full path to downloader
	fullPath, err := createDataDir(pathName, dataDir)
   
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    // Call downloader pass path and list of URL's
    // only return err status
    downloader(fullPath, *feedURL)

	pterm.Info.Println("Full Path: ", fullPath)
	pterm.Info.Println("Error: ", err)

}

func downloader(path string, f ) {

    for id, url := range *f {
        
        fmt.Println(id, url)
    }

}

func createDataDir(path string, dir string) (string, error) {
	pterm.Info.Println("Path: ", path)
	pterm.Info.Println("Directory: ", dir)

	fullPath := (path + "/" + dir)

	err := os.MkdirAll(fullPath, 0750)
	if err != nil {
		log.Fatal(err)
	}
	// Accept the specified directory and source folder name to create
	// full path. /tmp/downloads/ + threatview
	// Return error if directory can not be created or can not be
	// confirmed as existing.
	return fullPath, err

}
