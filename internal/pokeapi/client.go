package pokeapi

import (
	"net/http"
	"time"
	"github.com/NickG76/pokedexcli/internal/pokecache"
)

// client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}


// new client
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		cache: pokecache.NewCache(cacheInterval), 
		httpClient: http.Client{
			Timeout: timeout,
		},
	}


}
