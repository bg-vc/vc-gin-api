package pkg

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
)

func TestRedisDisabled(t *testing.T) {
	client, err := OpenRedis(context.Background(), RedisConfig{Enabled: false})
	if err != nil || client != nil {
		t.Fatalf("disabled Redis should return nil, nil; got %v, %v", client, err)
	}
}

func TestRedisCacheExample(t *testing.T) {
	server := miniredis.RunT(t)
	client, err := OpenRedis(context.Background(), RedisConfig{Enabled: true, Address: server.Addr()})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	type value struct {
		Name string `json:"name"`
	}
	if err := SetCache(context.Background(), client, "example", value{Name: "SIMPLE"}, time.Minute); err != nil {
		t.Fatal(err)
	}
	var loaded value
	found, err := GetCache(context.Background(), client, "example", &loaded)
	if err != nil || !found || loaded.Name != "SIMPLE" {
		t.Fatalf("unexpected cache result: found=%v value=%+v err=%v", found, loaded, err)
	}
}
