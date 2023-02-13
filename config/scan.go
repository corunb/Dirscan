package config

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
)

var BiaoJi []string
var w sync.WaitGroup

// Scans 单个扫描
func Scans(Turl string) {

	//状态码判断是否输入错误
	CodeIstrue(Codel(Rcode))

	//读取目录文件
	dic := Typeselection()

	//设置进度条
	var bar Bar
	bar.NewBar(0, len(dic))

	//设置字典管道
	pathChan := make(chan string, len(dic))

	//生产者
	for _, v := range dic {
		pathChan <- v
	}
	close(pathChan)

	//设置线程阻塞
	w.Add(Threads)
	//消费者
	for i := 0; i < Threads; i++ {
		//go func(url string,pathChan chan string,w sync.WaitGroup) {
		if Requestmode == "GET" {
			//fmt.Println("a:",runtime.NumGoroutine())
			go GetScan(Turl, pathChan, &w, &bar)
		} else if Requestmode == "HEAD" {
			go HeadScan(Turl, pathChan, &w, &bar)
		}
		//}(url,pathChan,w)
	}
	w.Wait()

	//递归扫描
	if Recursion == true {
		if len(BiaoJi) > 0 {
			//数组去重
			BiaoJi = RemoveRepByLoop(BiaoJi)
			//fmt.Println(BiaoJi)
			//进行递归扫描
			newurl := Urll(BiaoJi[0])
			//删除数据得第一个元素
			BiaoJi = BiaoJi[1:]
			fmt.Println(" ")
			color.Green.Printf("target: %v \n", newurl)
			time.Sleep(200 * time.Millisecond)
			Scans(newurl)
		} else {
			color.Red.Printf("\n[!] 递归扫描结束")
		}
	}
	//bar.Close()
}

// Scanes 批量扫描
func Scanes() {
	//读取url列表
	urls := ReadFile(Urlfile)

	//遍历url
	for _, surl := range urls {
		Turl := Urll(surl)
		if FindUrl(Turl) == true {
			fmt.Printf("\rtarget: %v\n",Turl)
			Scans(Turl)
		}else {
			fmt.Printf("\rtarget: %v\n",Turl)
		}
	}

}

// HeadScan Scan Head扫描
func HeadScan(Turl string, pathChan <-chan string, w *sync.WaitGroup, bar *Bar) {
	for path := range pathChan {
		resp := Request(Requestmode, Turl, path)
		if resp != nil {
			Rurl := resp.Header.Get("location") //获取302跳转的url

			respCode := resp.StatusCode //状态码

			//指定状态码排除
			codes := Codel(Rcode)
			nocodes := Codel(Neglect)
			newcodes := difference(codes, nocodes)
			for _, code := range newcodes {
				if respCode == code {
					HeadPrint(respCode, Turl, path, Rurl)
				}
			}

			if Recursion == true {
				Recursionchoose(respCode, Turl, path)
			}
		}
		//进度条计数
		bar.Add(1)

	}
	w.Done()

}

// GetScan Getscan  Get扫描
func GetScan(Turl string, pathChan <-chan string, w *sync.WaitGroup, bar *Bar) {
	for path := range pathChan {
		resp := Request(Requestmode, Turl, path)
		if resp != nil {
			Rurl := resp.Header.Get("location") //获取302跳转的url

			body, _ := ioutil.ReadAll(resp.Body)
			Bodylen := Storage(len(body)) //返回长度
			//fmt.Println(string(body))
			respCode := resp.StatusCode //状态码

			//指定状态码排除
			codes := Codel(Rcode)
			nocodes := Codel(Neglect)
			newcodes := difference(codes, nocodes)
			for _, code := range newcodes {
				if respCode == code {
					GetPrint(respCode, Bodylen, Turl, path, Rurl)
				}
			}

			if Recursion == true {
				Recursionchoose(respCode, Turl, path)
			}

		}
		//进度条计数
		bar.Add(1)
	}
	// 消费完毕则调用 Done，减少需要等待的线程
	w.Done()
}

// Request 封装请求
func Request(Requestmode string, Turl string, path string) *http.Response {

	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, //忽略http证书
		IdleConnTimeout:     5 * time.Second,
		MaxConnsPerHost:     5,  //每个host可以发出的最大连接个数（包括长链接和短链接），是非常有用的参数，可以用来对socket数量进行限制
		MaxIdleConns:        0,  //host的最大链接数量,0表示无限大
		MaxIdleConnsPerHost: 10, //连接池对每个host的最大链接数量
	}
	if Proxy != "" {
		u, _ := url.Parse(Proxy)
		if strings.ToLower(u.Scheme) != "socks5" { //判断前缀走什么代理
			tr.Proxy = http.ProxyURL(u) //代理 URL 返回一个代理函数（用于传输），该函数始终返回相同的 URL。
		} else {
			dialer, err := Socks5Dailer(u.String()) //获取dialer，走socks5代理
			tr.Dial = dialer.Dial                   //传入socks5参数
			if err != nil {
				fmt.Println("error:", err)
			}
		}
	}

	//设置客户端
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(Timeout) * time.Second, //等待时间
		CheckRedirect: func(req *http.Request, via []*http.Request) error { //获取302
			return http.ErrUseLastResponse
		},
	}
	//设置请求
	//Turlpath :=  url.QueryEscape(Turl +path)
	Turlpath := Turl + path
	req, err := http.NewRequest(Requestmode, Turlpath, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", Uas) //设置随机UA头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if Cookie != "null" {
		req.Header.Set("Cookie", Cookie) //设置cookie
	}

	//发送请求
	resp, _ := client.Do(req)
	//fmt.Println(resp)
	return resp

}
