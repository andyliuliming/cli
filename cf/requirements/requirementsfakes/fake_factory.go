// This file was generated by counterfeiter
package requirementsfakes

import (
	"sync"

	"github.com/blang/semver"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeFactory struct {
	NewApplicationRequirementStub        func(name string) requirements.ApplicationRequirement
	newApplicationRequirementMutex       sync.RWMutex
	newApplicationRequirementArgsForCall []struct {
		name string
	}
	newApplicationRequirementReturns struct {
		result1 requirements.ApplicationRequirement
	}
	NewDEAApplicationRequirementStub        func(name string) requirements.DEAApplicationRequirement
	newDEAApplicationRequirementMutex       sync.RWMutex
	newDEAApplicationRequirementArgsForCall []struct {
		name string
	}
	newDEAApplicationRequirementReturns struct {
		result1 requirements.DEAApplicationRequirement
	}
	NewDiegoApplicationRequirementStub        func(name string) requirements.DiegoApplicationRequirement
	newDiegoApplicationRequirementMutex       sync.RWMutex
	newDiegoApplicationRequirementArgsForCall []struct {
		name string
	}
	newDiegoApplicationRequirementReturns struct {
		result1 requirements.DiegoApplicationRequirement
	}
	NewServiceInstanceRequirementStub        func(name string) requirements.ServiceInstanceRequirement
	newServiceInstanceRequirementMutex       sync.RWMutex
	newServiceInstanceRequirementArgsForCall []struct {
		name string
	}
	newServiceInstanceRequirementReturns struct {
		result1 requirements.ServiceInstanceRequirement
	}
	NewLoginRequirementStub        func() requirements.Requirement
	newLoginRequirementMutex       sync.RWMutex
	newLoginRequirementArgsForCall []struct{}
	newLoginRequirementReturns     struct {
		result1 requirements.Requirement
	}
	NewRoutingAPIRequirementStub        func() requirements.Requirement
	newRoutingAPIRequirementMutex       sync.RWMutex
	newRoutingAPIRequirementArgsForCall []struct{}
	newRoutingAPIRequirementReturns     struct {
		result1 requirements.Requirement
	}
	NewSpaceRequirementStub        func(name string) requirements.SpaceRequirement
	newSpaceRequirementMutex       sync.RWMutex
	newSpaceRequirementArgsForCall []struct {
		name string
	}
	newSpaceRequirementReturns struct {
		result1 requirements.SpaceRequirement
	}
	NewTargetedSpaceRequirementStub        func() requirements.Requirement
	newTargetedSpaceRequirementMutex       sync.RWMutex
	newTargetedSpaceRequirementArgsForCall []struct{}
	newTargetedSpaceRequirementReturns     struct {
		result1 requirements.Requirement
	}
	NewTargetedOrgRequirementStub        func() requirements.TargetedOrgRequirement
	newTargetedOrgRequirementMutex       sync.RWMutex
	newTargetedOrgRequirementArgsForCall []struct{}
	newTargetedOrgRequirementReturns     struct {
		result1 requirements.TargetedOrgRequirement
	}
	NewOrganizationRequirementStub        func(name string) requirements.OrganizationRequirement
	newOrganizationRequirementMutex       sync.RWMutex
	newOrganizationRequirementArgsForCall []struct {
		name string
	}
	newOrganizationRequirementReturns struct {
		result1 requirements.OrganizationRequirement
	}
	NewDomainRequirementStub        func(name string) requirements.DomainRequirement
	newDomainRequirementMutex       sync.RWMutex
	newDomainRequirementArgsForCall []struct {
		name string
	}
	newDomainRequirementReturns struct {
		result1 requirements.DomainRequirement
	}
	NewUserRequirementStub        func(username string, wantGUID bool) requirements.UserRequirement
	newUserRequirementMutex       sync.RWMutex
	newUserRequirementArgsForCall []struct {
		username string
		wantGUID bool
	}
	newUserRequirementReturns struct {
		result1 requirements.UserRequirement
	}
	NewBuildpackRequirementStub        func(buildpack string) requirements.BuildpackRequirement
	newBuildpackRequirementMutex       sync.RWMutex
	newBuildpackRequirementArgsForCall []struct {
		buildpack string
	}
	newBuildpackRequirementReturns struct {
		result1 requirements.BuildpackRequirement
	}
	NewAPIEndpointRequirementStub        func() requirements.Requirement
	newAPIEndpointRequirementMutex       sync.RWMutex
	newAPIEndpointRequirementArgsForCall []struct{}
	newAPIEndpointRequirementReturns     struct {
		result1 requirements.Requirement
	}
	NewMinAPIVersionRequirementStub        func(commandName string, requiredVersion semver.Version) requirements.Requirement
	newMinAPIVersionRequirementMutex       sync.RWMutex
	newMinAPIVersionRequirementArgsForCall []struct {
		commandName     string
		requiredVersion semver.Version
	}
	newMinAPIVersionRequirementReturns struct {
		result1 requirements.Requirement
	}
	NewMaxAPIVersionRequirementStub        func(commandName string, maximumVersion semver.Version) requirements.Requirement
	newMaxAPIVersionRequirementMutex       sync.RWMutex
	newMaxAPIVersionRequirementArgsForCall []struct {
		commandName    string
		maximumVersion semver.Version
	}
	newMaxAPIVersionRequirementReturns struct {
		result1 requirements.Requirement
	}
	NewUsageRequirementStub        func(requirements.Usable, string, func() bool) requirements.Requirement
	newUsageRequirementMutex       sync.RWMutex
	newUsageRequirementArgsForCall []struct {
		arg1 requirements.Usable
		arg2 string
		arg3 func() bool
	}
	newUsageRequirementReturns struct {
		result1 requirements.Requirement
	}
	NewNumberArgumentsStub        func([]string, ...string) requirements.Requirement
	newNumberArgumentsMutex       sync.RWMutex
	newNumberArgumentsArgsForCall []struct {
		arg1 []string
		arg2 []string
	}
	newNumberArgumentsReturns struct {
		result1 requirements.Requirement
	}
}

func (fake *FakeFactory) NewApplicationRequirement(name string) requirements.ApplicationRequirement {
	fake.newApplicationRequirementMutex.Lock()
	fake.newApplicationRequirementArgsForCall = append(fake.newApplicationRequirementArgsForCall, struct {
		name string
	}{name})
	fake.newApplicationRequirementMutex.Unlock()
	if fake.NewApplicationRequirementStub != nil {
		return fake.NewApplicationRequirementStub(name)
	} else {
		return fake.newApplicationRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewApplicationRequirementCallCount() int {
	fake.newApplicationRequirementMutex.RLock()
	defer fake.newApplicationRequirementMutex.RUnlock()
	return len(fake.newApplicationRequirementArgsForCall)
}

func (fake *FakeFactory) NewApplicationRequirementArgsForCall(i int) string {
	fake.newApplicationRequirementMutex.RLock()
	defer fake.newApplicationRequirementMutex.RUnlock()
	return fake.newApplicationRequirementArgsForCall[i].name
}

func (fake *FakeFactory) NewApplicationRequirementReturns(result1 requirements.ApplicationRequirement) {
	fake.NewApplicationRequirementStub = nil
	fake.newApplicationRequirementReturns = struct {
		result1 requirements.ApplicationRequirement
	}{result1}
}

func (fake *FakeFactory) NewDEAApplicationRequirement(name string) requirements.DEAApplicationRequirement {
	fake.newDEAApplicationRequirementMutex.Lock()
	fake.newDEAApplicationRequirementArgsForCall = append(fake.newDEAApplicationRequirementArgsForCall, struct {
		name string
	}{name})
	fake.newDEAApplicationRequirementMutex.Unlock()
	if fake.NewDEAApplicationRequirementStub != nil {
		return fake.NewDEAApplicationRequirementStub(name)
	} else {
		return fake.newDEAApplicationRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewDEAApplicationRequirementCallCount() int {
	fake.newDEAApplicationRequirementMutex.RLock()
	defer fake.newDEAApplicationRequirementMutex.RUnlock()
	return len(fake.newDEAApplicationRequirementArgsForCall)
}

func (fake *FakeFactory) NewDEAApplicationRequirementArgsForCall(i int) string {
	fake.newDEAApplicationRequirementMutex.RLock()
	defer fake.newDEAApplicationRequirementMutex.RUnlock()
	return fake.newDEAApplicationRequirementArgsForCall[i].name
}

func (fake *FakeFactory) NewDEAApplicationRequirementReturns(result1 requirements.DEAApplicationRequirement) {
	fake.NewDEAApplicationRequirementStub = nil
	fake.newDEAApplicationRequirementReturns = struct {
		result1 requirements.DEAApplicationRequirement
	}{result1}
}

func (fake *FakeFactory) NewDiegoApplicationRequirement(name string) requirements.DiegoApplicationRequirement {
	fake.newDiegoApplicationRequirementMutex.Lock()
	fake.newDiegoApplicationRequirementArgsForCall = append(fake.newDiegoApplicationRequirementArgsForCall, struct {
		name string
	}{name})
	fake.newDiegoApplicationRequirementMutex.Unlock()
	if fake.NewDiegoApplicationRequirementStub != nil {
		return fake.NewDiegoApplicationRequirementStub(name)
	} else {
		return fake.newDiegoApplicationRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewDiegoApplicationRequirementCallCount() int {
	fake.newDiegoApplicationRequirementMutex.RLock()
	defer fake.newDiegoApplicationRequirementMutex.RUnlock()
	return len(fake.newDiegoApplicationRequirementArgsForCall)
}

func (fake *FakeFactory) NewDiegoApplicationRequirementArgsForCall(i int) string {
	fake.newDiegoApplicationRequirementMutex.RLock()
	defer fake.newDiegoApplicationRequirementMutex.RUnlock()
	return fake.newDiegoApplicationRequirementArgsForCall[i].name
}

func (fake *FakeFactory) NewDiegoApplicationRequirementReturns(result1 requirements.DiegoApplicationRequirement) {
	fake.NewDiegoApplicationRequirementStub = nil
	fake.newDiegoApplicationRequirementReturns = struct {
		result1 requirements.DiegoApplicationRequirement
	}{result1}
}

func (fake *FakeFactory) NewServiceInstanceRequirement(name string) requirements.ServiceInstanceRequirement {
	fake.newServiceInstanceRequirementMutex.Lock()
	fake.newServiceInstanceRequirementArgsForCall = append(fake.newServiceInstanceRequirementArgsForCall, struct {
		name string
	}{name})
	fake.newServiceInstanceRequirementMutex.Unlock()
	if fake.NewServiceInstanceRequirementStub != nil {
		return fake.NewServiceInstanceRequirementStub(name)
	} else {
		return fake.newServiceInstanceRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewServiceInstanceRequirementCallCount() int {
	fake.newServiceInstanceRequirementMutex.RLock()
	defer fake.newServiceInstanceRequirementMutex.RUnlock()
	return len(fake.newServiceInstanceRequirementArgsForCall)
}

func (fake *FakeFactory) NewServiceInstanceRequirementArgsForCall(i int) string {
	fake.newServiceInstanceRequirementMutex.RLock()
	defer fake.newServiceInstanceRequirementMutex.RUnlock()
	return fake.newServiceInstanceRequirementArgsForCall[i].name
}

func (fake *FakeFactory) NewServiceInstanceRequirementReturns(result1 requirements.ServiceInstanceRequirement) {
	fake.NewServiceInstanceRequirementStub = nil
	fake.newServiceInstanceRequirementReturns = struct {
		result1 requirements.ServiceInstanceRequirement
	}{result1}
}

func (fake *FakeFactory) NewLoginRequirement() requirements.Requirement {
	fake.newLoginRequirementMutex.Lock()
	fake.newLoginRequirementArgsForCall = append(fake.newLoginRequirementArgsForCall, struct{}{})
	fake.newLoginRequirementMutex.Unlock()
	if fake.NewLoginRequirementStub != nil {
		return fake.NewLoginRequirementStub()
	} else {
		return fake.newLoginRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewLoginRequirementCallCount() int {
	fake.newLoginRequirementMutex.RLock()
	defer fake.newLoginRequirementMutex.RUnlock()
	return len(fake.newLoginRequirementArgsForCall)
}

func (fake *FakeFactory) NewLoginRequirementReturns(result1 requirements.Requirement) {
	fake.NewLoginRequirementStub = nil
	fake.newLoginRequirementReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

func (fake *FakeFactory) NewRoutingAPIRequirement() requirements.Requirement {
	fake.newRoutingAPIRequirementMutex.Lock()
	fake.newRoutingAPIRequirementArgsForCall = append(fake.newRoutingAPIRequirementArgsForCall, struct{}{})
	fake.newRoutingAPIRequirementMutex.Unlock()
	if fake.NewRoutingAPIRequirementStub != nil {
		return fake.NewRoutingAPIRequirementStub()
	} else {
		return fake.newRoutingAPIRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewRoutingAPIRequirementCallCount() int {
	fake.newRoutingAPIRequirementMutex.RLock()
	defer fake.newRoutingAPIRequirementMutex.RUnlock()
	return len(fake.newRoutingAPIRequirementArgsForCall)
}

func (fake *FakeFactory) NewRoutingAPIRequirementReturns(result1 requirements.Requirement) {
	fake.NewRoutingAPIRequirementStub = nil
	fake.newRoutingAPIRequirementReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

func (fake *FakeFactory) NewSpaceRequirement(name string) requirements.SpaceRequirement {
	fake.newSpaceRequirementMutex.Lock()
	fake.newSpaceRequirementArgsForCall = append(fake.newSpaceRequirementArgsForCall, struct {
		name string
	}{name})
	fake.newSpaceRequirementMutex.Unlock()
	if fake.NewSpaceRequirementStub != nil {
		return fake.NewSpaceRequirementStub(name)
	} else {
		return fake.newSpaceRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewSpaceRequirementCallCount() int {
	fake.newSpaceRequirementMutex.RLock()
	defer fake.newSpaceRequirementMutex.RUnlock()
	return len(fake.newSpaceRequirementArgsForCall)
}

func (fake *FakeFactory) NewSpaceRequirementArgsForCall(i int) string {
	fake.newSpaceRequirementMutex.RLock()
	defer fake.newSpaceRequirementMutex.RUnlock()
	return fake.newSpaceRequirementArgsForCall[i].name
}

func (fake *FakeFactory) NewSpaceRequirementReturns(result1 requirements.SpaceRequirement) {
	fake.NewSpaceRequirementStub = nil
	fake.newSpaceRequirementReturns = struct {
		result1 requirements.SpaceRequirement
	}{result1}
}

func (fake *FakeFactory) NewTargetedSpaceRequirement() requirements.Requirement {
	fake.newTargetedSpaceRequirementMutex.Lock()
	fake.newTargetedSpaceRequirementArgsForCall = append(fake.newTargetedSpaceRequirementArgsForCall, struct{}{})
	fake.newTargetedSpaceRequirementMutex.Unlock()
	if fake.NewTargetedSpaceRequirementStub != nil {
		return fake.NewTargetedSpaceRequirementStub()
	} else {
		return fake.newTargetedSpaceRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewTargetedSpaceRequirementCallCount() int {
	fake.newTargetedSpaceRequirementMutex.RLock()
	defer fake.newTargetedSpaceRequirementMutex.RUnlock()
	return len(fake.newTargetedSpaceRequirementArgsForCall)
}

func (fake *FakeFactory) NewTargetedSpaceRequirementReturns(result1 requirements.Requirement) {
	fake.NewTargetedSpaceRequirementStub = nil
	fake.newTargetedSpaceRequirementReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

func (fake *FakeFactory) NewTargetedOrgRequirement() requirements.TargetedOrgRequirement {
	fake.newTargetedOrgRequirementMutex.Lock()
	fake.newTargetedOrgRequirementArgsForCall = append(fake.newTargetedOrgRequirementArgsForCall, struct{}{})
	fake.newTargetedOrgRequirementMutex.Unlock()
	if fake.NewTargetedOrgRequirementStub != nil {
		return fake.NewTargetedOrgRequirementStub()
	} else {
		return fake.newTargetedOrgRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewTargetedOrgRequirementCallCount() int {
	fake.newTargetedOrgRequirementMutex.RLock()
	defer fake.newTargetedOrgRequirementMutex.RUnlock()
	return len(fake.newTargetedOrgRequirementArgsForCall)
}

func (fake *FakeFactory) NewTargetedOrgRequirementReturns(result1 requirements.TargetedOrgRequirement) {
	fake.NewTargetedOrgRequirementStub = nil
	fake.newTargetedOrgRequirementReturns = struct {
		result1 requirements.TargetedOrgRequirement
	}{result1}
}

func (fake *FakeFactory) NewOrganizationRequirement(name string) requirements.OrganizationRequirement {
	fake.newOrganizationRequirementMutex.Lock()
	fake.newOrganizationRequirementArgsForCall = append(fake.newOrganizationRequirementArgsForCall, struct {
		name string
	}{name})
	fake.newOrganizationRequirementMutex.Unlock()
	if fake.NewOrganizationRequirementStub != nil {
		return fake.NewOrganizationRequirementStub(name)
	} else {
		return fake.newOrganizationRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewOrganizationRequirementCallCount() int {
	fake.newOrganizationRequirementMutex.RLock()
	defer fake.newOrganizationRequirementMutex.RUnlock()
	return len(fake.newOrganizationRequirementArgsForCall)
}

func (fake *FakeFactory) NewOrganizationRequirementArgsForCall(i int) string {
	fake.newOrganizationRequirementMutex.RLock()
	defer fake.newOrganizationRequirementMutex.RUnlock()
	return fake.newOrganizationRequirementArgsForCall[i].name
}

func (fake *FakeFactory) NewOrganizationRequirementReturns(result1 requirements.OrganizationRequirement) {
	fake.NewOrganizationRequirementStub = nil
	fake.newOrganizationRequirementReturns = struct {
		result1 requirements.OrganizationRequirement
	}{result1}
}

func (fake *FakeFactory) NewDomainRequirement(name string) requirements.DomainRequirement {
	fake.newDomainRequirementMutex.Lock()
	fake.newDomainRequirementArgsForCall = append(fake.newDomainRequirementArgsForCall, struct {
		name string
	}{name})
	fake.newDomainRequirementMutex.Unlock()
	if fake.NewDomainRequirementStub != nil {
		return fake.NewDomainRequirementStub(name)
	} else {
		return fake.newDomainRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewDomainRequirementCallCount() int {
	fake.newDomainRequirementMutex.RLock()
	defer fake.newDomainRequirementMutex.RUnlock()
	return len(fake.newDomainRequirementArgsForCall)
}

func (fake *FakeFactory) NewDomainRequirementArgsForCall(i int) string {
	fake.newDomainRequirementMutex.RLock()
	defer fake.newDomainRequirementMutex.RUnlock()
	return fake.newDomainRequirementArgsForCall[i].name
}

func (fake *FakeFactory) NewDomainRequirementReturns(result1 requirements.DomainRequirement) {
	fake.NewDomainRequirementStub = nil
	fake.newDomainRequirementReturns = struct {
		result1 requirements.DomainRequirement
	}{result1}
}

func (fake *FakeFactory) NewUserRequirement(username string, wantGUID bool) requirements.UserRequirement {
	fake.newUserRequirementMutex.Lock()
	fake.newUserRequirementArgsForCall = append(fake.newUserRequirementArgsForCall, struct {
		username string
		wantGUID bool
	}{username, wantGUID})
	fake.newUserRequirementMutex.Unlock()
	if fake.NewUserRequirementStub != nil {
		return fake.NewUserRequirementStub(username, wantGUID)
	} else {
		return fake.newUserRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewUserRequirementCallCount() int {
	fake.newUserRequirementMutex.RLock()
	defer fake.newUserRequirementMutex.RUnlock()
	return len(fake.newUserRequirementArgsForCall)
}

func (fake *FakeFactory) NewUserRequirementArgsForCall(i int) (string, bool) {
	fake.newUserRequirementMutex.RLock()
	defer fake.newUserRequirementMutex.RUnlock()
	return fake.newUserRequirementArgsForCall[i].username, fake.newUserRequirementArgsForCall[i].wantGUID
}

func (fake *FakeFactory) NewUserRequirementReturns(result1 requirements.UserRequirement) {
	fake.NewUserRequirementStub = nil
	fake.newUserRequirementReturns = struct {
		result1 requirements.UserRequirement
	}{result1}
}

func (fake *FakeFactory) NewBuildpackRequirement(buildpack string) requirements.BuildpackRequirement {
	fake.newBuildpackRequirementMutex.Lock()
	fake.newBuildpackRequirementArgsForCall = append(fake.newBuildpackRequirementArgsForCall, struct {
		buildpack string
	}{buildpack})
	fake.newBuildpackRequirementMutex.Unlock()
	if fake.NewBuildpackRequirementStub != nil {
		return fake.NewBuildpackRequirementStub(buildpack)
	} else {
		return fake.newBuildpackRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewBuildpackRequirementCallCount() int {
	fake.newBuildpackRequirementMutex.RLock()
	defer fake.newBuildpackRequirementMutex.RUnlock()
	return len(fake.newBuildpackRequirementArgsForCall)
}

func (fake *FakeFactory) NewBuildpackRequirementArgsForCall(i int) string {
	fake.newBuildpackRequirementMutex.RLock()
	defer fake.newBuildpackRequirementMutex.RUnlock()
	return fake.newBuildpackRequirementArgsForCall[i].buildpack
}

func (fake *FakeFactory) NewBuildpackRequirementReturns(result1 requirements.BuildpackRequirement) {
	fake.NewBuildpackRequirementStub = nil
	fake.newBuildpackRequirementReturns = struct {
		result1 requirements.BuildpackRequirement
	}{result1}
}

func (fake *FakeFactory) NewAPIEndpointRequirement() requirements.Requirement {
	fake.newAPIEndpointRequirementMutex.Lock()
	fake.newAPIEndpointRequirementArgsForCall = append(fake.newAPIEndpointRequirementArgsForCall, struct{}{})
	fake.newAPIEndpointRequirementMutex.Unlock()
	if fake.NewAPIEndpointRequirementStub != nil {
		return fake.NewAPIEndpointRequirementStub()
	} else {
		return fake.newAPIEndpointRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewAPIEndpointRequirementCallCount() int {
	fake.newAPIEndpointRequirementMutex.RLock()
	defer fake.newAPIEndpointRequirementMutex.RUnlock()
	return len(fake.newAPIEndpointRequirementArgsForCall)
}

func (fake *FakeFactory) NewAPIEndpointRequirementReturns(result1 requirements.Requirement) {
	fake.NewAPIEndpointRequirementStub = nil
	fake.newAPIEndpointRequirementReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

func (fake *FakeFactory) NewMinAPIVersionRequirement(commandName string, requiredVersion semver.Version) requirements.Requirement {
	fake.newMinAPIVersionRequirementMutex.Lock()
	fake.newMinAPIVersionRequirementArgsForCall = append(fake.newMinAPIVersionRequirementArgsForCall, struct {
		commandName     string
		requiredVersion semver.Version
	}{commandName, requiredVersion})
	fake.newMinAPIVersionRequirementMutex.Unlock()
	if fake.NewMinAPIVersionRequirementStub != nil {
		return fake.NewMinAPIVersionRequirementStub(commandName, requiredVersion)
	} else {
		return fake.newMinAPIVersionRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewMinAPIVersionRequirementCallCount() int {
	fake.newMinAPIVersionRequirementMutex.RLock()
	defer fake.newMinAPIVersionRequirementMutex.RUnlock()
	return len(fake.newMinAPIVersionRequirementArgsForCall)
}

func (fake *FakeFactory) NewMinAPIVersionRequirementArgsForCall(i int) (string, semver.Version) {
	fake.newMinAPIVersionRequirementMutex.RLock()
	defer fake.newMinAPIVersionRequirementMutex.RUnlock()
	return fake.newMinAPIVersionRequirementArgsForCall[i].commandName, fake.newMinAPIVersionRequirementArgsForCall[i].requiredVersion
}

func (fake *FakeFactory) NewMinAPIVersionRequirementReturns(result1 requirements.Requirement) {
	fake.NewMinAPIVersionRequirementStub = nil
	fake.newMinAPIVersionRequirementReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

func (fake *FakeFactory) NewMaxAPIVersionRequirement(commandName string, maximumVersion semver.Version) requirements.Requirement {
	fake.newMaxAPIVersionRequirementMutex.Lock()
	fake.newMaxAPIVersionRequirementArgsForCall = append(fake.newMaxAPIVersionRequirementArgsForCall, struct {
		commandName    string
		maximumVersion semver.Version
	}{commandName, maximumVersion})
	fake.newMaxAPIVersionRequirementMutex.Unlock()
	if fake.NewMaxAPIVersionRequirementStub != nil {
		return fake.NewMaxAPIVersionRequirementStub(commandName, maximumVersion)
	} else {
		return fake.newMaxAPIVersionRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewMaxAPIVersionRequirementCallCount() int {
	fake.newMaxAPIVersionRequirementMutex.RLock()
	defer fake.newMaxAPIVersionRequirementMutex.RUnlock()
	return len(fake.newMaxAPIVersionRequirementArgsForCall)
}

func (fake *FakeFactory) NewMaxAPIVersionRequirementArgsForCall(i int) (string, semver.Version) {
	fake.newMaxAPIVersionRequirementMutex.RLock()
	defer fake.newMaxAPIVersionRequirementMutex.RUnlock()
	return fake.newMaxAPIVersionRequirementArgsForCall[i].commandName, fake.newMaxAPIVersionRequirementArgsForCall[i].maximumVersion
}

func (fake *FakeFactory) NewMaxAPIVersionRequirementReturns(result1 requirements.Requirement) {
	fake.NewMaxAPIVersionRequirementStub = nil
	fake.newMaxAPIVersionRequirementReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

func (fake *FakeFactory) NewUsageRequirement(arg1 requirements.Usable, arg2 string, arg3 func() bool) requirements.Requirement {
	fake.newUsageRequirementMutex.Lock()
	fake.newUsageRequirementArgsForCall = append(fake.newUsageRequirementArgsForCall, struct {
		arg1 requirements.Usable
		arg2 string
		arg3 func() bool
	}{arg1, arg2, arg3})
	fake.newUsageRequirementMutex.Unlock()
	if fake.NewUsageRequirementStub != nil {
		return fake.NewUsageRequirementStub(arg1, arg2, arg3)
	} else {
		return fake.newUsageRequirementReturns.result1
	}
}

func (fake *FakeFactory) NewUsageRequirementCallCount() int {
	fake.newUsageRequirementMutex.RLock()
	defer fake.newUsageRequirementMutex.RUnlock()
	return len(fake.newUsageRequirementArgsForCall)
}

func (fake *FakeFactory) NewUsageRequirementArgsForCall(i int) (requirements.Usable, string, func() bool) {
	fake.newUsageRequirementMutex.RLock()
	defer fake.newUsageRequirementMutex.RUnlock()
	return fake.newUsageRequirementArgsForCall[i].arg1, fake.newUsageRequirementArgsForCall[i].arg2, fake.newUsageRequirementArgsForCall[i].arg3
}

func (fake *FakeFactory) NewUsageRequirementReturns(result1 requirements.Requirement) {
	fake.NewUsageRequirementStub = nil
	fake.newUsageRequirementReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

func (fake *FakeFactory) NewNumberArguments(arg1 []string, arg2 ...string) requirements.Requirement {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.newNumberArgumentsMutex.Lock()
	fake.newNumberArgumentsArgsForCall = append(fake.newNumberArgumentsArgsForCall, struct {
		arg1 []string
		arg2 []string
	}{arg1Copy, arg2})
	fake.newNumberArgumentsMutex.Unlock()
	if fake.NewNumberArgumentsStub != nil {
		return fake.NewNumberArgumentsStub(arg1, arg2...)
	} else {
		return fake.newNumberArgumentsReturns.result1
	}
}

func (fake *FakeFactory) NewNumberArgumentsCallCount() int {
	fake.newNumberArgumentsMutex.RLock()
	defer fake.newNumberArgumentsMutex.RUnlock()
	return len(fake.newNumberArgumentsArgsForCall)
}

func (fake *FakeFactory) NewNumberArgumentsArgsForCall(i int) ([]string, []string) {
	fake.newNumberArgumentsMutex.RLock()
	defer fake.newNumberArgumentsMutex.RUnlock()
	return fake.newNumberArgumentsArgsForCall[i].arg1, fake.newNumberArgumentsArgsForCall[i].arg2
}

func (fake *FakeFactory) NewNumberArgumentsReturns(result1 requirements.Requirement) {
	fake.NewNumberArgumentsStub = nil
	fake.newNumberArgumentsReturns = struct {
		result1 requirements.Requirement
	}{result1}
}

var _ requirements.Factory = new(FakeFactory)
