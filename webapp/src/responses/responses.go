package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a resposta de erro da API
type APIErro struct {
	Erro string `json:"erro"`
}

// JSON retorna um resposta em formato JSON para a requisicao
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// HandleErrorStatusCode
func HandleErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var erro APIErro
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
