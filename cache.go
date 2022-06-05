package cache

import (
	"time"
)

type Cache struct {
	//key      string
	//value    string
	arr []struct {
		key      string
		value    string
		deadline time.Time
	}
}

func NewCache() Cache {
	return Cache{}
}

func (receiver Cache) Get(key string) (string, bool) {

	for _, elem := range receiver.arr {
		if elem.key == key && (time.Now().Before(elem.deadline) || time.Now().Equal(elem.deadline)) {
			return elem.value, true
		} else {
			return elem.value, false
		}

	}
	return "", false
}

func (receiver Cache) Put(key, value string) {

	for _, elem := range receiver.arr {
		if elem.key == key {
			elem.value = value
			elem.deadline = time.Date(3999, 12, 31, 1, 1, 1, 1, time.Local)
		}

	}
}

func (receiver Cache) Keys() []string {

	var newArr []string
	for _, elem := range receiver.arr {
		if time.Now().Before(elem.deadline) || time.Now().Equal(elem.deadline) {
			newArr = append(newArr, elem.key)
		}

	}

	return newArr
}

func (receiver Cache) PutTill(key, value string, deadline time.Time) {

	for _, elem := range receiver.arr {
		if elem.key == key {
			elem.value = value
			elem.deadline = deadline
		}

	}
}
