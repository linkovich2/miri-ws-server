package util

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
)

func RunEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

// this isn't currently being used anywhere? Maybe in a cleanup cycle we should utilize this
// but it might not be very performant seeing as we're using reflect a little
// (but it shouldn't be much worse then anything else)
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

func Sample(i []interface{}) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())

	if len(i) <= 0 {
		return nil, errors.New("Sampled slice cannot be empty")
	}

	return i[rand.Intn(len(i))], nil
}
