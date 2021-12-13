package service

import (
	"context"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer ...
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()

	getR := r.Methods(http.MethodGet).Subrouter()

	guideR := getR.PathPrefix("/guidelines").Subrouter()
	guideR.Use(JSONHeader)

	guideR.Path("/documents").Handler(httptransport.NewServer(
		endpoints.GetGuidelinesDocument,
		decodeGetGuidelinesDocumentRequest,
		encodeResponse,
	))

	documentR := getR.PathPrefix("/guidelines").Subrouter()
	documentR.Use(DocumentHeader)

	documentR.Path("/documents/{filename}").HandlerFunc(documentHandler)

	return r
}

// JSONHeader ...
func JSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// DocumentHeader ...
func DocumentHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/pdf")
		next.ServeHTTP(w, r)
	})
}

func documentHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "image/jpeg")
	vars := mux.Vars(r)
	key := vars["filename"]
	url := fmt.Sprintf("assets/guidelines/documents/%s", key)

	http.ServeFile(w, r, url)
}
