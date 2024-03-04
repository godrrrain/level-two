package tools

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type port struct {
	Port string `json:"port"`
}

// ReadConfig считывает конфиг
func ReadConfig() string {
	defaultPort := "8080"
	data, err := os.ReadFile("config.json")
	if err != nil {
		return defaultPort
	}
	port := port{}
	err = json.Unmarshal(data, &port)
	if err != nil {
		return defaultPort
	}
	return port.Port
}

// RequestLogger оборачивает запрос для его вывода в логи
func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(
			"method", r.Method,
			"path", r.URL.EscapedPath(),
		)
		next(w, r)
	}
}
