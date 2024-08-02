package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/NChitty/artifactsmmo/pkg/schemas/requests"
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
    fmt.Println("Using a prebuilt http client")
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

func AddAuthorizationTokenRequestEditor(ctx context.Context, req *http.Request) error {
	token, isNotEmpty := os.LookupEnv("ARTIFACTS_MMO_TOKEN")
	if isNotEmpty {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
	return errors.New("Token is unset")
}

type ClientInterface interface {
	GetStatus(
		ctx context.Context,
		reqEditors ...RequestEditorFn,
	) (*http.Response, error)
	CharacterActionMove(
		ctx context.Context,
		name string,
		body requests.PositionSchema,
		reqEditors ...RequestEditorFn,
	) (*http.Response, error)
}

func (c *Client) GetStatus(
	ctx context.Context,
	reqEditors ...RequestEditorFn,
) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.Server, nil)
	if err != nil {
		return nil, err
	}
  req.Header.Add("Accepts", "application/json")
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CharacterActionMove(
	ctx context.Context,
	name string,
	body requests.PositionSchema,
	reqEditors ...RequestEditorFn,
) (*http.Response, error) {
	req, err := newCharacterActionMovePostRequest(c.Server, name, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func newCharacterActionMovePostRequest(
	server string,
	name string,
	body requests.PositionSchema,
) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return newCharacterActionMovePostRequestWithBody(server, name, "application/json", bodyReader)
}

func newCharacterActionMovePostRequestWithBody(
	server string,
	name string,
	contentType string,
	bodyReader io.Reader,
) (*http.Request, error) {
	var err error

	serverUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/my/%s/action/move", name)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryUrl, err := serverUrl.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accepts", contentType)

	return req, nil
}
