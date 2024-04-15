// Code generated by counterfeiter. DO NOT EDIT.
package requirementsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/cf/requirements"
)

type FakeConfigRefresher struct {
	RefreshStub        func() (coreconfig.Warning, error)
	refreshMutex       sync.RWMutex
	refreshArgsForCall []struct {
	}
	refreshReturns struct {
		result1 coreconfig.Warning
		result2 error
	}
	refreshReturnsOnCall map[int]struct {
		result1 coreconfig.Warning
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfigRefresher) Refresh() (coreconfig.Warning, error) {
	fake.refreshMutex.Lock()
	ret, specificReturn := fake.refreshReturnsOnCall[len(fake.refreshArgsForCall)]
	fake.refreshArgsForCall = append(fake.refreshArgsForCall, struct {
	}{})
	fake.recordInvocation("Refresh", []interface{}{})
	fake.refreshMutex.Unlock()
	if fake.RefreshStub != nil {
		return fake.RefreshStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.refreshReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConfigRefresher) RefreshCallCount() int {
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	return len(fake.refreshArgsForCall)
}

func (fake *FakeConfigRefresher) RefreshCalls(stub func() (coreconfig.Warning, error)) {
	fake.refreshMutex.Lock()
	defer fake.refreshMutex.Unlock()
	fake.RefreshStub = stub
}

func (fake *FakeConfigRefresher) RefreshReturns(result1 coreconfig.Warning, result2 error) {
	fake.refreshMutex.Lock()
	defer fake.refreshMutex.Unlock()
	fake.RefreshStub = nil
	fake.refreshReturns = struct {
		result1 coreconfig.Warning
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigRefresher) RefreshReturnsOnCall(i int, result1 coreconfig.Warning, result2 error) {
	fake.refreshMutex.Lock()
	defer fake.refreshMutex.Unlock()
	fake.RefreshStub = nil
	if fake.refreshReturnsOnCall == nil {
		fake.refreshReturnsOnCall = make(map[int]struct {
			result1 coreconfig.Warning
			result2 error
		})
	}
	fake.refreshReturnsOnCall[i] = struct {
		result1 coreconfig.Warning
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigRefresher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.refreshMutex.RLock()
	defer fake.refreshMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfigRefresher) recordInvocation(key string, args []interface{}) {
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

var _ requirements.ConfigRefresher = new(FakeConfigRefresher)
