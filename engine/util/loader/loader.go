package loader

import (
  "io/ioutil"
  "encoding/json"
)

func Grab(filename string, marshalTo interface{}) {
  res, err := ioutil.ReadFile("./data/" + filename)
  if err != nil {
    panic(err)
  }

  json.Unmarshal(res, marshalTo)
}
