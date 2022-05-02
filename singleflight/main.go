package main

import (
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/singleflight"
)

var group singleflight.Group

func main() {
	keys := []string{
		"key1", "key2", "key3", "key4", "key5",
		"key1", "key2", "key3", "key4", "key5",
	}

	for _, key := range keys {
		key := key
		go group.Do(key, func() (interface{}, error) {
			log.Println(key)
			time.Sleep(time.Second)
			return key, nil
		})
	}
	for {
		runtime.Gosched()
	}

}
