package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/EDDYCJY/gsema"
	"github.com/projectdiscovery/cdncheck"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
)

var client *cdncheck.Client
var err error
var semaphore *gsema.Semaphore
var cdnResult []string	// 保存CDN的域名
var nonCdnDomainResult []string // 保存无CDN的域名
var nonCdnIPResult []string	// 保存无CDN的IP

// 传入的参数
var target string
var nonCdnDomainSavePath string
var nonCdnIPSavePath string
var threads int

func checkCDN(domain string) {
	defer semaphore.Done()
	// 域名处理，可能传入是URL
	/*
	非正常格式：需要先变成标准格式 www.baidu.com
	https://www.baidu.com
	1.2.3.4:80
	www.baidu.com:443
	 */
	if !strings.HasPrefix(domain, "http") {
		domain = "http://" + domain
	}

	parse, err := url.Parse(domain)
	if err != nil {
		return
	}
	domain = parse.Hostname()

	// 域名解析成IP，可能有多个IP
	ips, err := net.LookupIP(domain)
	if err != nil {
		log.Println(domain, err)
		return
	}

	// 进行CDN验证
	for _, ip := range ips {
		// 过滤ipv6
		if strings.Contains(ip.String(), ":") {
			continue
		}

		if found, provider, err := client.Check(ip); found && err == nil {
			log.Println(fmt.Sprintf("%s ==> %s is part of %s cdn", domain, ip.String(), provider))

			cdnResult = append(cdnResult, domain)
		} else {
			log.Println(fmt.Sprintf("%s ==> %s has no cdn", domain, ip.String()))

			nonCdnDomainResult = append(nonCdnDomainResult, domain)
			nonCdnIPResult = append(nonCdnIPResult, ip.String())
		}
	}
}

// 数组去重
func removeRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func flagInit() {
	flag.StringVar(&target, "t", "", "需要扫描的文件")
	flag.StringVar(&nonCdnDomainSavePath, "nd", "", "无CDN域名保存地址，不保存置空即可")
	flag.StringVar(&nonCdnIPSavePath, "ni", "", "无CDN IP保存地址，不保存置空即可")
	flag.IntVar(&threads, "thread", 20, "并发数")
	flag.Parse()
}

func main() {
	flagInit()
	if target == "" {
		flag.Usage()
		return
	}
	semaphore = gsema.NewSemaphore(threads)
	client, err = cdncheck.NewWithCache()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(target)
	defer file.Close()
	if err != nil {
		log.Panicln(err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		semaphore.Add(1)
		go checkCDN(scanner.Text())
	}
	semaphore.Wait()

	if nonCdnDomainSavePath != "" {	// 写入无CDN的domain
		openFile, _ := os.OpenFile(nonCdnDomainSavePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
		defer openFile.Close()
		write := bufio.NewWriter(openFile)
		for _, domain := range removeRepeatedElement(nonCdnDomainResult) {
			write.WriteString(domain + "\n")
		}
		write.Flush()
	}

	if nonCdnIPSavePath != "" {	// 写入无CDN的domain
		openFile, _ := os.OpenFile(nonCdnIPSavePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
		defer openFile.Close()
		write := bufio.NewWriter(openFile)
		for _, ip := range removeRepeatedElement(nonCdnIPResult) {
			write.WriteString(ip + "\n")
		}
		write.Flush()
	}
}