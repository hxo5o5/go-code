/*练习：Web 爬虫
在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。
修改 Crawl 函数来并行地抓取 URL，并且保证不重复。
提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！
*/
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

//构建一个缓存map
var (
	m = make(map[string]int)
	l sync.Mutex     //访问互斥锁
	i sync.WaitGroup //等待组
)

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	defer i.Done()
	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}
	//存入数据 需要同步锁 因为这是在子线程中
	l.Lock()
	//如果没有爬
	if m[url] == 0 {
		//存入爬取的url
		m[url]++
		//深度递减
		depth--
		for _, u := range urls {
			i.Add(1)
			go Crawl(u, depth, fetcher) //继续爬取
		}
	}
	l.Unlock()
	// 下面并没有实现上面两种情况：
	/*if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return*/
}

func main() {
	i.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	i.Wait()
	for k, _ := range m {
		fmt.Println(k)
	}
	fmt.Println("over")
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
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
