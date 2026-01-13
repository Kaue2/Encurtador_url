package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

/* * 
* Abre a conexão com o Redis 
*/
func NewCache(addr string) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})

	// testando conexão
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("Erro: falha ao conectar no redis: %w", err)
	} 

	return &Cache{client: client}, nil 
}

func (c *Cache) Close() error {
	return c.client.Close()
}