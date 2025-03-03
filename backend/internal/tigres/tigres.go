package tigres

import (
	"context"
	"fmt"
	"io"

	"github.com/404th/Ink/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/transport/http"
	"go.uber.org/zap"
)

type tigres struct {
	cfg   *config.Config
	sugar *zap.SugaredLogger
}

func NewTigres(cfg *config.Config, sugar *zap.SugaredLogger) TigresI {
	return &tigres{
		cfg:   cfg,
		sugar: sugar,
	}
}

// ------------------- S3 Client Helper Func -------------------
// createS3Client creates and returns a new S3 client with explicit credentials and endpoint.
func createS3Client(cfg *config.Config) (*s3.Client, error) {
	// Load the AWS configuration.
	newAWSCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfg.AWSAccessKeyId, cfg.AWSSecretAccessKey, ""),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	return s3.NewFromConfig(newAWSCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.AWSEndpointUrlS3)
		o.Region = *aws.String(cfg.AWSRegion)
		o.UsePathStyle = false
	}), nil
}

type TigresI interface {
	UploadImage(ctx context.Context, key string, file io.Reader) (string, error)
	UploadVideo(ctx context.Context, key string, file io.Reader) (string, error)
	UploadFile(ctx context.Context, key string, file io.Reader) (string, error)
}

// ------------------- Upload Functions (Using Multipart Upload) -------------------

// UploadImage uploads an image file to S3 using streaming and multipart upload.
func (t *tigres) UploadImage(ctx context.Context, key string, file io.Reader) (string, error) {
	client, err := createS3Client(t.cfg)
	if err != nil {
		return "", fmt.Errorf("failed to create S3 client: %w", err)
	}

	// Create an uploader that will stream the file and perform multipart upload if needed.
	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(t.cfg.AWSImageBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}
	// result.Location is the URL to the uploaded object.
	return result.Location, nil
}

// UploadVideo uploads a video file to S3 using streaming and multipart upload.
func (t *tigres) UploadVideo(ctx context.Context, key string, file io.Reader) (string, error) {
	client, err := createS3Client(t.cfg)
	if err != nil {
		return "", fmt.Errorf("failed to create S3 client: %w", err)
	}

	// Create an uploader that will stream the file and perform multipart upload if needed.
	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(t.cfg.AWSVideoBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}
	// result.Location is the URL to the uploaded object.
	return result.Location, nil
}

// UploadFile uploads a generic file to S3 using streaming and multipart upload.
func (t *tigres) UploadFile(ctx context.Context, key string, file io.Reader) (string, error) {
	client, err := createS3Client(t.cfg)
	if err != nil {
		return "", fmt.Errorf("failed to create S3 client: %w", err)
	}

	// Create an uploader that will stream the file and perform multipart upload if needed.
	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(t.cfg.AWSFileBucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image: %w", err)
	}
	// result.Location is the URL to the uploaded object.
	return result.Location, nil
}

func WithHeader(key, value string) func(*s3.Options) {
	return func(options *s3.Options) {
		options.APIOptions = append(options.APIOptions, http.AddHeaderValue(key, value))
	}
}
