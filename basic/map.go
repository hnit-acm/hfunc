package basic

import (
	"hash/crc32"
	"sync"
)

type SafeMap interface {
	Set(key, val interface{})
	Get(key interface{}) (interface{}, bool)
	GetByCondition(f func(key, val interface{}) bool) (interface{}, bool)
	ForEach(f func(key, val interface{}))
	Delete(key interface{})
}

type HashFormatFunc func(key interface{}) uint32

var defaultHashFormatFunc = HashFormatFunc(func(key interface{}) uint32 {
	return crc32.ChecksumIEEE([]byte(key.(string)))
})

type ConcurrentHashMap struct {
	data []map[interface{}]interface{}
	mus  []sync.RWMutex
	// Partition size
	size uint32
	// Hash format function.
	// Example: from int to []byte].
	// func(key interface{}) []byte {
	//		return []byte(strings.Itoa(key))
	// }
	hashFormat HashFormatFunc
}

func NewConcurrentHashMap(bufSize uint32, hashFormats ...HashFormatFunc) *ConcurrentHashMap {
	// Allocate storage space for partition.
	d := make([]map[interface{}]interface{}, bufSize)
	// Allocate storage space for lock.
	m := make([]sync.RWMutex, bufSize)
	// Generate instance for partition and lock.
	for i := uint32(0); i < bufSize; i++ {
		d[i] = make(map[interface{}]interface{})
		m[i] = sync.RWMutex{}
	}
	hashFormat := defaultHashFormatFunc

	for _, format := range hashFormats {
		hashFormat = format
		continue
	}

	return &ConcurrentHashMap{
		size:       bufSize,
		data:       d,
		mus:        m,
		hashFormat: hashFormat,
	}
}

func (m *ConcurrentHashMap) Set(key, val interface{}) {
	hashVal := m.hashFormat(key) % m.size
	m.mus[hashVal].Lock()
	m.data[hashVal][key] = val
	m.mus[hashVal].Unlock()
}

func (m *ConcurrentHashMap) Delete(key interface{}) {
	hashVal := m.hashFormat(key) % m.size
	m.mus[hashVal].Lock()
	delete(m.data[hashVal], key)
	m.mus[hashVal].Unlock()
}

func (m *ConcurrentHashMap) Get(key interface{}) (interface{}, bool) {
	hashVal := m.hashFormat(key) % m.size
	m.mus[hashVal].RLock()
	data, ok := m.data[hashVal][key]
	m.mus[hashVal].RUnlock()
	return data, ok
}

func (m *ConcurrentHashMap) GetByCondition(f func(key, value interface{}) bool) (interface{}, bool) {
	for i := uint32(0); i < m.size; i++ {
		for k, v := range m.data[i] {
			if f(k, v) == true {
				return v, true
			}
		}
	}
	return nil, false
}

func (m *ConcurrentHashMap) ForEach(f func(key, value interface{})) {
	for i := uint32(0); i < m.size; i++ {
		for k, v := range m.data[i] {
			f(k, v)
		}
	}
}
