package apiserver

import (
	"testing"

	"github.com/JohnnyJa/http-rest-api/internal/app/model"
)

func FillTable(t *testing.T, s *APIServer) error {
	t.Helper()
	urls := []model.UrlPackage{
		{UrlString: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=1"},
		{UrlString: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=2"},
		{UrlString: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=3"},
		{UrlString: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=4"},
	}

	for _, url := range urls {
		if _, err := s.store.UrlPackage().Create(&url); err != nil {
			return err
		}
	}
	return nil
}
