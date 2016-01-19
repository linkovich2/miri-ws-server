package util

import (
	"errors"
	"math/rand"
	"reflect"
	"strings"
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

// We're not using this very often, maybe we need to make it more flexible
func Sample(i []string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	if len(i) <= 0 {
		return "", errors.New("Sampled slice cannot be empty")
	}

	return i[rand.Intn(len(i))], nil
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}
