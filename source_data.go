package ginsa

import (
	"archive/zip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	tagsURL       = "https://api.github.com/repos/zengin-code/source-data/git/refs/tags"
	archivePrefix = "https://github.com/zengin-code/source-data/archive/"
)

type SourceData struct {
	Tag string

	Banks    map[string]*Bank
	Branches map[string]map[string]*Branch
}

func FetchAllSourceData() ([]*SourceData, error) {
	resp, err := http.Get(tagsURL)
	if err != nil {
		return nil, err
	}

	refs := []Ref{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&refs)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	sources := make([]*SourceData, len(refs))
	for i, ref := range refs {
		sources[i] = &SourceData{
			Tag: ref.Ref[10:],
		}
	}

	return sources, nil
}

func (s *SourceData) DownloadURL() string {
	return archivePrefix + s.Tag + ".zip"
}

func (s *SourceData) Load() error {
	fp, err := ioutil.TempFile("", "ginsa-"+s.Tag)
	if err != nil {
		return err
	}

	defer os.Remove(fp.Name())

	res, err := http.Get(s.DownloadURL())
	if err != nil {
		return err
	}

	_, err = io.Copy(fp, res.Body)
	res.Body.Close()
	fp.Close()

	archive, err := zip.OpenReader(fp.Name())
	if err != nil {
		return err
	}
	defer archive.Close()

	s.Branches = map[string]map[string]*Branch{}

	for _, f := range archive.File {
		parts := strings.Split(f.Name, string(os.PathSeparator))

		if len(parts) < 3 {
			continue
		}

		if parts[2] == "banks.json" {
			banks := map[string]*Bank{}

			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			decoder := json.NewDecoder(rc)
			err = decoder.Decode(&banks)
			if err != nil {
				return err
			}

			s.Banks = banks
		} else if parts[2] == "branches" && filepath.Ext(f.Name) == ".json" {
			bankCode := filepath.Base(f.Name)[0:4]
			branches := map[string]*Branch{}

			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			decoder := json.NewDecoder(rc)
			err = decoder.Decode(&branches)
			if err != nil {
				return err
			}

			s.Branches[bankCode] = branches
		}
	}

	return nil
}
