package cache

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/zu1k/proxypool/proxy"
)

var c = cache.New(cache.NoExpiration, 10*time.Minute)

func GetProxies() proxy.ProxyList {
	result, found := c.Get("proxies")
	if found {
		log.Println(len(result.(proxy.ProxyList)))
		return result.(proxy.ProxyList)
	}
	log.Println("Cache not found")
	return nil
}

func SetProxies(proxies proxy.ProxyList) {
	c.Set("proxies", proxies, cache.NoExpiration)
}

func SetString(key, value string) {
	c.Set(key, value, cache.NoExpiration)
}

func GetString(key string) string {
	result, found := c.Get(key)
	if found {
		return result.(string)
	}
	return ""
}
