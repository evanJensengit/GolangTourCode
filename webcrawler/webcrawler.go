// package main

// import (
// 	"fmt"
// 	//"sync"
// )

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func mapLearning() {
	var a map[string]bool
	a = make(map[string]bool)
	s1 := "www.fudge.com"
	a[s1] = true
	fmt.Println(a[s1])

	fmt.Println(a["beans"])

	if a["beans"] {
		fmt.Println(a["beans"])
	} else {
		fmt.Println("Else")
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	M.Lock()
	if UrlsFound[url] {
		M.Unlock()
		return
	}
	UrlsFound[url] = true
	body, urls, err := fetcher.Fetch(url)

	M.Unlock()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		M.Lock()
		if !UrlsFound[u] {
			UrlsFound[u] = true
			M.Unlock()
			Crawl(u, depth-1, fetcher)
			M.Lock()
		}
		M.Unlock()

	}
	return
}

func main() {
	//Crawl("https://golang.org/", 4, fetcher)
	mapLearning()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

var M sync.Mutex
var UrlsFound = make(map[string]bool)

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
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
