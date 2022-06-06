package cache

import (
	"time"
)

type SomeStruct struct {
	value    string
	deadline time.Time
}
type Cache struct {
	arr map[string]SomeStruct
}

func NewCache() Cache {
	return Cache{arr: make(map[string]SomeStruct)}
}

func (receiver Cache) Get(key string) (string, bool) {

	for myKey, myValue := range receiver.arr {
		if myKey == key && (myValue.deadline.IsZero() || myValue.deadline.After(time.Now())) {
			return myValue.value, true
		}
	}

	return "", false
}

func (receiver *Cache) Put(key, value string) {

	receiver.arr[key] = SomeStruct{value: value, deadline: time.Time{}}
}

func (receiver *Cache) Keys() []string {

	var newArr []string
	for myKey, myValue := range receiver.arr {

		if myValue.deadline.IsZero() || myValue.deadline.After(time.Now()) {
			newArr = append(newArr, myKey)

		} else {
			delete(receiver.arr, myKey)
		}
	}

	return newArr
}

func (receiver *Cache) PutTill(key, value string, deadline time.Time) {

	receiver.arr[key] = SomeStruct{value: value, deadline: deadline}
}
