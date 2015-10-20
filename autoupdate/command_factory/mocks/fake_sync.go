// This file was generated by counterfeiter
package mocks

import (
	"sync"

	"github.com/cloudfoundry-incubator/ltc/autoupdate/command_factory"
	config_package "github.com/cloudfoundry-incubator/ltc/config"
)

type FakeSync struct {
	SyncLTCStub        func(ltcPath string, arch string, config *config_package.Config) error
	syncLTCMutex       sync.RWMutex
	syncLTCArgsForCall []struct {
		ltcPath string
		arch    string
		config  *config_package.Config
	}
	syncLTCReturns struct {
		result1 error
	}
}

func (fake *FakeSync) SyncLTC(ltcPath string, arch string, config *config_package.Config) error {
	fake.syncLTCMutex.Lock()
	fake.syncLTCArgsForCall = append(fake.syncLTCArgsForCall, struct {
		ltcPath string
		arch    string
		config  *config_package.Config
	}{ltcPath, arch, config})
	fake.syncLTCMutex.Unlock()
	if fake.SyncLTCStub != nil {
		return fake.SyncLTCStub(ltcPath, arch, config)
	} else {
		return fake.syncLTCReturns.result1
	}
}

func (fake *FakeSync) SyncLTCCallCount() int {
	fake.syncLTCMutex.RLock()
	defer fake.syncLTCMutex.RUnlock()
	return len(fake.syncLTCArgsForCall)
}

func (fake *FakeSync) SyncLTCArgsForCall(i int) (string, string, *config_package.Config) {
	fake.syncLTCMutex.RLock()
	defer fake.syncLTCMutex.RUnlock()
	return fake.syncLTCArgsForCall[i].ltcPath, fake.syncLTCArgsForCall[i].arch, fake.syncLTCArgsForCall[i].config
}

func (fake *FakeSync) SyncLTCReturns(result1 error) {
	fake.SyncLTCStub = nil
	fake.syncLTCReturns = struct {
		result1 error
	}{result1}
}

var _ command_factory.Sync = new(FakeSync)
