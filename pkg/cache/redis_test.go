package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"testing"
)

func TestRedisClient(t *testing.T) {
	var ctx = context.Background()
	opt, _ := redis.ParseURL(os.Getenv("REDIS_PASS"))
	client := redis.NewClient(opt)

	val := client.Get(ctx, "foo").Val()
	fmt.Println(val)
}

func TestRedis(t *testing.T) {

	// cache.SetString(context.Background(), "foo", "bar", 5*time.Second)
	// t.Run("Get", func(t *testing.T) {
	// 	foo := cache.GetString(context.Background(), "foo")
	// 	if foo != "bar" {
	// 		t.Fail()
	// 	}
	// })
	t.Run("Get", func(t *testing.T) {
		c := NewRedisCache()
		foo, err := c.GetString(context.Background(), "foo")
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		if foo != "bar" {
			fmt.Println(foo)
			t.Fail()
		}
	})
}
