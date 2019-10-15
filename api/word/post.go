package word

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/4Tune-Light/test-2/controllers/word/test0"
)

type wordType struct {
    Word string
}

type wordRes struct {
	Vowels int
	Consonants int
}

func DoPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
  var word wordType

  err := decoder.Decode(&word)
  if err != nil {
      fmt.Printf("Word is required")
      w.WriteHeader(http.StatusBadRequest)
      return
  }

  data := test0.Test0(word.Word)

  res := wordRes{data["vowels"], data["consonants"]}

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(res)
}