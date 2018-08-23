package genius

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestExtraction(t *testing.T) {
	var lyrics string
	var fin, err = os.Open("testdata/oxxxymiron-tentacles-lyrics.html")

	if err != nil {
		t.Fatalf("failed to open file: %s", err)
	}

	defer fin.Close()

	if lyrics, err = NewExtractor(fin).Extract(); err != nil {
		t.Errorf("failed to extract lyrics: %s", err)
	} else {
		t.Log(lyrics)
	}

	var reader = strings.NewReader(lyrics)
	var scanner = bufio.NewScanner(reader)
	var counter int

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" {
			counter++
		}
	}

	if counter != 100 {
		t.Errorf("wrong number of lines: %d", counter)
	}
}
