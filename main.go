package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

// docker run -it --rm -p 11211:11211 -p 11212:11212/udp memcached --udp-port 11212 -p 11211
func main() {
	c := memcache.New("localhost:11211")
	err := c.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	err = c.Set(&memcache.Item{
		Key: "key1",
		Value: append(
			bytes.Repeat([]byte("a"), 1400),
			bytes.Repeat([]byte("b"), 2800)...,
		),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	err = c.Set(&memcache.Item{
		Key: "key2",
		Value: append(
			bytes.Repeat([]byte("d"), 1500),
			bytes.Repeat([]byte("e"), 1500)...,
		),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	//t, err := c.Get("key3")
	//fmt.Printf("%s\n", t.Value)

	cr := memcache.New("udp://localhost:11212")
	c2 := memcache.NewSeparateReadWriter(cr, c)

	//t, err := c2.Get("key1")
	//fmt.Printf("%s\n", t.Value)
	//fmt.Println(string(t.Value), err)
	//return

	items, err := c2.GetMulti([]string{"key1", "key2"})
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, i := range items {
		fmt.Println(string(i.Value))
		fmt.Println("-----")
	}

	i, err := c2.Get("key1")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(i.Value))
}
