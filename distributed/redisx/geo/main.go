package main

import (
	. "distributed/redisx"
	"github.com/go-redis/redis"
)

func main() {
	Client.GeoAdd("user", &redis.GeoLocation{
		Name:      "user1",
		Longitude: 116.48105,
		Latitude:  39.996794,
		Dist:      0,
		GeoHash:   0,
	})

	Client.GeoAdd("user", &redis.GeoLocation{
		Name:      "user2",
		Longitude: 116.38105,
		Latitude:  39.996700,
		Dist:      0,
		GeoHash:   0,
	})

	Out(Client.GeoDist("user", "user1", "user2", "km").Result())

	result, _ := Client.GeoPos("user", "user1").Result()
	Out(result[0].Longitude, result[0].Latitude) // 2

	Out(Client.GeoHash("user", "user1").Result())

	Out(Client.GeoRadiusByMember("user", "user1", &redis.GeoRadiusQuery{
		Radius:      10,
		Unit:        "km",
		WithCoord:   false,
		WithDist:    false,
		WithGeoHash: false,
		Count:       10,
		Sort:        "asc",
		Store:       "",
		StoreDist:   "",
	}).Result())
}
