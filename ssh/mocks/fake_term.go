// This file was generated by counterfeiter
package mocks

import (
	"sync"

	"github.com/cloudfoundry-incubator/ltc/ssh"
	"github.com/docker/docker/pkg/term"
)

type FakeTerm struct {
	SetRawTerminalStub        func(fd uintptr) (*term.State, error)
	setRawTerminalMutex       sync.RWMutex
	setRawTerminalArgsForCall []struct {
		fd uintptr
	}
	setRawTerminalReturns struct {
		result1 *term.State
		result2 error
	}
	RestoreTerminalStub        func(fd uintptr, state *term.State) error
	restoreTerminalMutex       sync.RWMutex
	restoreTerminalArgsForCall []struct {
		fd    uintptr
		state *term.State
	}
	restoreTerminalReturns struct {
		result1 error
	}
	GetWinsizeStub        func(fd uintptr) (width int, height int)
	getWinsizeMutex       sync.RWMutex
	getWinsizeArgsForCall []struct {
		fd uintptr
	}
	getWinsizeReturns struct {
		result1 int
		result2 int
	}
	IsTTYStub        func(fd uintptr) bool
	isTTYMutex       sync.RWMutex
	isTTYArgsForCall []struct {
		fd uintptr
	}
	isTTYReturns struct {
		result1 bool
	}
}

func (fake *FakeTerm) SetRawTerminal(fd uintptr) (*term.State, error) {
	fake.setRawTerminalMutex.Lock()
	fake.setRawTerminalArgsForCall = append(fake.setRawTerminalArgsForCall, struct {
		fd uintptr
	}{fd})
	fake.setRawTerminalMutex.Unlock()
	if fake.SetRawTerminalStub != nil {
		return fake.SetRawTerminalStub(fd)
	} else {
		return fake.setRawTerminalReturns.result1, fake.setRawTerminalReturns.result2
	}
}

func (fake *FakeTerm) SetRawTerminalCallCount() int {
	fake.setRawTerminalMutex.RLock()
	defer fake.setRawTerminalMutex.RUnlock()
	return len(fake.setRawTerminalArgsForCall)
}

func (fake *FakeTerm) SetRawTerminalArgsForCall(i int) uintptr {
	fake.setRawTerminalMutex.RLock()
	defer fake.setRawTerminalMutex.RUnlock()
	return fake.setRawTerminalArgsForCall[i].fd
}

func (fake *FakeTerm) SetRawTerminalReturns(result1 *term.State, result2 error) {
	fake.SetRawTerminalStub = nil
	fake.setRawTerminalReturns = struct {
		result1 *term.State
		result2 error
	}{result1, result2}
}

func (fake *FakeTerm) RestoreTerminal(fd uintptr, state *term.State) error {
	fake.restoreTerminalMutex.Lock()
	fake.restoreTerminalArgsForCall = append(fake.restoreTerminalArgsForCall, struct {
		fd    uintptr
		state *term.State
	}{fd, state})
	fake.restoreTerminalMutex.Unlock()
	if fake.RestoreTerminalStub != nil {
		return fake.RestoreTerminalStub(fd, state)
	} else {
		return fake.restoreTerminalReturns.result1
	}
}

func (fake *FakeTerm) RestoreTerminalCallCount() int {
	fake.restoreTerminalMutex.RLock()
	defer fake.restoreTerminalMutex.RUnlock()
	return len(fake.restoreTerminalArgsForCall)
}

func (fake *FakeTerm) RestoreTerminalArgsForCall(i int) (uintptr, *term.State) {
	fake.restoreTerminalMutex.RLock()
	defer fake.restoreTerminalMutex.RUnlock()
	return fake.restoreTerminalArgsForCall[i].fd, fake.restoreTerminalArgsForCall[i].state
}

func (fake *FakeTerm) RestoreTerminalReturns(result1 error) {
	fake.RestoreTerminalStub = nil
	fake.restoreTerminalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTerm) GetWinsize(fd uintptr) (width int, height int) {
	fake.getWinsizeMutex.Lock()
	fake.getWinsizeArgsForCall = append(fake.getWinsizeArgsForCall, struct {
		fd uintptr
	}{fd})
	fake.getWinsizeMutex.Unlock()
	if fake.GetWinsizeStub != nil {
		return fake.GetWinsizeStub(fd)
	} else {
		return fake.getWinsizeReturns.result1, fake.getWinsizeReturns.result2
	}
}

func (fake *FakeTerm) GetWinsizeCallCount() int {
	fake.getWinsizeMutex.RLock()
	defer fake.getWinsizeMutex.RUnlock()
	return len(fake.getWinsizeArgsForCall)
}

func (fake *FakeTerm) GetWinsizeArgsForCall(i int) uintptr {
	fake.getWinsizeMutex.RLock()
	defer fake.getWinsizeMutex.RUnlock()
	return fake.getWinsizeArgsForCall[i].fd
}

func (fake *FakeTerm) GetWinsizeReturns(result1 int, result2 int) {
	fake.GetWinsizeStub = nil
	fake.getWinsizeReturns = struct {
		result1 int
		result2 int
	}{result1, result2}
}

func (fake *FakeTerm) IsTTY(fd uintptr) bool {
	fake.isTTYMutex.Lock()
	fake.isTTYArgsForCall = append(fake.isTTYArgsForCall, struct {
		fd uintptr
	}{fd})
	fake.isTTYMutex.Unlock()
	if fake.IsTTYStub != nil {
		return fake.IsTTYStub(fd)
	} else {
		return fake.isTTYReturns.result1
	}
}

func (fake *FakeTerm) IsTTYCallCount() int {
	fake.isTTYMutex.RLock()
	defer fake.isTTYMutex.RUnlock()
	return len(fake.isTTYArgsForCall)
}

func (fake *FakeTerm) IsTTYArgsForCall(i int) uintptr {
	fake.isTTYMutex.RLock()
	defer fake.isTTYMutex.RUnlock()
	return fake.isTTYArgsForCall[i].fd
}

func (fake *FakeTerm) IsTTYReturns(result1 bool) {
	fake.IsTTYStub = nil
	fake.isTTYReturns = struct {
		result1 bool
	}{result1}
}

var _ ssh.Term = new(FakeTerm)