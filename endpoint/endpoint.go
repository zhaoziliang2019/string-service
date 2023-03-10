package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/zhaoziliang2019/string-service/service"
	"strings"
)

type StringEndpoints struct {
	StringEndpoint      endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
}

var ErrInvalidRequestType = errors.New("RequestType has only towtype:Concat,Diff")

type StringRequest struct {
	RequestType string `json:"request_type"`
	A           string `json:"a"`
	B           string `json:"b"`
}
type StringResponse struct {
	Result string `json:"result"`
	Error  error  `json:"error"`
}

func MakeStringEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(StringRequest)
		var (
			res, a, b string
			opError   error
		)
		a = req.A
		b = req.B
		if strings.EqualFold(req.RequestType, "Concat") { //不区分大小写
			res, _ = svc.Concat(a, b)
		} else if strings.EqualFold(req.RequestType, "Diff") {
			res, _ = svc.Diff(a, b)
		} else {
			return nil, ErrInvalidRequestType
		}
		return StringResponse{
			Result: res,
			Error:  opError,
		}, nil
	}
}

type HealthResponse struct {
	Status bool `json:"status"`
}
type HealthRequest struct {
}

func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthResponse{status}, nil
	}
}
