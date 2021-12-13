package service

import (
	"context"
	"encoding/json"
	"net/http"
	"sadhelx-be-guidelines/util"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type (
	// Endpoints define all endpoint
	Endpoints struct {
		GetGuidelinesDocument endpoint.Endpoint
		GetDocument           endpoint.Endpoint
	}

	// Response format
	Response struct {
		Status  bool        `json:"status"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data,omitempty"`
	}

	// GetDocumentReq ...
	GetDocumentReq struct {
		Filename string `json:"filename"`
	}
)

// MakeGuidelinesEndpoints ...
func MakeGuidelinesEndpoints(svc Service) Endpoints {
	return Endpoints{
		GetGuidelinesDocument: makeGetGuidelinesDocument(svc),
		GetDocument:           makeGetDocument(svc),
	}
}

func makeGetGuidelinesDocument(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(EmailAvailabilityReq)

		res, err := svc.GetGuidelinesDocument(ctx)
		if err != nil {
			return Response{Status: false, Message: err.Error()}, nil
		}
		return Response{Status: true, Message: util.MsgGetDocument, Data: res}, nil
	}
}

func decodeGetGuidelinesDocumentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	return r, nil
}

func makeGetDocument(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(EmailAvailabilityReq)

		req := request.(GetDocumentReq)

		err := svc.GetDocument(ctx, req.Filename)
		if err != nil {
			return Response{Status: false, Message: err.Error()}, nil
		}
		return Response{Status: true, Message: "File fetched"}, nil
	}
}

func decodeGetDocumentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req GetDocumentReq
	params := mux.Vars(r)
	filename := params["filename"]
	req.Filename = filename
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(Response)
	sc := util.StatusCode(res.Message)
	if sc == 0 {
		sc = 500
	}
	w.WriteHeader(sc)
	return json.NewEncoder(w).Encode(&res)
}
