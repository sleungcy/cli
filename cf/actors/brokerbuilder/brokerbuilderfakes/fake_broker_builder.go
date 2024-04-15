// Code generated by counterfeiter. DO NOT EDIT.
package brokerbuilderfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/actors/brokerbuilder"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeBrokerBuilder struct {
	AttachBrokersToServicesStub        func([]models.ServiceOffering) ([]models.ServiceBroker, error)
	attachBrokersToServicesMutex       sync.RWMutex
	attachBrokersToServicesArgsForCall []struct {
		arg1 []models.ServiceOffering
	}
	attachBrokersToServicesReturns struct {
		result1 []models.ServiceBroker
		result2 error
	}
	attachBrokersToServicesReturnsOnCall map[int]struct {
		result1 []models.ServiceBroker
		result2 error
	}
	AttachSpecificBrokerToServicesStub        func(string, []models.ServiceOffering) (models.ServiceBroker, error)
	attachSpecificBrokerToServicesMutex       sync.RWMutex
	attachSpecificBrokerToServicesArgsForCall []struct {
		arg1 string
		arg2 []models.ServiceOffering
	}
	attachSpecificBrokerToServicesReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	attachSpecificBrokerToServicesReturnsOnCall map[int]struct {
		result1 models.ServiceBroker
		result2 error
	}
	GetAllServiceBrokersStub        func() ([]models.ServiceBroker, error)
	getAllServiceBrokersMutex       sync.RWMutex
	getAllServiceBrokersArgsForCall []struct {
	}
	getAllServiceBrokersReturns struct {
		result1 []models.ServiceBroker
		result2 error
	}
	getAllServiceBrokersReturnsOnCall map[int]struct {
		result1 []models.ServiceBroker
		result2 error
	}
	GetBrokerWithAllServicesStub        func(string) (models.ServiceBroker, error)
	getBrokerWithAllServicesMutex       sync.RWMutex
	getBrokerWithAllServicesArgsForCall []struct {
		arg1 string
	}
	getBrokerWithAllServicesReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	getBrokerWithAllServicesReturnsOnCall map[int]struct {
		result1 models.ServiceBroker
		result2 error
	}
	GetBrokerWithSpecifiedServiceStub        func(string) (models.ServiceBroker, error)
	getBrokerWithSpecifiedServiceMutex       sync.RWMutex
	getBrokerWithSpecifiedServiceArgsForCall []struct {
		arg1 string
	}
	getBrokerWithSpecifiedServiceReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	getBrokerWithSpecifiedServiceReturnsOnCall map[int]struct {
		result1 models.ServiceBroker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBrokerBuilder) AttachBrokersToServices(arg1 []models.ServiceOffering) ([]models.ServiceBroker, error) {
	var arg1Copy []models.ServiceOffering
	if arg1 != nil {
		arg1Copy = make([]models.ServiceOffering, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.attachBrokersToServicesMutex.Lock()
	ret, specificReturn := fake.attachBrokersToServicesReturnsOnCall[len(fake.attachBrokersToServicesArgsForCall)]
	fake.attachBrokersToServicesArgsForCall = append(fake.attachBrokersToServicesArgsForCall, struct {
		arg1 []models.ServiceOffering
	}{arg1Copy})
	fake.recordInvocation("AttachBrokersToServices", []interface{}{arg1Copy})
	fake.attachBrokersToServicesMutex.Unlock()
	if fake.AttachBrokersToServicesStub != nil {
		return fake.AttachBrokersToServicesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.attachBrokersToServicesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesCallCount() int {
	fake.attachBrokersToServicesMutex.RLock()
	defer fake.attachBrokersToServicesMutex.RUnlock()
	return len(fake.attachBrokersToServicesArgsForCall)
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesCalls(stub func([]models.ServiceOffering) ([]models.ServiceBroker, error)) {
	fake.attachBrokersToServicesMutex.Lock()
	defer fake.attachBrokersToServicesMutex.Unlock()
	fake.AttachBrokersToServicesStub = stub
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesArgsForCall(i int) []models.ServiceOffering {
	fake.attachBrokersToServicesMutex.RLock()
	defer fake.attachBrokersToServicesMutex.RUnlock()
	argsForCall := fake.attachBrokersToServicesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesReturns(result1 []models.ServiceBroker, result2 error) {
	fake.attachBrokersToServicesMutex.Lock()
	defer fake.attachBrokersToServicesMutex.Unlock()
	fake.AttachBrokersToServicesStub = nil
	fake.attachBrokersToServicesReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) AttachBrokersToServicesReturnsOnCall(i int, result1 []models.ServiceBroker, result2 error) {
	fake.attachBrokersToServicesMutex.Lock()
	defer fake.attachBrokersToServicesMutex.Unlock()
	fake.AttachBrokersToServicesStub = nil
	if fake.attachBrokersToServicesReturnsOnCall == nil {
		fake.attachBrokersToServicesReturnsOnCall = make(map[int]struct {
			result1 []models.ServiceBroker
			result2 error
		})
	}
	fake.attachBrokersToServicesReturnsOnCall[i] = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServices(arg1 string, arg2 []models.ServiceOffering) (models.ServiceBroker, error) {
	var arg2Copy []models.ServiceOffering
	if arg2 != nil {
		arg2Copy = make([]models.ServiceOffering, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.attachSpecificBrokerToServicesMutex.Lock()
	ret, specificReturn := fake.attachSpecificBrokerToServicesReturnsOnCall[len(fake.attachSpecificBrokerToServicesArgsForCall)]
	fake.attachSpecificBrokerToServicesArgsForCall = append(fake.attachSpecificBrokerToServicesArgsForCall, struct {
		arg1 string
		arg2 []models.ServiceOffering
	}{arg1, arg2Copy})
	fake.recordInvocation("AttachSpecificBrokerToServices", []interface{}{arg1, arg2Copy})
	fake.attachSpecificBrokerToServicesMutex.Unlock()
	if fake.AttachSpecificBrokerToServicesStub != nil {
		return fake.AttachSpecificBrokerToServicesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.attachSpecificBrokerToServicesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesCallCount() int {
	fake.attachSpecificBrokerToServicesMutex.RLock()
	defer fake.attachSpecificBrokerToServicesMutex.RUnlock()
	return len(fake.attachSpecificBrokerToServicesArgsForCall)
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesCalls(stub func(string, []models.ServiceOffering) (models.ServiceBroker, error)) {
	fake.attachSpecificBrokerToServicesMutex.Lock()
	defer fake.attachSpecificBrokerToServicesMutex.Unlock()
	fake.AttachSpecificBrokerToServicesStub = stub
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesArgsForCall(i int) (string, []models.ServiceOffering) {
	fake.attachSpecificBrokerToServicesMutex.RLock()
	defer fake.attachSpecificBrokerToServicesMutex.RUnlock()
	argsForCall := fake.attachSpecificBrokerToServicesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesReturns(result1 models.ServiceBroker, result2 error) {
	fake.attachSpecificBrokerToServicesMutex.Lock()
	defer fake.attachSpecificBrokerToServicesMutex.Unlock()
	fake.AttachSpecificBrokerToServicesStub = nil
	fake.attachSpecificBrokerToServicesReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) AttachSpecificBrokerToServicesReturnsOnCall(i int, result1 models.ServiceBroker, result2 error) {
	fake.attachSpecificBrokerToServicesMutex.Lock()
	defer fake.attachSpecificBrokerToServicesMutex.Unlock()
	fake.AttachSpecificBrokerToServicesStub = nil
	if fake.attachSpecificBrokerToServicesReturnsOnCall == nil {
		fake.attachSpecificBrokerToServicesReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBroker
			result2 error
		})
	}
	fake.attachSpecificBrokerToServicesReturnsOnCall[i] = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokers() ([]models.ServiceBroker, error) {
	fake.getAllServiceBrokersMutex.Lock()
	ret, specificReturn := fake.getAllServiceBrokersReturnsOnCall[len(fake.getAllServiceBrokersArgsForCall)]
	fake.getAllServiceBrokersArgsForCall = append(fake.getAllServiceBrokersArgsForCall, struct {
	}{})
	fake.recordInvocation("GetAllServiceBrokers", []interface{}{})
	fake.getAllServiceBrokersMutex.Unlock()
	if fake.GetAllServiceBrokersStub != nil {
		return fake.GetAllServiceBrokersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAllServiceBrokersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokersCallCount() int {
	fake.getAllServiceBrokersMutex.RLock()
	defer fake.getAllServiceBrokersMutex.RUnlock()
	return len(fake.getAllServiceBrokersArgsForCall)
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokersCalls(stub func() ([]models.ServiceBroker, error)) {
	fake.getAllServiceBrokersMutex.Lock()
	defer fake.getAllServiceBrokersMutex.Unlock()
	fake.GetAllServiceBrokersStub = stub
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokersReturns(result1 []models.ServiceBroker, result2 error) {
	fake.getAllServiceBrokersMutex.Lock()
	defer fake.getAllServiceBrokersMutex.Unlock()
	fake.GetAllServiceBrokersStub = nil
	fake.getAllServiceBrokersReturns = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetAllServiceBrokersReturnsOnCall(i int, result1 []models.ServiceBroker, result2 error) {
	fake.getAllServiceBrokersMutex.Lock()
	defer fake.getAllServiceBrokersMutex.Unlock()
	fake.GetAllServiceBrokersStub = nil
	if fake.getAllServiceBrokersReturnsOnCall == nil {
		fake.getAllServiceBrokersReturnsOnCall = make(map[int]struct {
			result1 []models.ServiceBroker
			result2 error
		})
	}
	fake.getAllServiceBrokersReturnsOnCall[i] = struct {
		result1 []models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServices(arg1 string) (models.ServiceBroker, error) {
	fake.getBrokerWithAllServicesMutex.Lock()
	ret, specificReturn := fake.getBrokerWithAllServicesReturnsOnCall[len(fake.getBrokerWithAllServicesArgsForCall)]
	fake.getBrokerWithAllServicesArgsForCall = append(fake.getBrokerWithAllServicesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetBrokerWithAllServices", []interface{}{arg1})
	fake.getBrokerWithAllServicesMutex.Unlock()
	if fake.GetBrokerWithAllServicesStub != nil {
		return fake.GetBrokerWithAllServicesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBrokerWithAllServicesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesCallCount() int {
	fake.getBrokerWithAllServicesMutex.RLock()
	defer fake.getBrokerWithAllServicesMutex.RUnlock()
	return len(fake.getBrokerWithAllServicesArgsForCall)
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesCalls(stub func(string) (models.ServiceBroker, error)) {
	fake.getBrokerWithAllServicesMutex.Lock()
	defer fake.getBrokerWithAllServicesMutex.Unlock()
	fake.GetBrokerWithAllServicesStub = stub
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesArgsForCall(i int) string {
	fake.getBrokerWithAllServicesMutex.RLock()
	defer fake.getBrokerWithAllServicesMutex.RUnlock()
	argsForCall := fake.getBrokerWithAllServicesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesReturns(result1 models.ServiceBroker, result2 error) {
	fake.getBrokerWithAllServicesMutex.Lock()
	defer fake.getBrokerWithAllServicesMutex.Unlock()
	fake.GetBrokerWithAllServicesStub = nil
	fake.getBrokerWithAllServicesReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetBrokerWithAllServicesReturnsOnCall(i int, result1 models.ServiceBroker, result2 error) {
	fake.getBrokerWithAllServicesMutex.Lock()
	defer fake.getBrokerWithAllServicesMutex.Unlock()
	fake.GetBrokerWithAllServicesStub = nil
	if fake.getBrokerWithAllServicesReturnsOnCall == nil {
		fake.getBrokerWithAllServicesReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBroker
			result2 error
		})
	}
	fake.getBrokerWithAllServicesReturnsOnCall[i] = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedService(arg1 string) (models.ServiceBroker, error) {
	fake.getBrokerWithSpecifiedServiceMutex.Lock()
	ret, specificReturn := fake.getBrokerWithSpecifiedServiceReturnsOnCall[len(fake.getBrokerWithSpecifiedServiceArgsForCall)]
	fake.getBrokerWithSpecifiedServiceArgsForCall = append(fake.getBrokerWithSpecifiedServiceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetBrokerWithSpecifiedService", []interface{}{arg1})
	fake.getBrokerWithSpecifiedServiceMutex.Unlock()
	if fake.GetBrokerWithSpecifiedServiceStub != nil {
		return fake.GetBrokerWithSpecifiedServiceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBrokerWithSpecifiedServiceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceCallCount() int {
	fake.getBrokerWithSpecifiedServiceMutex.RLock()
	defer fake.getBrokerWithSpecifiedServiceMutex.RUnlock()
	return len(fake.getBrokerWithSpecifiedServiceArgsForCall)
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceCalls(stub func(string) (models.ServiceBroker, error)) {
	fake.getBrokerWithSpecifiedServiceMutex.Lock()
	defer fake.getBrokerWithSpecifiedServiceMutex.Unlock()
	fake.GetBrokerWithSpecifiedServiceStub = stub
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceArgsForCall(i int) string {
	fake.getBrokerWithSpecifiedServiceMutex.RLock()
	defer fake.getBrokerWithSpecifiedServiceMutex.RUnlock()
	argsForCall := fake.getBrokerWithSpecifiedServiceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceReturns(result1 models.ServiceBroker, result2 error) {
	fake.getBrokerWithSpecifiedServiceMutex.Lock()
	defer fake.getBrokerWithSpecifiedServiceMutex.Unlock()
	fake.GetBrokerWithSpecifiedServiceStub = nil
	fake.getBrokerWithSpecifiedServiceReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) GetBrokerWithSpecifiedServiceReturnsOnCall(i int, result1 models.ServiceBroker, result2 error) {
	fake.getBrokerWithSpecifiedServiceMutex.Lock()
	defer fake.getBrokerWithSpecifiedServiceMutex.Unlock()
	fake.GetBrokerWithSpecifiedServiceStub = nil
	if fake.getBrokerWithSpecifiedServiceReturnsOnCall == nil {
		fake.getBrokerWithSpecifiedServiceReturnsOnCall = make(map[int]struct {
			result1 models.ServiceBroker
			result2 error
		})
	}
	fake.getBrokerWithSpecifiedServiceReturnsOnCall[i] = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerBuilder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.attachBrokersToServicesMutex.RLock()
	defer fake.attachBrokersToServicesMutex.RUnlock()
	fake.attachSpecificBrokerToServicesMutex.RLock()
	defer fake.attachSpecificBrokerToServicesMutex.RUnlock()
	fake.getAllServiceBrokersMutex.RLock()
	defer fake.getAllServiceBrokersMutex.RUnlock()
	fake.getBrokerWithAllServicesMutex.RLock()
	defer fake.getBrokerWithAllServicesMutex.RUnlock()
	fake.getBrokerWithSpecifiedServiceMutex.RLock()
	defer fake.getBrokerWithSpecifiedServiceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBrokerBuilder) recordInvocation(key string, args []interface{}) {
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

var _ brokerbuilder.BrokerBuilder = new(FakeBrokerBuilder)
