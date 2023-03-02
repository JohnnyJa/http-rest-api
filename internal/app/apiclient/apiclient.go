package apiclient

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

type APIClient struct {
	config *Config
	client *http.Client
}

var (
	validate *validator.Validate
)

func New(config *Config) *APIClient {
	return &APIClient{
		config: config,
		client: &http.Client{},
	}
}

type PriceResult struct{
	Price float64
	Error error
}



func (c *APIClient)GetPrice(url string, ch chan PriceResult){
	validate = validator.New()

	r, err := c.client.Get(url)
	if err != nil {
		ch <- PriceResult{0, err}
		
	}
	defer r.Body.Close()

	type response struct{
		Price float64 `json:"price" validate:"required"`
	}

	resp := &response{}
	if err := json.NewDecoder(r.Body).Decode(resp); err != nil {
		ch <- PriceResult{0, err}
		return
	}

	if err:=validate.Struct(resp); err != nil {
		ch <-PriceResult{0, err} 
		return
	}

	ch <- PriceResult{resp.Price,  nil}
}





