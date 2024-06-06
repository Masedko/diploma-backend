package storage

import (
	"context"
	"fmt"
	"io"
	"strings"

	"cloud.google.com/go/storage"
)

const (
	ImageBucket = "diploma-images"
)

type Bucket struct {
	*storage.BucketHandle
	Name string
}

// NewBucket Creates Bucket that can be used to upload and download files to a single bucket
func newBucket(c *storage.Client, bucketName string) *Bucket {
	bkt := c.Bucket(bucketName)
	return &Bucket{bkt, bucketName}
}

func (b *Bucket) upload(ctx context.Context, fileName string, file io.Reader) (string, error) {
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

func (b *Bucket) get(ctx context.Context, filePath string) (file io.Reader, err error) {
	fileName := strings.TrimLeft(filePath, fmt.Sprintf("%s/", b.Name))
	obj := b.Object(fileName)
	rc, err := obj.NewReader(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot create reader: %v", err)
	}
	return rc, nil
}
