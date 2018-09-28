// Utility lyrics allow to fetch lyrics from internet databases and store it on
// filesystem.
package main

import (
	"errors"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/daskol/go.lyrics/logging"
)

var logger = logging.Default()
var accessToken = os.Getenv("GENIUS_ACCESS_TOKEN")
var artistIDs = flag.String("artist-ids", "", "Identifier of artist to fetch.")
var outdir = flag.String("outdir", "lyrics", "Output directory.")

func init() {
	flag.Parse()
}

func getArtistIDs() ([]int, error) {
	var parts = strings.Split(*artistIDs, ",")
	var artistIDs = make([]int, len(parts))

	if len(parts) == 0 {
		return nil, errors.New("too few artist ids")
	}

	for idx, part := range parts {
		if val, err := strconv.Atoi(part); err != nil {
			return nil, err
		} else {
			artistIDs[idx] = val
		}
	}

	return artistIDs, nil
}

func main() {
	var fetcher *LyricsFetcher
	var artistIDs, err = getArtistIDs()

	if err != nil {
		logger.Fatalf("failed to parse artist ids: %s", err)
	}

	logger.Infof("create instance of lyrics fetcher")

	if fetcher, err = NewLyricsFetcher(accessToken, *outdir); err != nil {
		logger.Fatalf("failed to intantiate lyrics fetcher: %s", err)
	}

	logger.Infof("fetch lyrics for artists: ids=%+v", artistIDs)

	if err = fetcher.FetchForArtists(artistIDs); err != nil {
		logger.Fatalf("failed to fetch lyrics: %s", err)
	}

	logger.Infof("done.")
}
