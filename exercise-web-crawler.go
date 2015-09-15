package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var mod_fmt = log.New(os.Stdout, "", 0)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Page struct {
	url   string
	depth int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	ch := make(chan *map[string]bool)
	m := make(map[string]bool)
	var wg sync.WaitGroup
	var Fetch func(url string, depth int)
	Fetch = func(url string, depth int) {
		defer func() {
			ch <- &m
			wg.Done()
		}()

		if depth <= 0 {
			return
		}
		m := <-ch
		_, exists := (*m)[url]
		if exists {
			return
		}
		body, urls, err := fetcher.Fetch(url)
		(*m)[url] = true
		if err != nil {
			mod_fmt.Println(err)
			return
		}
		mod_fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			go Fetch(u, depth-1)
		}
	}

	go Fetch(url, depth)
	ch <- &m

	wg.Wait()
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
