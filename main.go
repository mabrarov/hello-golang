package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CountDownLatch struct {
	cond  *sync.Cond
	value int
}

func NewCountDownLatch(value int) *CountDownLatch {
	return &CountDownLatch{cond: sync.NewCond(&sync.Mutex{}), value: value}
}

func (c *CountDownLatch) CountDown() {
	c.cond.L.Lock()
	if c.value > 0 {
		c.value--
	}
	if c.value == 0 {
		c.cond.Broadcast()
	}
	c.cond.L.Unlock()
}

func (c *CountDownLatch) Wait() {
	c.cond.L.Lock()
	for c.value > 0 {
		c.cond.Wait()
	}
	c.cond.L.Unlock()
}

type FetchCache struct {
	mu    sync.Mutex
	cache map[string]bool
}

func (c *FetchCache) TryPut(value string) bool {
	c.mu.Lock()
	if c.cache[value] {
		c.mu.Unlock()
		return false
	}
	c.cache[value] = true
	c.mu.Unlock()
	return true
}

var urlCache = FetchCache{cache: make(map[string]bool)}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	if !urlCache.TryPut(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	latch := NewCountDownLatch(len(urls))
	for _, u := range urls {
		go func() {
			Crawl(u, depth-1, fetcher)
			latch.CountDown()
		}()
	}
	fmt.Printf("found: %s %q\n", url, body)
	latch.Wait()
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Duration(rand.Intn(500-10)+10) * time.Millisecond)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
