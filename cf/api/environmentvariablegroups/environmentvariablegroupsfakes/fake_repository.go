// Code generated by counterfeiter. DO NOT EDIT.
package environmentvariablegroupsfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api/environmentvariablegroups"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeRepository struct {
	ListRunningStub        func() ([]models.EnvironmentVariable, error)
	listRunningMutex       sync.RWMutex
	listRunningArgsForCall []struct {
	}
	listRunningReturns struct {
		result1 []models.EnvironmentVariable
		result2 error
	}
	listRunningReturnsOnCall map[int]struct {
		result1 []models.EnvironmentVariable
		result2 error
	}
	ListStagingStub        func() ([]models.EnvironmentVariable, error)
	listStagingMutex       sync.RWMutex
	listStagingArgsForCall []struct {
	}
	listStagingReturns struct {
		result1 []models.EnvironmentVariable
		result2 error
	}
	listStagingReturnsOnCall map[int]struct {
		result1 []models.EnvironmentVariable
		result2 error
	}
	SetRunningStub        func(string) error
	setRunningMutex       sync.RWMutex
	setRunningArgsForCall []struct {
		arg1 string
	}
	setRunningReturns struct {
		result1 error
	}
	setRunningReturnsOnCall map[int]struct {
		result1 error
	}
	SetStagingStub        func(string) error
	setStagingMutex       sync.RWMutex
	setStagingArgsForCall []struct {
		arg1 string
	}
	setStagingReturns struct {
		result1 error
	}
	setStagingReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) ListRunning() ([]models.EnvironmentVariable, error) {
	fake.listRunningMutex.Lock()
	ret, specificReturn := fake.listRunningReturnsOnCall[len(fake.listRunningArgsForCall)]
	fake.listRunningArgsForCall = append(fake.listRunningArgsForCall, struct {
	}{})
	fake.recordInvocation("ListRunning", []interface{}{})
	fake.listRunningMutex.Unlock()
	if fake.ListRunningStub != nil {
		return fake.ListRunningStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listRunningReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) ListRunningCallCount() int {
	fake.listRunningMutex.RLock()
	defer fake.listRunningMutex.RUnlock()
	return len(fake.listRunningArgsForCall)
}

func (fake *FakeRepository) ListRunningCalls(stub func() ([]models.EnvironmentVariable, error)) {
	fake.listRunningMutex.Lock()
	defer fake.listRunningMutex.Unlock()
	fake.ListRunningStub = stub
}

func (fake *FakeRepository) ListRunningReturns(result1 []models.EnvironmentVariable, result2 error) {
	fake.listRunningMutex.Lock()
	defer fake.listRunningMutex.Unlock()
	fake.ListRunningStub = nil
	fake.listRunningReturns = struct {
		result1 []models.EnvironmentVariable
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) ListRunningReturnsOnCall(i int, result1 []models.EnvironmentVariable, result2 error) {
	fake.listRunningMutex.Lock()
	defer fake.listRunningMutex.Unlock()
	fake.ListRunningStub = nil
	if fake.listRunningReturnsOnCall == nil {
		fake.listRunningReturnsOnCall = make(map[int]struct {
			result1 []models.EnvironmentVariable
			result2 error
		})
	}
	fake.listRunningReturnsOnCall[i] = struct {
		result1 []models.EnvironmentVariable
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) ListStaging() ([]models.EnvironmentVariable, error) {
	fake.listStagingMutex.Lock()
	ret, specificReturn := fake.listStagingReturnsOnCall[len(fake.listStagingArgsForCall)]
	fake.listStagingArgsForCall = append(fake.listStagingArgsForCall, struct {
	}{})
	fake.recordInvocation("ListStaging", []interface{}{})
	fake.listStagingMutex.Unlock()
	if fake.ListStagingStub != nil {
		return fake.ListStagingStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStagingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) ListStagingCallCount() int {
	fake.listStagingMutex.RLock()
	defer fake.listStagingMutex.RUnlock()
	return len(fake.listStagingArgsForCall)
}

func (fake *FakeRepository) ListStagingCalls(stub func() ([]models.EnvironmentVariable, error)) {
	fake.listStagingMutex.Lock()
	defer fake.listStagingMutex.Unlock()
	fake.ListStagingStub = stub
}

func (fake *FakeRepository) ListStagingReturns(result1 []models.EnvironmentVariable, result2 error) {
	fake.listStagingMutex.Lock()
	defer fake.listStagingMutex.Unlock()
	fake.ListStagingStub = nil
	fake.listStagingReturns = struct {
		result1 []models.EnvironmentVariable
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) ListStagingReturnsOnCall(i int, result1 []models.EnvironmentVariable, result2 error) {
	fake.listStagingMutex.Lock()
	defer fake.listStagingMutex.Unlock()
	fake.ListStagingStub = nil
	if fake.listStagingReturnsOnCall == nil {
		fake.listStagingReturnsOnCall = make(map[int]struct {
			result1 []models.EnvironmentVariable
			result2 error
		})
	}
	fake.listStagingReturnsOnCall[i] = struct {
		result1 []models.EnvironmentVariable
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) SetRunning(arg1 string) error {
	fake.setRunningMutex.Lock()
	ret, specificReturn := fake.setRunningReturnsOnCall[len(fake.setRunningArgsForCall)]
	fake.setRunningArgsForCall = append(fake.setRunningArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetRunning", []interface{}{arg1})
	fake.setRunningMutex.Unlock()
	if fake.SetRunningStub != nil {
		return fake.SetRunningStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setRunningReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) SetRunningCallCount() int {
	fake.setRunningMutex.RLock()
	defer fake.setRunningMutex.RUnlock()
	return len(fake.setRunningArgsForCall)
}

func (fake *FakeRepository) SetRunningCalls(stub func(string) error) {
	fake.setRunningMutex.Lock()
	defer fake.setRunningMutex.Unlock()
	fake.SetRunningStub = stub
}

func (fake *FakeRepository) SetRunningArgsForCall(i int) string {
	fake.setRunningMutex.RLock()
	defer fake.setRunningMutex.RUnlock()
	argsForCall := fake.setRunningArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) SetRunningReturns(result1 error) {
	fake.setRunningMutex.Lock()
	defer fake.setRunningMutex.Unlock()
	fake.SetRunningStub = nil
	fake.setRunningReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetRunningReturnsOnCall(i int, result1 error) {
	fake.setRunningMutex.Lock()
	defer fake.setRunningMutex.Unlock()
	fake.SetRunningStub = nil
	if fake.setRunningReturnsOnCall == nil {
		fake.setRunningReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setRunningReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetStaging(arg1 string) error {
	fake.setStagingMutex.Lock()
	ret, specificReturn := fake.setStagingReturnsOnCall[len(fake.setStagingArgsForCall)]
	fake.setStagingArgsForCall = append(fake.setStagingArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetStaging", []interface{}{arg1})
	fake.setStagingMutex.Unlock()
	if fake.SetStagingStub != nil {
		return fake.SetStagingStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setStagingReturns
	return fakeReturns.result1
}

func (fake *FakeRepository) SetStagingCallCount() int {
	fake.setStagingMutex.RLock()
	defer fake.setStagingMutex.RUnlock()
	return len(fake.setStagingArgsForCall)
}

func (fake *FakeRepository) SetStagingCalls(stub func(string) error) {
	fake.setStagingMutex.Lock()
	defer fake.setStagingMutex.Unlock()
	fake.SetStagingStub = stub
}

func (fake *FakeRepository) SetStagingArgsForCall(i int) string {
	fake.setStagingMutex.RLock()
	defer fake.setStagingMutex.RUnlock()
	argsForCall := fake.setStagingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) SetStagingReturns(result1 error) {
	fake.setStagingMutex.Lock()
	defer fake.setStagingMutex.Unlock()
	fake.SetStagingStub = nil
	fake.setStagingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) SetStagingReturnsOnCall(i int, result1 error) {
	fake.setStagingMutex.Lock()
	defer fake.setStagingMutex.Unlock()
	fake.SetStagingStub = nil
	if fake.setStagingReturnsOnCall == nil {
		fake.setStagingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setStagingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listRunningMutex.RLock()
	defer fake.listRunningMutex.RUnlock()
	fake.listStagingMutex.RLock()
	defer fake.listStagingMutex.RUnlock()
	fake.setRunningMutex.RLock()
	defer fake.setRunningMutex.RUnlock()
	fake.setStagingMutex.RLock()
	defer fake.setStagingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ environmentvariablegroups.Repository = new(FakeRepository)
