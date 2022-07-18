## 介绍

使用`github.com/projectdiscovery/cdncheck`来快速判断是否存在CDN

## 安装编译

### 自动

```shell
go install github.com/damit5/cdnCheck_go@latest
```



### 手动

```shell
git clone https://github.com/damit5/cdnCheck_go
cd cdnCheck_go

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath -o release/cdnCheck_go_linux_amd64
GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o release/cdnCheck_darwin
GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o release/cdnCheck.exe
```

## 用法

```shell
Usage of ./cdnCheck_darwin:
  -nd string
    	无CDN域名保存地址，不保存置空即可
  -ni string
    	无CDN IP保存地址，不保存置空即可
  -t string
    	需要扫描的文件
  -thread int
    	并发数 (default 20)
    	
./cdnCheck_darwin -t ../test/baidu.com.txt -thread 100 -nd ../test/nocdndomain.txt -ni ../test/nocdnip.txt
```

![image-20220718142155174](README.assets/image-20220718142155174.png)

## 疑难杂症

* 建议修改本地NS Server的地址为`8.8.8.8`或者其他大DNS，不然可能因为并发太大导致结果很多请求被拒绝