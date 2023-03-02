package store_test

import (
	"testing"

	"github.com/JohnnyJa/http-rest-api/internal/app/model"
	"github.com/JohnnyJa/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUrlPackageRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("UrlPackage")

	u, err := s.UrlPackage().Create(&model.UrlPackage{
		UrlString: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=1",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUrlPackageRepository_FindById(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("UrlPackage")

	id := 1
	_, err :=s.UrlPackage().FindById(id)
	assert.Error(t, err)

	s.UrlPackage().Create(&model.UrlPackage{
		UrlString: "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=1",
	})

	u, err :=s.UrlPackage().FindById(id)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
