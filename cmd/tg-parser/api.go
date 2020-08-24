package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getMe(clientStorage *ClientStorage) {
	http.HandleFunc("/v1/getMe", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		res, _ := clientStorage.SendAndCatch(Update{
			"@type": "getMe",
		})

		result, _ := json.Marshal(res)
		_, _ = fmt.Fprintln(w, string(result))
	})
}
