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
	assert.EqualValues(t, 0, res.Price)
	assert.Error(t, res.Error)

	urls := []string{
		"http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=1",
	 "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=2",
	 "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=3",
	 "http://inv-nets.admixer.net/test-dsp/dsp?responseType=1&profile=4",
	}

	for _, url := range urls {
		go cl.GetPrice(url, ch)
	}

	prices := make([]float64, 0)

	for i := 0; i < len(urls); i++ {
		res = <-ch
		prices = append(prices, res.Price)
	}

	log.Println(prices)
}