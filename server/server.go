package server

import (
	"net/http"
	"encoding/json"
	"github.com/uvgroovy/findwords/wordmap"
)

type WordsHandler struct {
	Words wordmap.WordsMap
}

func (wh *WordsHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {    

	params := r.URL.Query()

	lettersArr, ok := params["letters"]
	if ! ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(lettersArr)!= 1 {
		w.WriteHeader(http.StatusBadRequest)
		return	
	}
	
 	w.Header().Set("Content-Type", "application/json")	
	
	letters := lettersArr[0]

	result := make([]string, 0)
	
	for word := range wh.Words.GetWords(letters) {
		result = append(result, word)
	}
	
	encoder := json.NewEncoder(w)
	
	err := encoder.Encode(result)

    if err != nil {
        panic(err.Error())
    }
	
}
