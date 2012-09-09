/*
 * http://tour.golang.org/#70
 * Exercise: Web Crawler
 * 
 * In this exercise you'll use Go's concurrency features
 * to parallelize a web crawler.
 * 
 * Modify the Crawl function to fetch URLs in parallel
 * without fetching the same URL twice.
 */
// 注：此实现暂未添加url重复判断的功能
package main

import (
	"fmt"
)

// 定义Fether接口，包含一个Fetch函数
type Fetcher interface {
	// Fetch 返回一个对应url的body和此网页中包含的url链接
	Fetch(url string) (body string, urls []string, err error)
}

// 定义一个Fetch结果的结构体，包含网页信息和包含的内容
type FetchResult struct {
	url  string     // 当前网页的url
	depth int       // 当前网页所处的爬取深度
	body string     // 当前网页的内容
	urls []string   // 当前网页包含的其他链接
	err  error      // 当前网页包含的错误
}

// Crawl 使用 fetcher 来递归的爬网页
// 从url开始直到最大深度
func Crawl(url string, depth int, fetcher Fetcher, ch chan FetchResult) {
	var r FetchResult

	r.depth = depth
	if depth <= 0 {
		// 如果到达最大深度，返回此深度，并返回
		ch <- r
		return
	}

	r.url = url
	r.body, r.urls, r.err = fetcher.Fetch(url)
	if r.err != nil {
		// 如果网页包含错误，返回
		ch <- r
		return
	}

	// 为下层爬取准备管道
	subch := make(chan FetchResult, len(r.urls))
	// 递归爬取下层网页
	// 使用并发机制，同时访问所有的下层网页
	for _, u := range r.urls {
		go Crawl(u, depth-1, fetcher, subch)
	}

	// 统计深度比当前网页高的网页数目，如果等于下层网页数目，
	// 则说明下层网页已经爬取完毕，退出循环
	for count := 0; count < len(r.urls); {
		tmpr := <- subch
		if tmpr.depth == depth-1 {
			count++
		}
		if len(tmpr.url) > 0 {
			ch <- tmpr
		}
	}
	
	ch <- r
	return
}

func main() {
	ch := make(chan FetchResult, 1)
	// 新建线程进行爬取
	go Crawl("http://golang.org/", 4, fetcher, ch)

	// 利用管道来收集爬取到的网页信息
	// 一旦接收到第一层深度的网页信息，代表所有网页已经访问完毕，
	// 可以退出了。
	ok := true
	for ok {
		r := <- ch
		if r.depth == 4 {
			ok = false
		}
		// 如果有错误则打印错误，并跳至下一循环
		if (r.err != nil) {
			fmt.Print(r.err)
			continue
		}
		// 打印网页信息
		fmt.Printf("found: %s %q\n", r.url, r.body)
	}
}

////////////////////////////////////////////////////////////
// 下面是此练习提供的伪网页接口，用于测试，可不必理会

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls     []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
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

