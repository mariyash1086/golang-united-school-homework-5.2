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

	for key1, value1 := range receiver.arr {
		if key1 == key && (value1.deadline.After(time.Now()) || value1.deadline.IsZero()) {
			return value1.value, true
		} else {
			return value1.value, false
		}
	}
	return key, false
}

func (receiver Cache) Put(key, value string) {

	receiver.arr[key] = SomeStruct{value: value, deadline: time.Time{}}
}

func (receiver Cache) Keys() []string {

	var newArr []string
	for key1, value1 := range receiver.arr {

		if value1.deadline.After(time.Now()) || value1.deadline.IsZero() {
			newArr = append(newArr, key1)

		} else {
			delete(receiver.arr, key1)
		}
	}

	return newArr
}

func (receiver Cache) PutTill(key, value string, deadline time.Time) {

	receiver.arr[key] = SomeStruct{value: value, deadline: deadline}
}
