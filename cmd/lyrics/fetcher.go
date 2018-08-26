package main

import (
	"errors"
	"os"
	"path"

	"github.com/daskol/go.lyrics/genius"
	"github.com/daskol/go.lyrics/logging"
)

// LyricsFetcher incapsulate operation on top of API clients and filesystem
// operations.
type LyricsFetcher struct {
	logger logging.Logger
	api    *genius.API
	outdir string
}

func NewLyricsFetcher(token, outdir string) (*LyricsFetcher, error) {
	var logger = logging.Default()

	if token == "" {
		return nil, errors.New("empty token")
	}

	if outdir == "" {
		outdir = "."
	}

	if _, err := os.Stat(outdir); os.IsExist(err) {
		logger.Infof("output directory exists: %s", outdir)
	} else if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		return nil, err
	}

	logger.Infof("create genius api client")
	var api = genius.New(token, logger)
	var fetcher = &LyricsFetcher{
		logger: logger,
		api:    api,
		outdir: outdir,
	}

	return fetcher, nil
}

func (l *LyricsFetcher) FetchForArtist(artistID int) error {
	l.logger.Infof("get artist info: artist_id=%d", artistID)
	var err error
	var artist *genius.Artist
	var songs []genius.Song

	if artist, err = l.api.GetArtist(artistID); err != nil {
		l.logger.Errorf("failed to get artist: %s", err)
		return err
	}

	l.logger.Infof("fetch song list of artist: name=%s", artist.Name)

	if songs, err = l.api.GetArtistSongs(artistID); err != nil {
		l.logger.Errorf("failed to load songs: %s", err)
		return err
	}

	var dirname = path.Join(l.outdir, artist.Name)

	l.logger.Infof("create output directory for lyrics: dirname=%s", dirname)

	if _, err = os.Stat(dirname); os.IsExist(err) {
		l.logger.Infof("artist directory exists: %s", dirname)
	} else if err = os.MkdirAll(dirname, os.ModePerm); err != nil {
		l.logger.Errorf("failed to create artist directory: %s", err)
		return err
	}

	l.logger.Infof("start fetching lyrics: nosongs=%d", len(songs))

	for i, song := range songs {
		l.logger.Infof("[%d] fetch lyrics from %s", i, song.URL)

		var filename = path.Join(dirname, song.Path)
		var text string

		if text, err = l.api.GetLyrics(song.URL); err != nil {
			l.logger.Errorf("failed get page of lyrics: %s", err)
			continue
		}

		fout, err := os.Create(filename)

		if err != nil {
			l.logger.Errorf("failed to create file: %s", err)
			continue
		}

		defer fout.Close()

		if _, err := fout.WriteString(text); err != nil {
			l.logger.Errorf("failed to write lyics to file: %s", err)
		}
	}

	return nil
}

func (l *LyricsFetcher) FetchForArtists(artistIDs []int) error {
	var err error

	for _, artistID := range artistIDs {
		l.logger.Infof("fetch lyrics for #%3d: artist_id=%d", artistID)
		if err = l.FetchForArtist(artistID); err != nil {
			break
		}
	}

	return err
}
