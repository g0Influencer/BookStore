package app

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) (map[string]interface{}){ // возвращаем сообщение в виде объекта мапы
	return map[string]interface{}{"status":status, "message":message}
}
func Respond(w http.ResponseWriter, data map[string]interface{}){ // записываем данные в формате json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}