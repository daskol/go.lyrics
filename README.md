# Go.Lyrics

*lyrics fetcher written in Go for machine learning stuff*

## Overview

This package provides bindings to some internet databases of lyrics. At present
only [Genius](https://genius.com/) is supported. Its API client is located
[here](genius/).

Also, there is [a CLI tool](cmd/lyrics) which allow to fetch lyrics and store
it on filesystem. It requires access token that could be issued
[here](https://docs.genius.com). After that one can build utility and fetch all
songs of specified artist by artist id. It is allowed to pass multiple artist
ids separated with comma.

```bash
    go get -u github.com/daskol/go.lyrics/...
    export GENIUS_ACCESS_TOKEN=sdfakp9_dnsdf-_ufndslundlfkbsnd734fNUIbk_DDd
    lyrics --outdir lyrics --artist-ids 1949
```

## [Documentation](https://godoc.org/github.com/daskol/go.lyrics)

See documentation on [GoDoc](https://godoc.org/github.com/daskol/go.lyrics).
