package client

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientCanHitAPI(t *testing.T) {
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
