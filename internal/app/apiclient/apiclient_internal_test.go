package apiclient

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiClient_GetPrice(t *testing.T){
	cl := TestApiClient(t)
	ch := make(chan PriceResult)
	url :="http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=1"
	go cl.GetPrice(url, ch)

	res := <-ch

	assert.NoError(t, res.Error)
	assert.EqualValues(t, 10, res.Price)

	url ="http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=3"

	go cl.GetPrice(url, ch)

	res = <-ch
	log.Println(res.Price)
	assert.Error(t, res.Error)
}