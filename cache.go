package cache

import (
	"time"
)

type SomeStruct struct {
	value    string
	deadline time.Time
	dead     bool
}
type Cache struct {
	//key      string
	//value    string
	arr map[string]SomeStruct
}

func NewCache() Cache {
	return Cache{arr: make(map[string]SomeStruct)}
}

func (receiver Cache) Get(key string) (string, bool) {

	for key1, value1 := range receiver.arr {
		if key1 == key && !value1.deadline.IsZero() {
			return value1.value, value1.dead
		} else {
			return value1.value, false
		}

	}
	return "", false
}

func (receiver Cache) Put(key, value string) {

	for key1, value1 := range receiver.arr {
		if key1 == key {
			value1.value = value
			value1.deadline = time.Time{}
			value1.dead = false

		} else {
			receiver.arr[key] = SomeStruct{value: value, deadline: time.Time{}, dead: false}

		}

	}
}

func (receiver Cache) Keys() []string {

	var newArr []string
	for key1, value1 := range receiver.arr {
		if !value1.deadline.IsZero() {
			newArr = append(newArr, key1)
		}
	}

	return newArr
}

func (receiver Cache) PutTill(key, value string, deadline time.Time) {

	for key1, value1 := range receiver.arr {
		if key1 == key {
			value1.value = value
			value1.deadline = deadline
			value1.dead = false
		} else {
			receiver.arr[key] = SomeStruct{value: value, deadline: deadline, dead: false}
		}

	}
}
