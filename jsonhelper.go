package main
import (
	"encoding/json"
	"net/http"
	"log"
)

func respond(w http.ResponseWriter,code int,payload interface{}){
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to response")
		w.WriteHeader(500)
		return
	}	

	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
} 

func respondWithError(w http.ResponseWriter,code int,message string){

	if code > 499 {
		log.Println("Responding with 5xx error",message)
	}
	type errResponse struct{
		Error string `json:"error"` // i want key for the field as `error`
	}

	respond(w,code,errResponse{Error:message})
}