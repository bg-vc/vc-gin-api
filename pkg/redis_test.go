package pkg

import (
	"fmt"
	"testing"
)

func TestSetRedisVal(t *testing.T) {
	Init("../conf/config.yaml")
	key := fmt.Sprintf(`vc-gin-api:test:user:%v`, 1)
	user := &User{
		Name:    "vince",
		Age:     24,
		Address: "beijing",
	}
	if err := SetRedisVal(key, user, 0); err != nil {
		t.Errorf("pkg.SetRedisVal error, key:%v, err:%v", key, err)
		return
	}
	t.Log("done success")
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func TestGetRedisVal(t *testing.T) {
	Init("../conf/config.yaml")
	key := fmt.Sprintf(`vc-gin-api:test:user:%v`, 1)
	result := GetRedisVal(key)
	t.Logf("result:%v\n", result)
}

func TestRedisExists(t *testing.T) {
	Init("../conf/config.yaml")
	key := fmt.Sprintf(`vc-gin-api:test:user:%v`, 1)
	result := RedisExists(key)
	t.Logf("result:%v\n", result)
}
