package crawler

import (
	"Dirscan/config"
	"github.com/gookit/color"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Crawler(Url string)  {
	resp, _ := http.Get(Url)
	links := All(resp.Body)
	var jishu int
	//var outcome  string
	for _, v := range links {
		absolute := urlJoin(v,Url)
		if absolute != "" {
			//fmt.Println("parse Urll", absolute)
			jishu++
			color.Green.Printf("\rresults: %v \n",absolute)
			config.Write("[**] results: "+absolute,Url)
			//outcome = absolute
		}
	}
	color.Green.Printf("\r成功爬取 %v 条url\n",jishu)


}

func urlJoin(href ,base string) string {
	uri, _ := url.Parse(href)
	baseUrl, _ := url.Parse(base)
	newurl := baseUrl.ResolveReference(uri).String()
	newUrl ,_ := url.Parse(newurl)
	var outcome string
	if newUrl.IsAbs() == true {
		a := newUrl.String()
		if a != "javascript:;"{
			outcome = a
		}

	}
	return outcome
}


func All(httpBody io.Reader) []string {
	var links []string
	var col []string
	page := html.NewTokenizer(httpBody)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					tl := trimHash(attr.Val)
					col = append(col, tl)
					resolv(&links, col)
				}
			}
		}
	}
}

// trimHash slices a hash # from the link
func trimHash(l string) string {
	if strings.Contains(l, "#") {
		var index int
		for n, str := range l {
			if strconv.QuoteRune(str) == "'#'" {
				index = n
				break
			}
		}
		return l[:index]
	}
	return l
}

// check looks to see if a url exits in the slice.
func check(sl []string, s string) bool {
	var check bool
	for _, str := range sl {
		if str == s {
			check = true
			break
		}
	}
	return check
}

// resolv adds links to the link slice and insures that there is no repetition
// in our collection.
func resolv(sl *[]string, ml []string) {
	for _, str := range ml {
		if check(*sl, str) == false {
			*sl = append(*sl, str)
		}
	}
}