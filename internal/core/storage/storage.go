package storage

import (
	"context"
	"io"

	"cloud.google.com/go/storage"

	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
)

type Client struct {
	client  *storage.Client
	Buckets map[string]*Bucket
}

// New Creates Client that can be used to access several buckets
func New(bucketNames []string) (*Client, error) {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	var buckets map[string]*Bucket
	for _, name := range bucketNames {
		buckets[name] = newBucket(c, name)
	}
	return &Client{c, buckets}, nil
}

func (c *Client) getBucket(bucketName string) (*Bucket, error) {
	b, ok := c.Buckets[bucketName]
	if !ok {
		return nil, pkgerrors.NewError("Bucket with name %s does not exist", nil)
	}
	return b, nil
}

func (c *Client) UploadToBucket(ctx context.Context, bucketName string, fileName string, file io.Reader) (string, error) {
	b, err := c.getBucket(bucketName)
	if err != nil {
		return "", err
	}
	return b.upload(ctx, fileName, file)
}

func (c *Client) GetFromBucket(ctx context.Context, bucketName string, filePath string) (io.Reader, error) {
	b, err := c.getBucket(bucketName)
	if err != nil {
		return nil, err
	}
	return b.get(ctx, filePath)
}
