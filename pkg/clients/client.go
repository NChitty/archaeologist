package clients

import (
	"context"
	"net/http"
	"strings"
)

type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestEditorFn func(ctx context.Context, req *http.Request) error

type Client struct {
	Server         string
	Client         HttpRequestDoer
	RequestEditors []RequestEditorFn
}

type ClientOption func(*Client) error

func NewClient(server string, opts ...ClientOption) (*Client, error) {
	client := Client{
		Server: server,
	}

	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

func (c *Client) applyEditors(
	ctx context.Context,
	req *http.Request,
	additionalEditors []RequestEditorFn,
) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

func WithHttpClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}
