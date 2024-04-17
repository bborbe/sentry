// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sentry_test

import (
	"context"
	"errors"

	runmocks "github.com/bborbe/run/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/sentry"
	sentrymocks "github.com/bborbe/sentry/mocks"
)

var _ = Describe("SkipErrorAndReport", func() {
	var ctx context.Context
	var err error
	var sentryClient *sentrymocks.SentryClient
	var runnable *runmocks.Runnable
	BeforeEach(func() {
		ctx = context.Background()
		runnable = &runmocks.Runnable{}
		sentryClient = &sentrymocks.SentryClient{}
	})
	JustBeforeEach(func() {
		skipErrorAndReport := sentry.NewSkipErrorAndReport(sentryClient, runnable)
		err = skipErrorAndReport.Run(ctx)
	})
	Context("success", func() {
		BeforeEach(func() {
			runnable.RunReturns(nil)
		})
		It("calls no captureException", func() {
			Expect(sentryClient.CaptureExceptionCallCount()).To(Equal(0))
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("failure", func() {
		BeforeEach(func() {
			runnable.RunReturns(errors.New("banana"))
		})
		It("calls captureException", func() {
			Expect(sentryClient.CaptureExceptionCallCount()).To(Equal(1))
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
	})
})
