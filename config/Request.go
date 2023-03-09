package config

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request 封装请求
func Request(Targeturl string) *http.Response {

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
	req, err := http.NewRequest(Requestmode, Targeturl, nil)
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

// GETRequest 用于单次结果内容判断
func GETRequest(Targeturl string) *http.Response {

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
	req, err := http.NewRequest("GET", Targeturl, nil)
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

// Socks5Dailer 代理设置
func Socks5Dailer(Socks5proxy string) (proxy.Dialer, error) {
	u, err := url.Parse(Socks5proxy)
	if err != nil {
		return nil, err
	}
	//if strings.ToLower(u.Scheme) != "socks5" { //strings.ToLower  进行全部小写
	//	return nil, errors.New("Only support socks5")
	//}
	address := u.Host //取host
	var auth proxy.Auth
	var dailer proxy.Dialer
	if u.User.String() != "" {
		auth = proxy.Auth{}
		auth.User = u.User.Username()    //取用户名
		password, _ := u.User.Password() //取密码
		auth.Password = password
		dailer, err = proxy.SOCKS5("tcp", address, &auth, proxy.Direct) //设置用户名密码的socks5
	} else {
		dailer, err = proxy.SOCKS5("tcp", address, nil, proxy.Direct) //没有设置用户密码的socks5
	}

	if err != nil {
		return nil, err
	}
	return dailer, nil
}
