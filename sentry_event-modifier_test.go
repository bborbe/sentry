// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sentry_test

import (
	"github.com/getsentry/sentry-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	libsentry "github.com/bborbe/sentry"
)

var _ = Describe("EventModifier", func() {
	Describe("EventModifierFunc", func() {
		It("implements EventModifier interface", func() {
			modifier := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					return event
				},
			)
			var _ libsentry.EventModifier = modifier
		})

		It("applies function to event", func() {
			event := &sentry.Event{Message: "original"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			modifier := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					event.Message = "modified"
					return event
				},
			)

			result := modifier.ApplyToEvent(event, hint, client)

			Expect(result).To(Equal(event))
			Expect(result.Message).To(Equal("modified"))
		})

		It("can return different event", func() {
			originalEvent := &sentry.Event{Message: "original"}
			newEvent := &sentry.Event{Message: "new"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			modifier := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					return newEvent
				},
			)

			result := modifier.ApplyToEvent(originalEvent, hint, client)

			Expect(result).To(Equal(newEvent))
			Expect(result.Message).To(Equal("new"))
		})

		It("receives all parameters", func() {
			event := &sentry.Event{Message: "test"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			var receivedEvent *sentry.Event
			var receivedHint *sentry.EventHint
			var receivedClient *sentry.Client

			modifier := libsentry.EventModifierFunc(
				func(e *sentry.Event, h *sentry.EventHint, c *sentry.Client) *sentry.Event {
					receivedEvent = e
					receivedHint = h
					receivedClient = c
					return e
				},
			)

			modifier.ApplyToEvent(event, hint, client)

			Expect(receivedEvent).To(Equal(event))
			Expect(receivedHint).To(Equal(hint))
			Expect(receivedClient).To(Equal(client))
		})
	})

	Describe("EventModifierList", func() {
		It("implements EventModifier interface", func() {
			var modifierList libsentry.EventModifierList
			var _ libsentry.EventModifier = modifierList
		})

		It("applies no modifiers when empty", func() {
			event := &sentry.Event{Message: "original"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			var modifierList libsentry.EventModifierList

			result := modifierList.ApplyToEvent(event, hint, client)

			Expect(result).To(Equal(event))
			Expect(result.Message).To(Equal("original"))
		})

		It("applies single modifier", func() {
			event := &sentry.Event{Message: "original"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			modifier := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					event.Message = "modified"
					return event
				},
			)

			modifierList := libsentry.EventModifierList{modifier}

			result := modifierList.ApplyToEvent(event, hint, client)

			Expect(result).To(Equal(event))
			Expect(result.Message).To(Equal("modified"))
		})

		It("applies multiple modifiers in sequence", func() {
			event := &sentry.Event{Message: "original"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			modifier1 := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					event.Message = event.Message + "-1"
					return event
				},
			)

			modifier2 := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					event.Message = event.Message + "-2"
					return event
				},
			)

			modifier3 := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					event.Message = event.Message + "-3"
					return event
				},
			)

			modifierList := libsentry.EventModifierList{modifier1, modifier2, modifier3}

			result := modifierList.ApplyToEvent(event, hint, client)

			Expect(result).To(Equal(event))
			Expect(result.Message).To(Equal("original-1-2-3"))
		})

		It("passes modified event to subsequent modifiers", func() {
			originalEvent := &sentry.Event{Message: "original"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			var eventsReceived []*sentry.Event

			modifier1 := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					eventsReceived = append(eventsReceived, event)
					newEvent := &sentry.Event{Message: "modified"}
					return newEvent
				},
			)

			modifier2 := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					eventsReceived = append(eventsReceived, event)
					return event
				},
			)

			modifierList := libsentry.EventModifierList{modifier1, modifier2}

			result := modifierList.ApplyToEvent(originalEvent, hint, client)

			Expect(eventsReceived).To(HaveLen(2))
			Expect(eventsReceived[0]).To(Equal(originalEvent))
			Expect(eventsReceived[0].Message).To(Equal("original"))
			Expect(eventsReceived[1].Message).To(Equal("modified"))
			Expect(result.Message).To(Equal("modified"))
		})

		It("handles nil event from modifier", func() {
			event := &sentry.Event{Message: "original"}
			hint := &sentry.EventHint{}
			var client *sentry.Client

			modifier1 := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					return nil
				},
			)

			modifier2 := libsentry.EventModifierFunc(
				func(event *sentry.Event, hint *sentry.EventHint, client *sentry.Client) *sentry.Event {
					if event == nil {
						return &sentry.Event{Message: "recreated"}
					}
					return event
				},
			)

			modifierList := libsentry.EventModifierList{modifier1, modifier2}

			result := modifierList.ApplyToEvent(event, hint, client)

			Expect(result).ToNot(BeNil())
			Expect(result.Message).To(Equal("recreated"))
		})
	})
})
