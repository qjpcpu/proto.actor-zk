package zk

import (
	"time"
)

const baseKey = `/protoactor`

type Option func(*config)

// WithAuth set zk auth
func WithAuth(scheme string, credential string) Option {
	return func(o *config) {
		o.Auth = authConfig{Scheme: scheme, Credential: credential}
	}
}

// WithBaseKey set actors base key
func WithBaseKey(key string) Option {
	return func(o *config) {
		if isStrBlank(key) {
			o.BaseKey = baseKey
		} else {
			o.BaseKey = formatBaseKey(key)
		}
	}
}

// WithSessionTimeout set zk session timeout
func WithSessionTimeout(tm time.Duration) Option {
	return func(o *config) {
		o.SessionTimeout = tm
	}
}

func withEndpoints(e []string) Option {
	return func(o *config) {
		o.Endpoints = e
	}
}

type authConfig struct {
	Scheme     string
	Credential string
}

func (za authConfig) isEmpty() bool {
	return za.Scheme == "" && za.Credential == ""
}

type config struct {
	BaseKey        string
	Endpoints      []string
	SessionTimeout time.Duration
	Auth           authConfig
}

func defaultConfig() *config {
	return &config{
		BaseKey:        baseKey,
		SessionTimeout: time.Second * 10,
	}
}
