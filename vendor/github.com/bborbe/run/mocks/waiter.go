// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"
	"time"

	"github.com/bborbe/run"
)

type Waiter struct {
	WaitStub        func(context.Context, time.Duration) error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
		arg1 context.Context
		arg2 time.Duration
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Waiter) Wait(arg1 context.Context, arg2 time.Duration) error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
		arg1 context.Context
		arg2 time.Duration
	}{arg1, arg2})
	stub := fake.WaitStub
	fakeReturns := fake.waitReturns
	fake.recordInvocation("Wait", []interface{}{arg1, arg2})
	fake.waitMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Waiter) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *Waiter) WaitCalls(stub func(context.Context, time.Duration) error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *Waiter) WaitArgsForCall(i int) (context.Context, time.Duration) {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	argsForCall := fake.waitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Waiter) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *Waiter) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Waiter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Waiter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ run.Waiter = new(Waiter)