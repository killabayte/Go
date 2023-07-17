package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const bucketName = "aws-golang-sdk-test-20231707"
const regionName = "eu-central-1"

func main() {
	var (
		s3Client *s3.Client
		err      error
	)
	ctx := context.Background()
	if s3Client, err = initS3Client(ctx); err != nil {
		fmt.Printf("Init S3Client error: %s", err)
		os.Exit(1)
	}
	if err = createS3Bucket(ctx, s3Client); err != nil {
		fmt.Printf("Create S3Client error: %s", err)
		os.Exit(1)
	}
	if err = uploadToS3Bucket(ctx, s3Client); err != nil {
		fmt.Printf("Upload S3Client error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Upload complete\n")
}

func initS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(regionName))
	if err != nil {
		return nil, fmt.Errorf("Unable to load SDK config, %s", err)
	}

	return s3.NewFromConfig(cfg), nil
}

func createS3Bucket(ctx context.Context, s3Client *s3.Client) error {
	allBuckets, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("ListBuckets error: %s", err)
	}

	found := false
	for _, bucket := range allBuckets.Buckets {
		if *bucket.Name == bucketName {
			found = true
		}
	}
	if !found {
		_, err := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: regionName,
			},
		})
		if err != nil {
			return fmt.Errorf("CreateBucket error: %s", err)
		}
	}
	return nil
}

// func uploadToS3Bucket(ctx context.Context, s3Client *s3.Client) error {
// 	uploader := manager.NewUploader(s3Client)
// 	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String("upload.test"),
// 		Body:   strings.NewReader("This is upload test from golang sdk"),
// 	})
// 	if err != nil {
// 		return fmt.Errorf("Upload error: %s", err)
// 	}
// 	return nil
// }

func uploadToS3Bucket(ctx context.Context, s3Client *s3.Client) error {
	testFile, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return fmt.Errorf("Readfile error: %s", err)
	}
	uploader := manager.NewUploader(s3Client)
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("upload.test"),
		Body:   bytes.NewReader(testFile),
	})
	if err != nil {
		return fmt.Errorf("Upload error: %s", err)
	}
	return nil
}
