// A minimal mock of the REST Countries v3.1 API (https://restcountries.com),
// which was retired. It only knows a few countries (see countries.json) and
// answers like the original API: an array with the country, or a 404 error.
package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed countries.json
var countriesData []byte

func main() {
	countries := map[string]json.RawMessage{}
	if err := json.Unmarshal(countriesData, &countries); err != nil {
		log.Fatalf("cannot parse countries.json: %s", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v3.1/alpha/{code}", func(w http.ResponseWriter, r *http.Request) {
		code := strings.ToUpper(r.PathValue("code"))
		country, ok := countries[code]
		if !ok {
			log.Printf("GET %s -> 404 (unknown country %q)", r.URL.Path, code)
			writeJSON(w, http.StatusNotFound, []byte(`{"status":404,"message":"Not Found"}`))
			return
		}
		log.Printf("GET %s -> 200", r.URL.Path)
		// The v3.1 API always returns a list, even for a single country
		writeJSON(w, http.StatusOK, append(append([]byte("["), country...), ']'))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("REST Countries mock listening on :%s with %d countries", port, len(countries))
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func writeJSON(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}
