package cache

import (
	"time"
)

type Cache struct {
	//key      string
	//value    string
	arr      map[string]string
	deadline time.Time
}

func NewCache() Cache {
	return Cache{}
}

func (receiver Cache) Get(key string) (string, bool) {

	for key1, value1 := range receiver.arr {
		if key1 == key && (time.Now().Before(receiver.deadline) || time.Now().Equal(receiver.deadline)) {
			return value1, true
		} else {
			return value1, false
		}

	}
	return "", false
}

func (receiver Cache) Put(key, value string) {

	for key1, _ := range receiver.arr {
		if key1 == key {
			receiver.arr[key1] = value
			receiver.deadline = time.Date(3999, 12, 31, 1, 1, 1, 1, time.Local)
		}

	}
}

func (receiver Cache) Keys() []string {

	var newArr []string
	for key1, _ := range receiver.arr {
		if time.Now().Before(receiver.deadline) || time.Now().Equal(receiver.deadline) {
			newArr = append(newArr, key1)
		}

	}

	return newArr
}

func (receiver Cache) PutTill(key, value string, deadline time.Time) {

	for key1, _ := range receiver.arr {
		if key1 == key {
			receiver.arr[key1] = value
			receiver.deadline = deadline
		}

	}
}
