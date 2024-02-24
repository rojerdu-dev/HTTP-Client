package client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Parallel()
	t.Run("happy path - can hit the api and return a pokemon", func(*testing.T) {
		c := NewClient()
		pokemon := "pikachu"
		result, err := c.GetPokemonByName(context.Background(), pokemon)
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", result.Name)
		fmt.Println(result)
	})

	t.Run("sad path - return an error when pokemon doesn't exist", func(*testing.T) {
		c := NewClient()
		_, err := c.GetPokemonByName(context.Background(), "not-in-existence")
		assert.Error(t, err)
	})
}

func TestWithAPIURL(t *testing.T) {
	t.Parallel()
	t.Run("happy path - testing the WithAPIURL option function", func(*testing.T) {
		c := NewClient(
			WithAPIURL("my-test-url"),
		)
		assert.Equal(t, "my-test-url", c.apiURL)
	})
}

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()
	t.Run("happy path - tests with httpclient works", func(*testing.T) {
		c := NewClient(
			WithAPIURL("my-test-url"),
			WithHTTPClient(&http.Client{
				Timeout: 1 * time.Second,
			}),
		)
		assert.Equal(t, 1*time.Second, c.httpClient.Timeout)
	})
}
