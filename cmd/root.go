package cmd

import (
	"Dirscan/config"
	"Dirscan/crawler"
	"fmt"
	"github.com/gookit/color"
)

func Run() {
	if config.Url != "" && config.Urlfile == "" && config.Crawler != true{
		//进行单url扫描
		config.Tishi()
		Turl := config.Urll(config.Url)
		if config.FindUrl(Turl) {
			config.Processchecks(Turl)
			if  config.Antirecursion == true {
				color.Red.Printf("\r[+] 进行反递归扫描！\n")
				config.AntiScans(Turl)
			}
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

}