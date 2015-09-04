// This file was generated by counterfeiter
package fake_secure_dialer

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/secure_shell"
	"golang.org/x/crypto/ssh"
)

type FakeSecureDialer struct {
	DialStub        func(network, addr string, config *ssh.ClientConfig) (secure_shell.SecureClient, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
		network string
		addr    string
		config  *ssh.ClientConfig
	}
	dialReturns struct {
		result1 secure_shell.SecureClient
		result2 error
	}
}

func (fake *FakeSecureDialer) Dial(network string, addr string, config *ssh.ClientConfig) (secure_shell.SecureClient, error) {
	fake.dialMutex.Lock()
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
		network string
		addr    string
		config  *ssh.ClientConfig
	}{network, addr, config})
	fake.dialMutex.Unlock()
	if fake.DialStub != nil {
		return fake.DialStub(network, addr, config)
	} else {
		return fake.dialReturns.result1, fake.dialReturns.result2
	}
}

func (fake *FakeSecureDialer) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *FakeSecureDialer) DialArgsForCall(i int) (string, string, *ssh.ClientConfig) {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return fake.dialArgsForCall[i].network, fake.dialArgsForCall[i].addr, fake.dialArgsForCall[i].config
}

func (fake *FakeSecureDialer) DialReturns(result1 secure_shell.SecureClient, result2 error) {
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 secure_shell.SecureClient
		result2 error
	}{result1, result2}
}

var _ secure_shell.SecureDialer = new(FakeSecureDialer)