package state

import (
	"github.com/hnit-acm/hfunc/basich"
)

// todo 完善状态机逻辑
// labels: enhancement
// assignees: nieaowei
type EventMap map[StateKey]func(key StateKey) StateFunc

// 状态key
type StateKey interface{}

// 状态实例
type StateMachineMap basich.MapFunc

type StateFunc func()

type SetStateFunc func(key StateKey, stateFunc StateFunc)
type DelStateFunc func(key StateKey)
type ExecStateFunc func(key StateKey, sync bool) bool

type StateMachineFunc func() (SetStateFunc, ExecStateFunc, DelStateFunc)

var StateMachine = StateMachineFunc(func() (SetStateFunc, ExecStateFunc, DelStateFunc) {
	// 状态维护
	get, set, del := basich.NewSyncMapFunc(1024)
	SetStateFunc := SetStateFunc(func(key StateKey, stateFunc StateFunc) {
		set(key, stateFunc)
	})

	ExecStateFunc := ExecStateFunc(func(key StateKey, sync bool) bool {
		if sync {
			f, ok := get(key)
			if ok {
				f.(StateFunc)()
			} else {
				return ok
			}
		} else {
			f, ok := get(key)
			if ok {
				f.(StateFunc)()
			} else {
				return ok
			}
		}
		return true
	})

	DelStateFunc := DelStateFunc(func(key StateKey) {
		del(key)
	})
	return SetStateFunc, ExecStateFunc, DelStateFunc
})
