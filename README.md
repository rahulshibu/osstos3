

# OssToS3

OssToS3 is a free data transfer tool for transfering file and folders from [OSS](https://www.alibabacloud.com/product/oss) (Alibaba Cloud Object Storage Service) to [S3](https://aws.amazon.com/s3/) (Amazon Simple Storage Service). This package or tool is build on _[Go](https://golang.org/)_ ( _[Go](https://golang.org/)_ is a statically typed, compiled programming language designed at Google). The main advantage is that it runs in windows/linux/mac without extra dependencies

## Setting Up And Running
- Clone or download the project from the github repository 
- Set up the s3 and oss configuration in the `config.toml` in the project root folder
```toml
# This is a TOML document.
[oss]
endpoint = "oss-ap-xxxx-5.aliyuncs.com" 
accesskeyid = "XxXXXXXxXXXxxxXx" 
secretaccesskey = "XxXNxXxNxXXxxXxxxxxxxxXXXLXxxX" 
bucket = "bucket-name" 

[s3]
region = "ap-xxxx-1" 
accesskeyid = "XXXNXXXNXXNNXXN" 
secretaccesskey = "XNXNxxNXnXnnXNXNXnxnnnxXnXnX" 
bucket = "bucket-name" 
token=""
```
- ##### macOS/Linux:
In case of macOS/Linux,open terminal and go to the project-directory. Inside the project directory there will be an executable file `osstos3`.Execute `osstos3` after configuring the `config.toml`. If there is any error in the configuration it will be be printed in the terminal.
```
$ cd project-directory
$ ./osstos3
```
- ##### Windows:
In case of windows,open cmd and go to the project-directory. Inside the project directory there will be an executable file `osstos3.exe`.Execute `osstos3.exe` after configuring the `config.toml`. If there is any error in the configuration it will be be printed in the terminal.
```
cd C:\download-path\project-directory
start osstos3.exe
```
If you have already installed GO in your system, you can also run by specifying the `go run`  command. All you have to do is go the project directory and and follow the below commands.Then `go run` command execute the main function in the `osstos3.go` file. 
```
$ cd project-directory
$ go run osstos3.go
```


## Configuring Go and Taking Build for Different Operating System

The GO executable can be create not only for Mac,Linux and Windows but also for other platforms like android,darwin,freebsd and many more.You can visit the offical GO documentation link by clicking _[here](https://golang.org/)_ and setup for you own system. The documentation by  _[digital ocean](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)_ have also mentioned easy methods for building and installing Go programs for different platform.

## License

