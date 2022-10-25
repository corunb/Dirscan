package config

import (
	"crypto/tls"
	"fmt"
	"github.com/EDDYCJY/fake-useragent"
	"github.com/go-resty/resty/v2"
	"github.com/gookit/color"
	"strings"
	"sync"
	"time"
)



var Redirect string
var BiaoJi []string
var w sync.WaitGroup


// Scans 单个扫描
func Scans(url string) {

	//状态码判断是否输入错误
	CodeIstrue(Codel(Rcode))


	//读取目录文件
	dic := ReadFile(Pathfile)


	//设置进度条
	var bar Bar
	bar.NewBar(0,len(dic))




	//设置字典管道
	pathChan:= make(chan string, len(dic))

	//生产者
	for _,v := range dic {
		pathChan  <- v
	}
	close(pathChan)



	//设置线程阻塞
	w.Add(Threads)
	//消费者
	for i:=0; i<Threads; i++{
		//go func(url string,pathChan chan string,w sync.WaitGroup) {
			if Requestmode == "G" {
				//fmt.Println("a:",runtime.NumGoroutine())
				go GetScan(url,pathChan,&w,&bar)
			}else if Requestmode == "H" {
				go HeadScan(url,pathChan,&w,&bar)
			}
		//}(url,pathChan,w)
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
			color.Red.Printf("[!] 递归扫描结束\n")
		}
	}

	bar.Close()
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
func HeadScan(url string,pathChan <- chan string,w *sync.WaitGroup,bar *Bar) {
	for path := range pathChan {
		client := resty.New().SetTimeout(time.Duration(Timeout)*time.Second).
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			SetRedirectPolicy(resty.FlexibleRedirectPolicy(20), resty.DomainCheckRedirectPolicy("host1.com", "host2.org", "host3.net"))
		//client.SetTimeout(3 * time.Second)
		//随机UA
		client.SetHeader("User-Agent", browser.Random())
		client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
		//http代理
		//if Httpproxy != "" {
		//	client.SetProxy(Httpproxy)
		//}

		resp, err := client.R().
			Head(url + path)
		if err != nil {
			s1 := strings.Replace(err.Error(), "\": redirect is not allowed as per DomainCheckRedirectPolicy", "", -1)
			s2 := strings.Replace(s1, "Head \"", "--> ", -1)
			Redirect = s2
		}

		respCode := resp.StatusCode()

			//指定状态码排除
			codes := Codel(Rcode)
			nocodes := Codel(Neglect)
			newcodes := difference(codes, nocodes)
			for _, code := range newcodes {
				if respCode == code {
					HeadPrint(respCode, url, path)
				}
			}



		if Recursion == true {
			Recursionchoose(respCode, url, path)
		}
		//进度条计数
		bar.Add(1)

	}
	w.Done()


}

// GetScan Getscan  Get扫描
func GetScan(url string,pathChan <- chan string,w *sync.WaitGroup,bar *Bar)  {
		for path := range pathChan {
			//延迟
			client := resty.New().SetTimeout(time.Duration(Timeout) * time.Millisecond).
				//忽略https证书
				SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
				//302跳转
				SetRedirectPolicy(resty.FlexibleRedirectPolicy(20),
					resty.DomainCheckRedirectPolicy("host1.com", "host2.org", "host3.net"))
			//随机UA
			client.SetHeader("User-Agent", browser.Random())
			//设置Content-Type
			client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
			//http代理
			//if Httpproxy != "" {
			//	client.SetProxy(Httpproxy)
			//}
			//Get请求
			resp, err := client.R().
				Get(url + path)
			if err != nil {
				//302跳转返回值处理
				s1 := strings.Replace(err.Error(), "\": redirect is not allowed as per DomainCheckRedirectPolicy", "", -1)
				s2 := strings.Replace(s1, "Get \"", "--> ", -1)
				Redirect = s2
			}

			//respTime := resp.Time()   //从发送请求到收到响应的时间；
			//respRece :=resp.ReceivedAt() // 接收到响应的时刻
			respCode := resp.StatusCode()
			Bodylen := Storage(len(resp.Body()))

				//指定状态码排除
				codes := Codel(Rcode)
				nocodes := Codel(Neglect)
				newcodes := difference(codes, nocodes)
				for _, code := range newcodes {
					if respCode == code {
						GetPrint(respCode, Bodylen, url, path)

					}
				}


			if Recursion == true {
				Recursionchoose(respCode, url, path)
			}

		//进度条计数
			bar.Add(1)

		}
	// 消费完毕则调用 Done，减少需要等待的线程
	w.Done()

}





