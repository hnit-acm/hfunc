package basic

import (
	"hash/crc32"
	"sync"
)

// 函数式map

type SetFunc func(key, val interface{})
type GetFunc func(key interface{}) (interface{}, bool)
type DelFunc func(key interface{})

type HashFormatFunc func(key interface{}) uint32

var defaultHashFormatFunc = HashFormatFunc(func(key interface{}) uint32 {
	return crc32.ChecksumIEEE([]byte(key.(string)))
})

type HashMap []map[interface{}]interface{}

type MapFunc func(bufSize uint32, hashFormats ...HashFormatFunc) (GetFunc, SetFunc, DelFunc)

var NewHashMapFunc = MapFunc(
	func(bufSize uint32, hashFormats ...HashFormatFunc) (GetFunc, SetFunc, DelFunc) {
		var (
			data       HashMap
			mus        []sync.RWMutex
			hashFormat HashFormatFunc
		)
		data = make(HashMap, bufSize)
		mus = make([]sync.RWMutex, bufSize)
		hashFormat = defaultHashFormatFunc
		for _, format := range hashFormats {
			hashFormat = format
			break
		}
		get := GetFunc(func(key interface{}) (interface{}, bool) {
			hashVal := hashFormat(key) % bufSize
			mus[hashVal].RLock()
			data, ok := data[hashVal][key]
			mus[hashVal].RUnlock()
			return data, ok
		})
		set := SetFunc(func(key, val interface{}) {
			hashVal := hashFormat(key) % bufSize
			mus[hashVal].Lock()
			if data[hashVal] == nil {
				data[hashVal] = make(map[interface{}]interface{})
			}
			data[hashVal][key] = val
			mus[hashVal].Unlock()
		})
		del := DelFunc(func(key interface{}) {
			hashVal := hashFormat(key) % bufSize
			mus[hashVal].Lock()
			delete(data[hashVal], key)
			mus[hashVal].Unlock()
		})
		return get, set, del
	},
)

// 面向对象式map

type Map interface {
	Set() SetFunc
	Get() GetFunc
	Del() DelFunc
}

type MapStruct struct {
	set SetFunc
	get GetFunc
	del DelFunc
}

func NewHashMapStruct(bufSize uint32, hashFormats ...HashFormatFunc) Map {
	get, set, del := NewHashMapFunc(bufSize, hashFormats...)
	return MapStruct{
		set: set,
		get: get,
		del: del,
	}
}

func (m MapStruct) Set() SetFunc {
	return m.set
}

func (m MapStruct) Get() GetFunc {
	return m.get
}

func (m MapStruct) Del() DelFunc {
	return m.del
}

var NewSyncMapFunc = MapFunc(
	func(bufSize uint32, hashFormats ...HashFormatFunc) (GetFunc, SetFunc, DelFunc) {
		var (
			data sync.Map
		)
		data = sync.Map{}

		get := GetFunc(func(key interface{}) (interface{}, bool) {
			return data.Load(key)
		})
		set := SetFunc(func(key, val interface{}) {
			data.Store(key, val)
		})
		del := DelFunc(func(key interface{}) {
			data.Delete(key)
		})
		return get, set, del
	},
)

func NewSyncMapStruct(bufSize uint32, hashFormats ...HashFormatFunc) Map {
	get, set, del := NewSyncMapFunc(bufSize, hashFormats...)
	return MapStruct{
		set: set,
		get: get,
		del: del,
	}
}
