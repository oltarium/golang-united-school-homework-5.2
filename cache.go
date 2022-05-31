package cache

import "time"

type Value struct {
	value     string
	exp       time.Time
	canExpire bool
}

type Cache struct {
	cache map[string]Value
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]Value),
	}
}

func (cache Cache) Get(key string) (string, bool) {
	value, ok := cache.cache[key]
	if !value.canExpire || (value.canExpire && value.exp.After(time.Now())) {
		return value.value, ok
	}
	return "", false
}

func (cache Cache) Put(key, value string) {
	cache.cache[key] = Value{value: value, canExpire: false}
}

func (cache Cache) Keys() []string {
	keys := []string{}
	for key, value := range cache.cache {
		if !value.canExpire || (value.canExpire && value.exp.After(time.Now())) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (cache Cache) PutTill(key, value string, deadline time.Time) {
	cache.cache[key] = Value{value: value, exp: deadline, canExpire: true}
}

func (cache Cache) Clear() {
	cache.cache = make(map[string]Value)
}
