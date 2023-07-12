package s3pkg

import (
	"context"
	"mime/multipart"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// CreateBucket creates a bucket
func CreateBucket(bucketName, region string, client *s3.Client, ctx context.Context) (*s3.CreateBucketOutput, error) {
	// Create the S3 Bucket
	return client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
}

// GetBucket determines whether we have this bucket
func GetBucket(bucketName string, client *s3.Client, ctx context.Context) (*s3.HeadBucketOutput, error) {
	// Do we have this Bucket
	return client.HeadBucket(ctx, &s3.HeadBucketInput{Bucket: aws.String(bucketName)})
}

// PutObject puts object to s3
func PutObject(bucketName, fileName string, file multipart.File, client *s3.Client, ctx context.Context) (*s3.PutObjectOutput, error) {
	return client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
}

// PutObject puts object to s3
func PutObjectWG(bucketName, fileName string, file multipart.File, wg *sync.WaitGroup, client *s3.Client, ctx context.Context) (*s3.PutObjectOutput, error) {
	defer func() {
		wg.Done()
	}()
	return client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
}

// GetObject retrieve object from s3
func GetObject(fileName string, file multipart.File, bucketName, objectKey string, client *s3.Client, ctx context.Context) (*s3.GetObjectOutput, error) {
	return client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
}
