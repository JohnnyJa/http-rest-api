package apiclient

import "testing"

func TestApiClient(t *testing.T) *APIClient{
	t.Helper()

	config := NewConfig()

	cl := New(config)
	return cl
}