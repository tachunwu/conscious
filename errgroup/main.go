package main

import (
	"errors"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	keys := []string{
		"key1", "key2", "key3", "key4", "key5",
		"key6", "key7", "key8", "key9", "key10",
	}
	for i, key := range keys {
		i, key := i, key // create locals for closure below
		g.Go(func() error {

			if i == 9 {
				return errors.New("error case")
			}
			log.Println(key)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		log.Println(err)
	}
}
