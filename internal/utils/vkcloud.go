package utils

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"net/http"
)

var s3Client *s3.S3
var bucket string

func InitVKCloudService() {
	region := "kz-ast"
	accessKey := "gkf3GircGx4R8FgoSiKNUn"
	secretKey := "hTG2nkHZYBoGvm1MzHHmhTnTmayP3EYpDFhvo9ZCy3GE"
	bucket = "itfest"

	log.Printf("Initializing VK Cloud Service with region: %s, bucket: %s", region, bucket)
	log.Printf("Access Key: %s", accessKey)
	log.Printf("Secret Key: %s", secretKey)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey, secretKey, ""),
		Endpoint:         aws.String("https://hb.kz-ast.vkcs.cloud"),
		S3ForcePathStyle: aws.Bool(true), // Required for VK Cloud
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	s3Client = s3.New(sess)

	testDNS(bucket, region)
}

func testDNS(bucket, region string) {
	host := "hb.kz-ast.vkcs.cloud"
	resp, err := http.Get(fmt.Sprintf("https://%s", host))
	if err != nil {
		log.Printf("DNS resolution failed for host %s: %v", host, err)
	} else {
		log.Printf("DNS resolution successful for host %s, status: %s", host, resp.Status)
	}
}

func UploadFile(key string, file []byte) (string, error) {
	log.Printf("Bucket: %s", bucket)
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(file),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		log.Printf("Failed to upload file to bucket: %v", err)
		return "", fmt.Errorf("failed to upload file: %v", err)
	}
	return key, nil
}