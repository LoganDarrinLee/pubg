package common

import (
  "net/http"
)

// Handle error message to the client.
func HandleClientError(w http.ResponseWriter, e error) {
  switch e {

  // TODO: Actually handle error
  default: 
    w.Write([]byte(e.Error()))
  }
}
