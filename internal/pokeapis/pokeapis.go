package pokeapis

import (
	"net/http"
	"time"

	"github.com/keshavsharma98/pokedexcli/internal/pokecache"
)

func NewCLient(interval time.Duration) *Client {
	return &Client{
		cache: *pokecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
