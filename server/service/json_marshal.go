package service

import (
	"encoding/json"
	"log"
)

//JsonMarshal ...
func JsonMarshal(unit interface{}) []byte {
	smth, err := json.Marshal(unit)
	if err != nil {
		log.Println("JSON Marshal error")
	}
	return smth
}
