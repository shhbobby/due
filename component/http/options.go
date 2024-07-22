package http

import (
	"github.com/dobyte/due/v2/etc"
	"github.com/dobyte/due/v2/registry"
	"github.com/dobyte/due/v2/transport"
	"github.com/gin-gonic/gin"
)

const (
	defaultName = "http" // 默认HTTP服务名称
	defaultAddr = ":0"   // 监听地址
)

const (
	defaultNameKey     = "etc.http.name"
	defaultAddrKey     = "etc.http.addr"
	defaultKeyFileKey  = "etc.http.keyFile"
	defaultCertFileKey = "etc.http.certFile"
)

type Option func(o *options)

type RouteHandler func(engine *gin.Engine)

type options struct {
	name        string                // HTTP服务名称
	addr        string                // 监听地址
	certFile    string                // 证书文件
	keyFile     string                // 秘钥文件
	handler     RouteHandler          // 路由处理器
	registry    registry.Registry     // 服务注册器
	transporter transport.Transporter // 消息传输器
}

func defaultOptions() *options {
	opts := &options{
		name:     defaultName,
		addr:     defaultAddr,
		keyFile:  etc.Get(defaultKeyFileKey).String(),
		certFile: etc.Get(defaultCertFileKey).String(),
	}

	if name := etc.Get(defaultNameKey).String(); name != "" {
		opts.name = name
	}

	if addr := etc.Get(defaultAddrKey).String(); addr != "" {
		opts.addr = addr
	}

	return opts
}

// WithName 设置实例名称
func WithName(name string) Option {
	return func(o *options) { o.name = name }
}

// WithAddr 设置监听地址
func WithAddr(addr string) Option {
	return func(o *options) { o.addr = addr }
}

// WithCredentials 设置证书和秘钥
func WithCredentials(certFile, keyFile string) Option {
	return func(o *options) { o.keyFile, o.certFile = keyFile, certFile }
}

// WithRouteHandler 设置路由处理器
func WithRouteHandler(handler RouteHandler) Option {
	return func(o *options) { o.handler = handler }
}

// WithRegistry 设置服务注册器
func WithRegistry(r registry.Registry) Option {
	return func(o *options) { o.registry = r }
}

// WithTransporter 设置消息传输器
func WithTransporter(transporter transport.Transporter) Option {
	return func(o *options) { o.transporter = transporter }
}