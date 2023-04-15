package utils

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
)

const (
	DefaultRTMPort   = 1935
	DefaultHTTPPort  = 80
	DefaultHTTPSPort = 443
	DefaultRTSPPort  = 554
	DefaultRTMPSPort = 443
	DefaultRTSPSPort = 322
)

type URLContext struct {
	URL string

	Scheme       string
	Username     string
	Password     string
	StdHost      string
	HostWithPort string
	Host         string
	Port         int
}

func ParseRTSPURL(rawURL string) (ctx URLContext, err error) {
	ctx, err = ParseURL(rawURL, -1)
	if err != nil {
		return
	}
	if (ctx.Scheme != "rtsp" && ctx.Scheme != "rtsps") || ctx.Host == "" {
		return ctx, fmt.Errorf("%w. url=%s", ErrInvalidUrl, rawURL)
	}
	return
}

// ParseURL
// @param defaultPort:
// 如果rawURL指定了端口，则该参数不生效
// 如果设置为-1，内部依然会对常见协议设置默认端口
func ParseURL(rawURL string, defaultPort int) (ctx URLContext, err error) {
	ctx.URL = rawURL

	stdURL, err := url.Parse(rawURL)
	if err != nil {
		return ctx, err
	}
	if stdURL.Scheme == "" {
		return ctx, fmt.Errorf("%w. url=%s", ErrInvalidUrl, rawURL)
	}
	if defaultPort == -1 {
		switch stdURL.Scheme {
		case "http":
			defaultPort = DefaultHTTPPort
		case "https":
			defaultPort = DefaultHTTPSPort
		case "rtmp":
			defaultPort = DefaultRTMPort
		case "rtsp":
			defaultPort = DefaultRTSPPort
		case "rtmps":
			defaultPort = DefaultRTMPSPort
		case "rtsps":
			defaultPort = DefaultRTSPSPort
		}
	}

	ctx.Scheme = stdURL.Scheme
	ctx.StdHost = stdURL.Host
	ctx.Username = stdURL.User.Username()
	ctx.Password, _ = stdURL.User.Password()

	h, p, err := net.SplitHostPort(stdURL.Host)
	if err != nil {
		ctx.Host = stdURL.Host
		if defaultPort == -1 {
			ctx.HostWithPort = stdURL.Host
		} else {
			ctx.HostWithPort = net.JoinHostPort(stdURL.Host, fmt.Sprintf("%d", defaultPort))
			ctx.Port = defaultPort
		}
	} else {
		ctx.Port, err = strconv.Atoi(p)
		if err != nil {
			return ctx, err
		}
		ctx.Host = h
		ctx.HostWithPort = stdURL.Host
	}
	return ctx, nil
}
