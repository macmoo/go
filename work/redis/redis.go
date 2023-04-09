package main

// -------------------------------------------------
// go get github.com/go-redis/redis/v8
// go get github.com/redis/go-redis/v9
// -------------------------------------------------
import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8" // v8
	// "github.com/redis/go-redis/v9" // v9
)

// ######################################################
// String Set
func StringSet(client *redis.Client, ctx context.Context, key string, val string) {
	err := client.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println("redis.Client.Set Error:", err)
	}
}

// ######################################################
// String Get
func StringGet(client *redis.Client, ctx context.Context, key string) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("redis.Clieng.Get Error:", err)
	}
	fmt.Printf("key:%s, value:%s\n", key, val)
}

// ######################################################
// List Set
func ListSet(client *redis.Client, ctx context.Context, key string, val []string) {
	err := client.RPush(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println("redis.Client.RPush Error:", err)
	}
}

// ######################################################
// List Get
func ListGetRange(client *redis.Client, ctx context.Context, key string) {
	val, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println("redis.Client.LRange Error:", err)
	}
	fmt.Println(val)
}

// ######################################################
// List Index Get
func ListGetIndex(client *redis.Client, ctx context.Context, key string) {
	val, err := client.LIndex(ctx, key, 2).Result()
	if err != nil {
		fmt.Println("redis.Client.LIndex Error:", err)
	}
	fmt.Println(val)
}

// ######################################################
// Map Set
func MapSet(client *redis.Client, ctx context.Context, key string, m map[string]string) {
	for field, val := range m {
		err := client.HSet(ctx, key, field, val).Err()
		if err != nil {
			fmt.Println("redis.Client.HSet Error:", err)
		}
	}
}

// ######################################################
// Map Get
func MapGet(client *redis.Client, ctx context.Context, key string, field string) {
	val, err := client.HGet(ctx, key, field).Result()
	if err != nil {
		fmt.Println("redis.Client.RPush Error:", err)
	}
	fmt.Println(val)
}

// ######################################################
// Map Get All
func MapGetAll(client *redis.Client, ctx context.Context, key string) {
	val, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println("redis.Client.RPush Error:", err)
	}
	fmt.Println(val)
}

// ######################################################
func main() {
	fmt.Println("---------------------------------")
	// -------------------------------------------------
	// Connnect
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// -------------------------------------------------
	// 다른 방법
	// opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	// if err != nil {
	// 	panic(err)
	// }
	// rdb := redis.NewClient(opt)
	// -------------------------------------------------
	ctx := context.Background()
	// Execute Command
	key1 := "test_key"
	val, err := rdb.Get(ctx, key1).Result()
	fmt.Println("val:", val)
	fmt.Println("err:", err)
	fmt.Println("---------------------------------")
	key2 := "test_key2"
	val, err = rdb.Get(ctx, key2).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val == "":
		fmt.Println("Value is empty")
	}
	fmt.Println("---------------------------------")
	key3 := "StringKey3"
	val3 := "StringValue3"
	StringSet(rdb, ctx, key3, val3) // String SET
	StringGet(rdb, ctx, key3)       // String GET
	fmt.Println("---------------------------------")
	key4 := "ListKey4"
	// val4 := []string{"ListVal1", "ListVal2", "ListVal3"}
	// rpush ListKey4 ListVal1
	// rpush ListKey4 ListVal2
	// rpush ListKey4 ListVal3
	// rpush ListKey4 ListVal4
	// ListSet(rdb, ctx, key4, val4)
	ListGetRange(rdb, ctx, key4)
	ListGetIndex(rdb, ctx, key4)
	fmt.Println("---------------------------------")
	key5 := "MapKey5"
	val5 := map[string]string{"MKey1": "MVal1", "MKey2": "MVal2", "MKey3": "MVal3"}
	for field, val := range val5 {
		fmt.Printf("key:%s, value:%s\n", field, val)
	}
	fmt.Println("---------------------------------")
	MapSet(rdb, ctx, key5, val5)
	MapGet(rdb, ctx, key5, "MKey2")
	MapGetAll(rdb, ctx, key5)
	fmt.Println("---------------------------------")
}

// ---------------------------------
// val: test_value
// err: <nil>
// ---------------------------------
// Key does not exist
// ---------------------------------
// key:StringKey3, value:StringValue3
// ---------------------------------
// [ListVal1 ListVal2 ListVal3 ListVal4]
// ListVal3
// ---------------------------------
// key:MKey1, value:MVal1
// key:MKey2, value:MVal2
// key:MKey3, value:MVal3
// ---------------------------------
// MVal2
// map[MKey1:MVal1 MKey2:MVal2 MKey3:MVal3]
// ---------------------------------
