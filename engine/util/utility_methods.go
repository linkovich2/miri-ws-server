package util

import (
	"reflect"
	"time"
)

func RunEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

func InArray(value interface{}, array interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(array))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = value == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}
