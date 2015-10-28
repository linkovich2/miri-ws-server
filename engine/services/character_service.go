package services

import (
  "net/http"

  "github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
)

func CreateCharacter(character *parameters.Character, userId string) (status int, body []byte) {
  // @todo stub
  return http.StatusCreated, []byte("Whatever")
}

func ListCharacters(userId string) (status int, body []byte) {
  // @todo stub
  return http.StatusOK, []byte("Whatever")
}
