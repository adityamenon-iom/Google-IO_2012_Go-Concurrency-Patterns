package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
        return func(query string) Result {
	          time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	          return Result(fmt.Sprintf("%s result for %q\n", kind, query))
        }
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
