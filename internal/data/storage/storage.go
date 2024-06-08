package storage

import (
	"context"

	"cloud.google.com/go/storage"
)

type Client struct {
	client *storage.Client
}

// New Creates Client that can be used to access several buckets
func New() (*Client, error) {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}
