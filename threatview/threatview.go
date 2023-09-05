package threatview

import (
	_ "fmt"
	"github.com/pterm/pterm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
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
	pterm.Info.Println("Threatview.io: Downloading Feeds")

	// Determine if pathName exists. If it doesn't then dataDir also
	// does not exist.
	_, err := os.Stat(pathName)
	pterm.Error.Println(err)

	fullPath := (pathName + "/" + dataDir)

	if err != nil {
		// Create pathName and dataDir
		createDataDir(fullPath)

		// Download feed data to pathName
		downloader(fullPath, feedURL)
	} else {

		// If fullPath does exist we need to check existing files timestamp
		// as files are only generated once a day @ 11PM UTC.
		dailyGenerationCheck(fullPath, feedURL)
	}

}

func downloader(path string, feedURL map[string]string) {

	// Set file download progress bar
	p, _ := pterm.DefaultProgressbar.WithTotal(len(feedURL)).WithTitle("Download Status").WithRemoveWhenDone(true).Start()

	for name, url := range feedURL {
		filename := filepath.Base(url)
		p.UpdateTitle("Downloading " + name + " from " + url)

		feed_data, err := http.Get(url)

		if err != nil {
			pterm.Error.Println("Error accessing threat feed: " + name)
		}

		defer feed_data.Body.Close()

		// Check response code to determine if file was accessible.
		if feed_data.StatusCode != 404 {

			date_stamp := time.Now().Format(time.DateOnly)
			file := path + "/" + date_stamp + "-" + filename
			f, err := os.Create(file)

			if err != nil {
				pterm.Error.Println("Error creating file: " + filename)
			}

			defer f.Close()

			_, err = io.Copy(f, feed_data.Body)

			if err != nil {
				pterm.Error.Println("Error with file data: " + filename)
			}

			pterm.Success.Println(name + " -> " + file)
		} else {
			pterm.Error.Println("Error downloading " + url + ". Status: " + feed_data.Status)
		}

		p.Increment()
	}

	//p.Stop()

}

func createDataDir(path string) {

	err := os.MkdirAll(path, 0750)
	if err != nil {
		pterm.Error.Println(path + "could not be created.")
		os.Exit(1)
	}
	pterm.Debug.Println("Directory created or already exists. " + path)
}

func dailyGenerationCheck(fullPath string, feedURL map[string]string) {

	date_stamp := time.Now().Format(time.DateOnly)
	var filenames []string

	for _, url := range feedURL {
		filename := (date_stamp + "-" + filepath.Base(url))
		pterm.Info.Println("Filename: " + filename)
		filenames = append(filenames, filename)
	}

	pterm.Info.Println(filenames)
	pterm.Info.Println(len(filenames))

	//filenames := feedFileSearch(filenames)

	//pterm.Info.Println("Passed path: " + fullPath)
	//pterm.Info.Println("Passed map: " + feedURL)
	/*
	   1. Generate "Todays" date (yyyy-mm-dd)
	   2. If files do not exist with todays date call downloader to get files
	   3. If files exist with todays date, check that we have all (8) files
	   4. If any files are missing build missingFeedURL, call downloader to get them.
	   5. If no files are missing, report files are up-to-date.
	*/
}

//func feedFileSearch(filenames string[]) string[] {

//}
