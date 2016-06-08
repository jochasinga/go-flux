package stores

import (
	"crypto/md5"
	"encoding/json"
)

type (
	BasicStore map[string][]interface{}
	HashStore  map[string][]interface{}
)

// Basic Store
func (b *BasicStore) Add(key string, data interface{}) {
	b[key] = append(b[key], data)
}

func (b *BasicStore) Del(key string) {
	delete(b, key)
}

func (b *BasicStore) Get(key string) interface{} {
	return b[key]
}

func (b *BasicStore) Set(key string, data interface{}) {
	b[key] = []interface{}{data}
}

// Hash Store
func (hs *HashStore) Add(data interface{}) (string, error) {
	bytedat, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	h := md5.New()
	key := string(h.Sum(bytedat)[:])
	hs[key] = data
	return key, nil
}

func (b *BasicStore) Del(key string) {
	delete(b, key)
}

func (b *BasicStore) Get(key string) interface{} {
	return b[key]
}

func (b *BasicStore) Set(key string, data interface{}) {
	b[key] = []interface{}{data}
}
