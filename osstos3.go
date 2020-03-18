package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	ossservice "github.com/rahulshibu/oss-service"
	"github.com/rahulshibu/s3service"
)

type (
	//tomlConfig is used for encoding the s3 and oss paramters provided in the config.toml in the root path.
	tomlConfig struct {
		S3  s3config
		OSS ossconfig
	}
	//ossconfig struct define the configurations needed for the Alibaba Cloud Object Storage Service (OSS).
	ossconfig struct {
		Endpoint        string //Specifies the OSS Region or Endpoint to send the request to.
		AccessKeyID     string //Specifies an OSS access key associated with a user or role.
		SecretAccessKey string //Specifies the secret key associated with the access key. This is essentially the "password" for the access key.
		Bucket          string //Specifies the place where object stored in Alibaba Cloud oss
	}
	//s3config struct define the configurations needed for the Amazon Simple Storage Service(OSS)
	s3config struct {
		Region          string //Specifies the AWS Region to send the request to.
		AccessKeyID     string //Specifies an AWS access key associated with an IAM user or role.
		SecretAccessKey string //Specifies the secret key associated with the access key. This is essentially the "password" for the access key.
		Bucket          string //Specifies the place where object stored in Amazon S3
		Token           string //Specifies the session token value that is required if you are using temporary security credentials that you retrieved directly from AWS STS operations
	}
)

var (
	//S3service variable stores the the s3service paramters
	S3service s3service.S3Service
	//Osservice variable stores the the ossservice paramters
	Osservice ossservice.OssService
	//config variable global stores config for the s3 and oss credentials
	config tomlConfig
)

func main() {
	//loading the config file from config.toml from the same directory
	loadAppConfig()

	//creating objects of s3 and oss
	S3service.Region = config.S3.Region
	S3service.AccessKeyID = config.S3.AccessKeyID
	S3service.SecretAccessKey = config.S3.SecretAccessKey
	S3service.Bucket = config.S3.Bucket
	S3service.Token = config.S3.Token
	Osservice.Endpoint = config.OSS.Endpoint
	Osservice.AccessKeyID = config.OSS.AccessKeyID
	Osservice.SecretAccessKey = config.OSS.SecretAccessKey
	Osservice.Bucket = config.OSS.Bucket

	//Specify the marker to return from a certain one
	marker := oss.Marker("")
	ossBucket := Osservice.GetOssBucket()

	//Specify max key and list all buckets objects with paging, return 1000 items each time.
	for {
		lor, err := ossBucket.ListObjects(oss.MaxKeys(1000), marker)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		marker = oss.Marker(lor.NextMarker)
		err = getObjectsAndUpload(lor)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if !lor.IsTruncated {
			break
		}
	}
	fmt.Println("All files copied ")
	os.Exit(1)

}

//getObjectsAndUpload get the each object name
func getObjectsAndUpload(lor oss.ListObjectsResult) error {

	for _, object := range lor.Objects {
		fmt.Println("Copying " + object.Key + " ...")
		//Gets the objects in the bucket as bytes
		buf, err := Osservice.GetByteObject(object.Key)
		if err != nil {
			return err
		}
		//Uploads the buffer byte objects to the Amazon S3 bucket to the specified path or object key
		_, err = S3service.UploadAsBuffer(buf, object.Key)
		if err != nil {
			return err
		}
		fmt.Println("Copied " + object.Key)
	}
	return nil
}

//loads the toml file in the base path
func loadAppConfig() tomlConfig {
	if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
		log.Fatalf(" %s", err)
		panic(fmt.Sprintf("%s", err))
	}
	return config
}
