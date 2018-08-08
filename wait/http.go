package wait

import (
	"time"
	"context"
)

// Implement interface
var _ WaitStrategy = (*httpWaitStrategy)(nil)

type httpWaitStrategy struct {
	// all WaitStrategies should have a startupTimeout to avoid waiting infinitely
	startupTimeout time.Duration

	// httpWaitStrategy has additional properties
	path              string
	statusCodeMatcher func(status int) bool
	useTls            bool
}

// Constructor
func HttpWaitStrategyNew(path string) *httpWaitStrategy {
	return &httpWaitStrategy{
		startupTimeout:    defaultStartupTimeout(),
		statusCodeMatcher: nil,
		path:              path,
		useTls:            false,
	}

}

// fluent builders for each property
// since go has neither covariance nor generics, the return type must be the type of the concrete implementation
// this is true for all properties, even the "shared" ones like startupTimeout
func (http *httpWaitStrategy) WithStartupTimeout(startupTimeout time.Duration) *httpWaitStrategy {
	http.startupTimeout = startupTimeout
	return http
}

func (http *httpWaitStrategy) WithStatusCodeMatcher(statusCodeMatcher func(status int) bool) *httpWaitStrategy {
	http.statusCodeMatcher = statusCodeMatcher
	return http
}

func (http *httpWaitStrategy) WithTls(useTls bool) *httpWaitStrategy {
	http.useTls = useTls
	return http
}

// Convenience method similar to Wait.java
// https://github.com/testcontainers/testcontainers-java/blob/1d85a3834bd937f80aad3a4cec249c027f31aeb4/core/src/main/java/org/testcontainers/containers/wait/strategy/Wait.java
func ForHttp(path string) *httpWaitStrategy {
	return HttpWaitStrategyNew(path)
}

// Implementation of WaitStrategy.WaitUntilReady
func (http *httpWaitStrategy) WaitUntilReady(ctx context.Context, waitStrategyTarget WaitStrategyTarget) error {
	panic("implement me")
}
