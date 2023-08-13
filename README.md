## 0x01 前言

```
Dirscan 是一款由 go 编写的高并发的目录扫描器，现在已经支持基础扫描功能,后续努力实现更多功能；
大佬们勿喷，有什么建议或者 bug 可以提 issues 。
欢迎一起讨论，提供思路或者字典。
```

### Corun项目成员： [getin63](https://github.com/getin63) & [goout2](https://github.com/goout2)   



### 新增功能变化：

[+] 新增反向递归扫描：

可反向对每级目录进行扫描；

```
./Dirscan -u http://xxx.xxx.xxx.xxx/aaa/bbb/ccc/ddd/eee -Ar
```

![image-20230626100721537](image//image-20230626100721537.png)




* 已实现功能：

```
Version--> V1.5.2 :
1. 支持并发快速扫描，可设置线程，延迟时间；
2. 过滤状态码；
3. 设置字典；
4. 批量扫描；
5. 结果文件写入，只保存 20x、30x、403 的扫描结果；
6. 302 跳转结果显示；
7. 递归扫描，支持 GET/HEAD ；
8. 结果文件默认为 日期+url 命名；
9. 使用加载配置设置默认参数；
10. 状态码排除；
11. 随机UA头
12. 进度条显示
13. 代理支持
14. 扫描前验证存活
15. 指定类型扫描
16. 指定Cookie
17. 添加爬虫爬取页面url
18. 对扫描的结果进行综合判断，提高准确性
19, 进行随机代理
20，添加反向递归扫描
```

* 计划实现：

```
1. 设置代理 ✔️
2. 设置忽略状态码  ✔️
3. 添加爬虫爬取页面url ✔️
4. 进度条显示 ✔️
5. 图形化显示
6. 随机UA头 ✔️
7. 反递归扫描✔️
```


## 0x02 编译

```
linux/mac: 
1. 终端翻墙，同步：go mod tidy
2. 一键编译4个版本的二进制文件，编译：make all 

windows: 
1. 终端翻墙，同步：go mod tidy
2. 编译：CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w " -trimpath -o Dirscan.exe main.go
```


## 0x03 使用方式

### 3.1 扫描参数

```
  -Ar  	进行反递归扫描
  -C    进行爬虫
  -Cs		进行每5秒检查网站存活，默认是开的
  -R string
        指定Get扫描还是Head扫描 (default "GET")
  -T int
        设置线程，默认30 (default 30)
  -c string
        设置Cookie，默认不加cookie (default "null")
  -ei string
        忽略指定状态码,示例：200,403,404,500或者200-400 (default "404")
  -f string
        指定目录字典 (default "./dic/dicc.txt")
  -i string
        筛选指定状态码,示例：200,403,404,500或者200-400 (default "100-403,405-599")
  -m string
        根据指定类型进行扫描，可设置php,asp,aspx,jsp
  -o string
        保存扫描结果,默认输出日期+地址
  -p string
        proxy，可设置http代理或socks5代理，socks5://admin:corun@x.x.x.x:1080
  -pf string
        指定Proxy列表,进行随机切换
  -r    进行递归扫描
  -t int
        设置延时时间，默认5s (default 5)
  -u string
        指定url
  -uf string
        指定url列表
```

### 3.2 参数详解

#### GET 扫描:

```
./Dirscan -u http://xxx.xxx.xxx.xxx 
默认以 日期+url 进行结果记录的文件名。
```

#### HEAD 扫描：

```
./Dirscan -u http://xxx.xxx.xxx.xxx -R HEAD
```

#### 筛选状态码：

```
./Dirscan -u http://xxx.xxx.xxx.xxx -i 200,302
筛选状态码可支持单个/多个，和区间筛选，区间可为 -i 200-302
```

#### 递归扫描：

```
./Dirscan -u http://xxx.xxx.xxx.xxx -r -i 200 
默认对状态码200、301、302、403的目录进行递归扫描。
```

#### 线程控制：

```
默认 30 线程，5 秒延迟
./Dirscan -u http://xxx.xxx.xxx.xxx -T 100 -t 10
```

#### 状态码排除

```
使用 -ei 301-401,404 可排除状态码显示。
```

#### 代理设置

```
http代理：
./Dirscan -u http://xxx.xxx.xxx.xxx -p http://127.0.0.1:8080

socks5代理：
./Dirscan -u http://xxx.xxx.xxx.xxx -p socks5://127.0.0.1:7891
./Dirscan -u http://xxx.xxx.xxx.xxx -p socks5://admin:corun@xx.xx.xx.xx:1080
```

#### 指定类型扫描

```
./Dirscan -u http://xxx.xxx.xxx.xxx -m php
可指定类型进行扫描，会将字典中的可变字典转换为指定类型的字典。
```

#### 字典选择：

```
./Dirscan -u http://xxx.xxx.xxx.xxx -f ./dic/php.txt
```



#### 默认配置加载

```
可更改 ./default/default.ini 文件中的默认配置信息，以改变默认设置。
```

#### 爬虫：

爬虫爬取页面url（原本想爬取的url进行递归扫描。但是发现爬取的url可能会很多，所以只增加单独的爬取功能，可以自行对爬取的url进行二次扫描）

```
./Dirscan -u http://xxx.xxx.xxx.xxx -C
```

#### 随机切换代理：

每隔2秒随机切换代理，可用于绕过waf或封ip。

```
./Dirscan -u  http://xxxxxxxxxxxxxxxxxxx -pf ./proxy.txt -T 5
使用代理时，线程最好不要太高
```

不进行随机代理切换时，每隔5秒验证目标存活，第一时间发现存在waf或网站访问不通。

```
可使用-Cs关闭探存活功能：
./Dirscan -u  http://xxxxxxxxxxxxxxxxxxx -Cs
```

![image-20230320100018765](image//image-20230320100018765.png)

#### 反向递归扫描

可反向对每级目录进行扫描；

```
./Dirscan -u http://xxx.xxx.xxx.xxx/aaa/bbb/ccc/ddd/eee -Ar
```

![image-20230626100721537](image//image-20230626100721537.png)



## 0x04 更新记录：

* [!] 2022.10.13 修复递归扫描时，302 与 200 时的字典重复扫描；
* [!] 2022.10.18 修复没有默认配置文件时退出并生成默认配置文件；
* [+] 2022.10.23 重构并发方式；
* [+] 2022.10.24 新增扫描时随机UA头；
* [+] 2022.10.25 新增进度条显示(windows,linux测试均可使用。)；
* [!] 2022.11.9 修复递归扫描已知问题
* [+] 2022.11.9 重构http请求，使用原生net/http包；新增代理支持，支持http代理，socks5代理；
* [+] 2022.11.9 新增扫描前探存活；
* [+] 2022.11.10 新增指定类型扫描；
* [+] 2022.11.14 新增指定Cookie; 
* [!] 2023.2.7 修复批量探测存活问题;
* [!] 2023.2.7 修复扫描结果显示200结果的大小；
* [!] 2023.3.1 修复批量探测存活显示错误问题;
* [+] 2023.3.2 新增爬虫爬取页面url；
* [+] 2023.3.8 新增对302页面的网站进行综合判断
* [+] 2023.3.9 新增针对扫描结果的200,3xx,4xx,5xx页面进行状态码，长度，内容综合判断，排除无效的页面，不用在众多400,500中找200的结果，提高目录扫描准确性。
* [!] 2023.3.9 修复HEAD扫描时，结果判断不生效。
* [!] 2023.3.13 修复已知逻辑问题
* [+] 2023.3.20 新增每隔2秒随机切换代理
* [+] 2023.3.20 新增不进行随机代理切换时，每隔5秒验证目标存活
* [!] 2023.3.21 修复已知bug
* [+] 2023.6.26 新增反向递归扫描
* [!] 2023.6.26 新增探存活开关
* [!] 2023.8.13 修复递归扫描已知bug






## 0x05 声明：

* 本工具仅用于个人安全研究学习。由于传播、利用本工具而造成的任何直接或者间接的后果及损失，均由使用者本人负责，工具作者不为此承担任何责任。

* 转载请注明来源！！！！
