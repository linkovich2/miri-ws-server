package loader

import (
	"encoding/json"
	"io/ioutil"
)

func Grab(filename string, marshalTo interface{}) {
	res, err := ioutil.ReadFile("./data/" + filename)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(res, marshalTo)
}
