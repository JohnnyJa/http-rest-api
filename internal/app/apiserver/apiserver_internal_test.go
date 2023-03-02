package apiserver

import (
	"testing"

	"github.com/JohnnyJa/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestApiServer_GetMaxSize(t *testing.T){
	s := New(NewConfig())
	s.Start()

	databaseURL := "sqlserver://MyPC/?database=UriDb"

	store.TestStore(t, databaseURL)

	ids := []int {1,2,3,4}

	_, err := s.GetMaxSize(ids)
	assert.Error(t, err)

	err = FillTable(t, s)	
	assert.NoError(t,err)


	ids = []int {1,2,4}
	max, err := s.GetMaxSize(ids)
	assert.NoError(t,err)
	assert.EqualValues(t, 20, max)
}