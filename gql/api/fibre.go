package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"
)


func QueryService() (string, string, error) {
	service_host := os.Getenv("FIBRE_HOST")
    if service_host == "" {
        service_host = "http://127.0.0.1"
    }

    str_host := fmt.Sprintf("%s:3000/", service_host)

	resp, err := http.Get(str_host)
	if err != nil {
		return "", "", err
	}

	var data map[string]interface{}
	errs := json.NewDecoder(resp.Body).Decode(&data)
	if errs != nil {
		log.Fatal(errs)
	}

	return data["ID"].(string), data["Msg"].(string), nil
}
