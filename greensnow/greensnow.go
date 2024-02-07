package greensnow

import (
	"github.com/pterm/pterm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"time"
)

type GreenSnow struct {
	dataDir  string
	dataPath string
	date     string
	feeds    []Feed
}

type Feed struct {
	name, filename, url string
}

var (
	dataDir string = "greensnow_data"
	feedURL        = map[string]string{
		"IPBlocklist": "https://blocklist.greensnow.co/greensnow.txt",
	}
)

func Download(pathName string) {
	pterm.Info.Println("Greensnow.co: Downloading Feeds")

	tf := newThreatfetch(pathName, feedURL, dataDir)

	//pterm.Info.Println(tf)

	// Determine if pathName exists. If it doesn't then dataDir also
	// does not exist.
	_, err := os.Stat(tf.dataPath)

	if err != nil {
		// Create pathName and dataDir
		createDataDir(tf.dataPath)

		// Download feed data to pathName
		downloader(tf)
	} else {
		pterm.Debug.Println("Directory exists.")

		// If fullPath does exist we need to check existing files timestamp
		// as files are only generated once a day @ 11PM UTC.
		dailyGenerationCheck(tf)
	}

}

func newThreatfetch(path string, feedURL map[string]string, dataDir string) *GreenSnow {

	t := new(GreenSnow)
	t.dataDir = path
	t.dataPath = (path + "/" + dataDir)
	t.date = time.Now().Format(time.DateOnly)
	t.feeds = make([]Feed, 0)

	for name, url := range feedURL {
		f := Feed{
			name:     name,
			filename: filepath.Base(url),
			url:      url,
		}
		t.feeds = append(t.feeds, f)
	}

	return t
}

func downloader(tf *GreenSnow) {

	// Set file download progress bar
	p, _ := pterm.DefaultProgressbar.WithTotal(len(feedURL)).WithTitle("Download Status").WithRemoveWhenDone(true).Start()

	for entry, _ := range tf.feeds {
		p.UpdateTitle("Downloading " + tf.feeds[entry].name + " from " + tf.feeds[entry].url)

		feed_data, err := http.Get(tf.feeds[entry].url)

		if err != nil {
			pterm.Error.Println("Error accessing threat feed: " + tf.feeds[entry].name)
		}

		defer feed_data.Body.Close()

		// Check response code to determine if file was accessible.
		if feed_data.StatusCode != 404 {

			file := (tf.dataPath + "/" + tf.date + "-" + tf.feeds[entry].filename)
			f, err := os.Create(file)

			if err != nil {
				pterm.Error.Println("Error creating file: " + tf.feeds[entry].filename)
			}

			defer f.Close()

			_, err = io.Copy(f, feed_data.Body)

			if err != nil {
				pterm.Error.Println("Error with file data: " + tf.feeds[entry].filename)
			}

			pterm.Success.Println(tf.feeds[entry].name + " -> " + file)
		} else {
			pterm.Error.Println("Error downloading " + tf.feeds[entry].url + ". Status: " + feed_data.Status)
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

func dailyGenerationCheck(tf *GreenSnow) {

	// empty struct for missing files
	var missing_tf GreenSnow
	missing_tf.dataDir = tf.dataDir
	missing_tf.dataPath = tf.dataPath
	missing_tf.date = tf.date
	missing_tf.feeds = make([]Feed, 0)
	pterm.Debug.Println(missing_tf)

	// Loop through feed files/urls to check if todays exists
	for f, _ := range tf.feeds {
		// create full file path for os.Stat
		file := (tf.dataPath + "/" + tf.date + "-" + tf.feeds[f].filename)

		// Call os.Stat on file
		fs, err := os.Stat(file)

		pterm.Debug.Println(fs)

		if fs == nil {
			pterm.Error.Println(err)
			f := Feed{
				name:     tf.feeds[f].name,
				filename: tf.feeds[f].filename,
				url:      tf.feeds[f].url,
			}
			missing_tf.feeds = append(missing_tf.feeds, f)

		}

		pterm.Debug.Println(missing_tf)
		pterm.Debug.Println(len(missing_tf.feeds))

	}

	// Compare missing_tf to passed tf. If same, call downloader with tf
	// else call downloader with missing_tf

	if reflect.DeepEqual(tf, missing_tf) {
		downloader(tf)
	} else if len(missing_tf.feeds) != 0 {
		downloader(&missing_tf)
	} else {
		pterm.Success.Println("All feed data updated.")
	}

}
