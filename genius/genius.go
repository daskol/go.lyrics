package genius

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type loggerWriter struct{}

func (l *loggerWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type API struct {
	baseURL url.URL
	client  http.Client
	logger  *log.Logger
	token   string
}

func New(token string, logger *log.Logger) *API {
	if logger == nil {
		logger = log.New(&loggerWriter{}, "", 0)
	}

	api := &API{}
	api.baseURL = url.URL{
		Scheme: "https",
		Host:   "api.genius.com",
	}
	api.logger = logger
	api.token = token
	return api
}

func (a *API) request(url string) (*Object, error) {
	var err error
	var req *http.Request
	var res *http.Response

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+a.token)

	if res, err = a.client.Do(req); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var code = strconv.Itoa(res.StatusCode)
		var msg = "genius: response status is [" + code + "] " + res.Status
		return nil, errors.New(msg)
	}

	var obj = new(Object)
	var dec = json.NewDecoder(res.Body)

	if err = dec.Decode(obj); err != nil {
		return nil, err
	}

	return obj, nil
}

func (a *API) GetArtist(id int) (*Artist, error) {
	var url url.URL = a.baseURL
	url.Path = "/artists/" + strconv.Itoa(id)

	a.logger.Printf("genius: fetch artist info: id=%d", id)

	if obj, err := a.request(url.String()); err != nil {
		return nil, err
	} else {
		return obj.Response.Artist, nil
	}
}

// getArtistSongs gets documents (songs) for the artist specified. By default,
// 20 items are returned for each request.
func (a *API) getArtistSongs(id, perPage, page int) ([]Song, error) {
	var params = make(url.Values)
	var url url.URL = a.baseURL

	params.Add("per_page", strconv.Itoa(perPage))
	params.Add("page", strconv.Itoa(page))

	url.Path = "/artists/" + strconv.Itoa(id) + "/songs"
	url.RawQuery = params.Encode()

	a.logger.Printf("genius: fetch page of songs: page=%d", page)

	if obj, err := a.request(url.String()); err != nil {
		return nil, err
	} else {
		return obj.Response.Songs, nil
	}
}

// GetArtistSongs gets all documents (songs) for the artist specified.
func (a *API) GetArtistSongs(id int) ([]Song, error) {
	const perPage = 20

	var err error
	var allSongs, songs []Song

	a.logger.Printf("genius: start fetching of all songs for artist %d", id)

	for page := 1; ; page++ {
		if songs, err = a.getArtistSongs(id, perPage, page); err != nil {
			return nil, err
		} else if len(songs) == 0 {
			break
		} else {
			allSongs = append(allSongs, songs...)
		}
	}

	return allSongs, nil
}

func (a *API) GetLyrics(uri string) (string, error) {
	var err error
	var req *http.Request
	var res *http.Response

	if req, err = http.NewRequest("GET", uri, nil); err != nil {
		return "", err
	}

	if res, err = a.client.Do(req); err != nil {
		return "", err
	}

	defer res.Body.Close()
	return NewExtractor(res.Body).Extract()
}
