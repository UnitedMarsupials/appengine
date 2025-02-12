// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package internal

import (
	"context"
	"errors"
	"os"

	"google.golang.org/protobuf/proto"
)

type ctxKey string

func (c ctxKey) String() string {
	return "appengine aeContext key: " + string(c)
}

var errNotAppEngineContext = errors.New("not an App Engine aeContext")

type CallOverrideFunc func(ctx context.Context, service, method string, in, out proto.Message) error

var callOverrideKey = "holds []CallOverrideFunc"

func WithCallOverride(ctx context.Context, f CallOverrideFunc) context.Context {
	// We avoid appending to any existing call override
	// so we don't risk overwriting a popped stack below.
	var cofs []CallOverrideFunc
	if uf, ok := ctx.Value(&callOverrideKey).([]CallOverrideFunc); ok {
		cofs = append(cofs, uf...)
	}
	cofs = append(cofs, f)
	return context.WithValue(ctx, &callOverrideKey, cofs)
}

func callOverrideFromContext(ctx context.Context) (CallOverrideFunc, context.Context, bool) {
	cofs, _ := ctx.Value(&callOverrideKey).([]CallOverrideFunc)
	if len(cofs) == 0 {
		return nil, nil, false
	}
	// We found a list of overrides; grab the last, and reconstitute a
	// aeContext that will hide it.
	f := cofs[len(cofs)-1]
	ctx = context.WithValue(ctx, &callOverrideKey, cofs[:len(cofs)-1])
	return f, ctx, true
}

type logOverrideFunc func(level int64, format string, args ...interface{})

var logOverrideKey = "holds a logOverrideFunc"

func WithLogOverride(ctx context.Context, f logOverrideFunc) context.Context {
	return context.WithValue(ctx, &logOverrideKey, f)
}

var appIDOverrideKey = "holds a string, being the full app ID"

func WithAppIDOverride(ctx context.Context, appID string) context.Context {
	return context.WithValue(ctx, &appIDOverrideKey, appID)
}

var apiHostOverrideKey = ctxKey("holds a string, being the alternate API_HOST")

func withAPIHostOverride(ctx context.Context, apiHost string) context.Context {
	return context.WithValue(ctx, apiHostOverrideKey, apiHost)
}

var apiPortOverrideKey = ctxKey("holds a string, being the alternate API_PORT")

func withAPIPortOverride(ctx context.Context, apiPort string) context.Context {
	return context.WithValue(ctx, apiPortOverrideKey, apiPort)
}

var namespaceKey = "holds the namespace string"

func withNamespace(ctx context.Context, ns string) context.Context {
	return context.WithValue(ctx, &namespaceKey, ns)
}

func NamespaceFromContext(ctx context.Context) string {
	// If there's no namespace, return the empty string.
	ns, _ := ctx.Value(&namespaceKey).(string)
	return ns
}

// FullyQualifiedAppID returns the fully-qualified application ID.
// This may contain a partition prefix (e.g. "s~" for High Replication apps),
// or a domain prefix (e.g. "example.com:").
func FullyQualifiedAppID(ctx context.Context) string {
	if id, ok := ctx.Value(&appIDOverrideKey).(string); ok {
		return id
	}
	return fullyQualifiedAppID(ctx)
}

func Logf(ctx context.Context, level int64, format string, args ...interface{}) {
	if f, ok := ctx.Value(&logOverrideKey).(logOverrideFunc); ok {
		f(level, format, args...)
		return
	}
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	logf(c, level, format, args...)
}

// NamespacedContext wraps a Context to support namespaces.
func NamespacedContext(ctx context.Context, namespace string) context.Context {
	return withNamespace(ctx, namespace)
}

// SetTestEnv sets the env variables for testing background ticket in Flex.
func SetTestEnv() func() {
	var environ = []struct {
		key, value string
	}{
		{"GAE_LONG_APP_ID", "my-app-id"},
		{"GAE_MINOR_VERSION", "067924799508853122"},
		{"GAE_MODULE_INSTANCE", "0"},
		{"GAE_MODULE_NAME", "default"},
		{"GAE_MODULE_VERSION", "20150612t184001"},
	}

	for _, v := range environ {
		old := os.Getenv(v.key)
		os.Setenv(v.key, v.value)
		v.value = old
	}
	return func() { // Restore old environment after the test completes.
		for _, v := range environ {
			if v.value == "" {
				os.Unsetenv(v.key)
				continue
			}
			os.Setenv(v.key, v.value)
		}
	}
}
