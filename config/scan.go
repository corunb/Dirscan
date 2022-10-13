package config

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"
	"golang.org/x/net/context"
	"golang.org/x/sync/semaphore"
	"strings"
	"sync"
	"time"
)

const (
	Weight = 1 // 每个goroutine获取信号量资源的权重
)

var Redirect string
var BiaoJi []string



// Scans 单个扫描
func Scans(url string) {

	//状态码判断是否输入错误
	CodeIstrue(Codel(Rcode))
	//CodeIstrue(Codel(Neglect))

	//读取目录文件
	dic := ReadFile(Pathfile)


	s := semaphore.NewWeighted(int64(Threads)) //设置线程
	var w sync.WaitGroup

	//并发扫描
	for _,path := range dic {
		w.Add(1)
		go func(url string,path string) {
			_ = s.Acquire(context.Background(), Weight)
			if Requestmode == "G" {
				GetScan(url,path)
			}else if Requestmode == "H" {
				HeadScan(url,path)
			}

			s.Release(Weight)
			w.Done()
		}(url,path)
	}
	w.Wait()

	//递归扫描
	if Recursion == true {
		if len(BiaoJi)> 0 {
			//数组去重
			BiaoJi = RemoveRepByLoop(BiaoJi)
			//fmt.Println(BiaoJi)
			//进行递归扫描
			newurl := Urll(BiaoJi[0])
			//删除数据得第一个元素
			BiaoJi = BiaoJi[1:]
			fmt.Println("")
			color.Green.Printf("target: %v \n",newurl)
			time.Sleep(200 * time.Millisecond)
			Scans(newurl)
		}else {
			color.Red.Println("[!] 递归扫描结束")
		}
	}

}


// Scanes 批量扫描
func Scanes() {
	//读取url列表
	urls := ReadFile(Urlfile)

	//遍历url
	for _,url := range urls {

		Scans(Urll(url))
	}

}

// HeadScan Scan Head扫描
func HeadScan(url string,path string) {

		client := resty.New().SetTimeout(time.Duration(Timeout) * time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)).SetRedirectPolicy(resty.FlexibleRedirectPolicy(20),resty.DomainCheckRedirectPolicy("host1.com", "host2.org", "host3.net"))
		resp, err := client.R().
			Head(url + path)
		if err != nil {
			s1 := strings.Replace(err.Error(), "\": redirect is not allowed as per DomainCheckRedirectPolicy", "", -1)
			s2 := strings.Replace(s1, "Head \"", "--> ", -1)
			Redirect = s2
		}

		respCode := resp.StatusCode()

	if Neglect == "" {
		//状态码筛选
		codes := Codel(Rcode)
		for _, code := range codes {
			if respCode == code {
				HeadPrint(respCode,url,path)
			}
		}
	}else {
		//指定状态码排除
		codes := Codel(Rcode)
		nocodes :=Codel(Neglect)
		newcodes := difference(codes, nocodes)
		for _, code := range newcodes {
			if respCode == code {
				HeadPrint(respCode,url,path)
			}
		}

	}

	if Recursion == true {
		Recursionchoose(respCode,url,path)
	}

}

// GetScan Getscan  Get扫描
func GetScan(url string, path string) {
	client := resty.New().SetTimeout(time.Duration(Timeout) *time.Second).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)).SetRedirectPolicy(resty.FlexibleRedirectPolicy(20), resty.DomainCheckRedirectPolicy("host1.com", "host2.org", "host3.net"))
	resp, err := client.R().
		Get(url + path)
	if err != nil {
		//fmt.Println(err)
		s1 := strings.Replace(err.Error(), "\": redirect is not allowed as per DomainCheckRedirectPolicy", "", -1)
		s2 := strings.Replace(s1, "Get \"", "--> ", -1)
		Redirect = s2
	}

	respCode := resp.StatusCode()
	Bodylen := Storage(len(resp.Body()))


	//fmt.Println(url+path, respCode)
	if Neglect == "" {
		//状态码筛选
		codes := Codel(Rcode)
		for _, code := range codes {
			if respCode == code {
				GetPrint(respCode,Bodylen,url,path)
			}
		}
	}else {
		//指定状态码排除
		codes := Codel(Rcode)
		nocodes :=Codel(Neglect)
		newcodes := difference(codes, nocodes)
		for _, code := range newcodes {
			if respCode == code {
				GetPrint(respCode,Bodylen,url,path)
			}
		}
	}

	if Recursion == true {
		Recursionchoose(respCode,url,path)
	}

}





