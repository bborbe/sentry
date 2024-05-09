// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	stderrors "errors"
	"flag"
	"runtime"
	"time"

	"github.com/bborbe/errors"
	"github.com/getsentry/sentry-go"
	"github.com/golang/glog"

	libsentry "github.com/bborbe/sentry"
)

var dsnPtr = flag.String("dsn", "", "sentry dsn")

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	runtime.GOMAXPROCS(runtime.NumCPU())
	_ = flag.Set("logtostderr", "true")
	_ = flag.Set("v", "2")

	time.Local = time.UTC
	glog.V(2).Infof("set global timezone to UTC")

	ctx := context.Background()

	flag.Parse()

	client, err := libsentry.NewClient(ctx, sentry.ClientOptions{
		Dsn:  *dsnPtr,
		Tags: map[string]string{"options-data": "options-value"},
	})
	if err != nil {
		glog.Exitf("create client failed: %v", err)
	}
	defer func() {
		_ = client.Flush(2 * time.Second)
		_ = client.Close()
	}()

	ctx = errors.AddToContext(ctx, "context-data", "context-value")
	err = errors.AddDataToError(
		stderrors.New("banana"),
		map[string]string{"error-data": "error-value"},
	)
	client.CaptureException(err, &sentry.EventHint{
		OriginalException: err,
		Context:           ctx,
	}, sentry.NewScope())

	glog.V(2).Infof("done")
}
