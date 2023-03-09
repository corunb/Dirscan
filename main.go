package main

import (
	"Dirscan/config"
	"Dirscan/crawler"
	"fmt"
	"github.com/gookit/color"
	"time"
)




func main() {
	start := time.Now()

	if config.Url != "" && config.Urlfile == "" && config.Crawler != true{
		//进行单url扫描
		config.Tishi()
		Turl := config.Urll(config.Url)
		if config.FindUrl(Turl) == true {
			config.Scans(Turl)
		}else {
			//fmt.Println("[!] 目标url无法访问")
			color.Red.Printf("\rtarget: %v  [!] 目标url无法访问\n",Turl)
		}
		//批量扫描
	}else if config.Urlfile != "" && config.Url == ""  && config.Crawler != true{
		config.Tishis()
		config.Scanes()
		//爬虫
	}else if config.Url != "" && config.Urlfile == "" && config.Crawler == true{
		crawler.Crawler(config.Url)
	}else {
		fmt.Println("请输入-h查看帮助！")
	}



	end := time.Since(start)

	color.HiGreen.Printf("\n运行时间为: %v\n", end)
}
