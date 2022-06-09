package client

import (
	"time"
)

var (
	DefaultRedirectLimit       = 5
	DefaultClientTimeout       = time.Second * 10
	DefaultMaxIdleConns        = 32
	DefaultIdleConnTimeout     = time.Second * 60
	DefaultMaxConnsPerHost     = 128
	DefaultMaxIdleConnsPerHost = 32
)
