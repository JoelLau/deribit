package jsonrpchttp

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"os"

	"github.com/JoelLau/deribit/pkgs/http"
)

// "low" level client that operates on MessageRequests and MessagesResponses
type JsonRpcClient struct {
	Http   Httper
	config ClientConfig
	loggr  Logger // avoid direct usage, use Logger() instead
}

type Logger interface {
	Info(string, ...any)
	Debug(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
}

// Craft, Perform HTTP requests
type Httper interface {
	Get(url string, opts ...http.OptFunc) (res http.Response, err error)
	Post(url string, body string, opt ...http.OptFunc) (res http.Response, err error)
}

type ClientConfig struct {
	// Base url to connect to:
	// - "test.deribit.net" for testnet, and
	// - "deribit.net"
	Host string

	// Whether or not to start logging
	Verbose bool
}

func NewJsonRpcClient(config ClientConfig) (c *JsonRpcClient, err error) {
	c = &JsonRpcClient{}
	c.Http = http.NewHttpClient()
	c.config = config
	c.loggr = c.Logger()

	return
}

func (c *JsonRpcClient) Logger() Logger {
	if c.loggr != nil {
		return c.loggr
	}

	opts := &slog.HandlerOptions{Level: slog.Level(99)}
	if c.config.Verbose {
		opts.Level = slog.LevelDebug
	}
	c.loggr = slog.New(slog.NewTextHandler(os.Stderr, opts))

	return c.loggr
}

func (c *JsonRpcClient) Get(request MessageRequest) (res MessageResponse, err error) {
	u, err := c.Url(request)
	if err != nil {
		err = fmt.Errorf("error deriving URL from request: %w", err)
		c.Logger().Error(err.Error(), slog.Any("request", request))
		return
	}

	rawRes, rawErr := c.Http.Get(u.String())
	if rawErr != nil {
		err = fmt.Errorf("error while making GET request: %w", rawErr)
		c.Logger().Error(err.Error(), slog.Any("request", request), slog.Any("url", u.String()))
		return
	}

	rawBytes, err := io.ReadAll(rawRes.Body)
	if err != nil {
		err = fmt.Errorf("error reading response bytes: %w", err)
		c.Logger().Error(err.Error(), slog.Any("url", u.String()), slog.Any("response", rawRes))
		return
	}

	if err = json.Unmarshal(rawBytes, &res); err != nil {
		err = fmt.Errorf("error parsing message request: %w", err)
		c.Logger().Error(err.Error(), slog.Any("rawBytes", string(rawBytes)))
		return
	}

	return
}

func (c *JsonRpcClient) Post(request MessageRequest) (res MessageResponse, err error) {
	u, err := c.Url(request)
	if err != nil {
		err = fmt.Errorf("error deriving URL from request: %w", err)
		c.Logger().Error(err.Error(), slog.Any("request", request))
		return
	}

	rawRes, rawErr := c.Http.Post(u.String(), "")
	if rawErr != nil {
		err = fmt.Errorf("error while making POST request: %w", rawErr)
		c.Logger().Error(err.Error(), slog.Any("request", request), slog.Any("url", u.String()))
		return
	}

	fmt.Printf("%+v", rawRes)
	return
}

func (c *JsonRpcClient) Url(request MessageRequest) (u url.URL, err error) {
	u = url.URL{}
	u.Scheme = "https"
	u.Host = c.config.Host
	u.Path = fmt.Sprintf("api/v2/%s", request.Method)

	return
}
