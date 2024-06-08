package storage

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

type Bucket struct {
	*storage.BucketHandle
	Name string
}

// NewBucket Creates Bucket that can be used to upload and download files to a single bucket
func NewBucket(c *Client, bucketName string) *Bucket {
	bkt := c.client.Bucket(bucketName)
	return &Bucket{bkt, bucketName}
}

func (b *Bucket) Upload(ctx context.Context, fileName string, file io.Reader) (string, error) {
	obj := b.Object(fileName)
	wc := obj.NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("cannot copy file: %v", err)
	}
	err := wc.Close()
	if err != nil {
		return "", fmt.Errorf("cannot close file writer: %v", err)
	}
	return fmt.Sprintf("%s/%s", b.Name, fileName), nil
}

func (b *Bucket) Delete(ctx context.Context, fileName string) error {
	obj := b.Object(fileName)
	err := obj.Delete(ctx)
	if err != nil {
		return fmt.Errorf("cannot delete file: %v", err)
	}
	return nil
}
