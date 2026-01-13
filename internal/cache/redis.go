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

func (c *Cache) Get(ctx context.Context, code string) (string, error){
	val, err := c.client.Get(ctx, code).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("chave não encontrada")
	}

	if err  != nil {
		return "", err
	}

	return val, nil
}

func (c *Cache) Save(ctx context.Context, 
					code string, 
					originalUrl string, 
					ttl time.Duration) error {
	return c.client.Set(ctx, code, originalUrl, ttl).Err()
} 

func (c *Cache) Close() error {
	return c.client.Close()
}