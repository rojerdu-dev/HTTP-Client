package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, PokemonFetchErr{
			Message:    "non-200 status code from the API",
			StatusCode: http.StatusNotFound}, err)
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

func TestHTTPServer(t *testing.T) {
	t.Parallel()
	t.Run("happy path - able to hit locally running server", func(*testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `{"name":"pikachu", "height":10}`)
			}),
		)
		defer ts.Close()
		fmt.Println(ts.URL)

		c := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := c.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, 10, poke.Height)
	})
	t.Run("sad path - able to handle 500 status from API", func(*testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))
		defer ts.Close()
		c := NewClient(
			WithAPIURL(ts.URL),
		)
		poke, err := c.GetPokemonByName(context.Background(), "pikachu")
		assert.Error(t, err)
		assert.Equal(t, 0, poke.Height)
	})
}
