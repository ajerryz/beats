package main

import (
	"github.com/google/uuid"
	"log/slog"
)

type HttpRequest struct {
	Protocol string `json:"protocol"`
	Method   string `json:"method"`
	Host     string `json:"host"`
	Uri      string `json:"uri"`
	Body     any    `json:"body"`
}

type HttpResponse struct {
	Headers map[string]any `json:"headers"`
	Body    any            `json:"body"`
}

var getHttpRequest = HttpRequest{
	Protocol: "HTTP/1.1",
	Method:   "GET",
	Host:     "www.demo.com",
	Uri:      "/v1/hello",
	Body:     nil,
}

var getHttpResponse = HttpResponse{
	Headers: map[string]any{
		"Content-Type":   "application/json",
		"Content-Length": 10,
	},
	Body: map[string]interface{}{
		"data": "hello world",
	},
}

func printGetLog(logger *slog.Logger) {
	reqId := uuid.New().String()
	logger.Info("this is a info message", "requestId", reqId, "request", getHttpRequest, "response", getHttpResponse)
}

var postHttpRequest = HttpRequest{
	Protocol: "HTTP/1.1",
	Method:   "POST",
	Host:     "www.demo.com",
	Uri:      "/v1/hello",
	Body: map[string]string{
		"name": "zhangning",
	},
}

var postHttpResponse = HttpResponse{
	Headers: map[string]any{
		"Content-Type":   "application/json",
		"Content-Length": 10,
	},
	Body: map[string]interface{}{
		"code": 0,
		"msg":  "success",
	},
}

func printPostLog(logger *slog.Logger) {
	reqId := uuid.New().String()
	logger.Info("this is a post request info message",
		"requestId", reqId,
		"request", postHttpRequest,
		"response", postHttpResponse)
}

var putHttpRequest = HttpRequest{
	Protocol: "HTTP/1.1",
	Method:   "PUT",
	Host:     "www.demo.com",
	Uri:      "/v1/hello",
	Body:     nil,
}

var putHttpResponse = HttpResponse{
	Headers: map[string]any{
		"Content-Type":   "application/json",
		"Content-Length": 10,
	},
	Body: map[string]interface{}{
		"code": 0,
		"msg":  "success",
	},
}

func printPutLog(logger *slog.Logger) {
	reqId := uuid.New().String()
	logger.Info("this is a put request info message",
		"requestId", reqId,
		"request", putHttpRequest,
		"response", putHttpResponse)
}

var deleteHttpRequest = HttpRequest{
	Protocol: "HTTP/1.1",
	Method:   "DELETE",
	Host:     "www.test.com",
	Uri:      "/v1/demo-123",
	Body:     nil,
}

var deleteHttpResponse = HttpResponse{
	Headers: map[string]any{
		"Content-Type":   "application/json",
		"Content-Length": 10,
	},
	Body: map[string]interface{}{
		"code": 0,
		"msg":  "success",
	},
}

func printDeleteLog(logger *slog.Logger) {
	reqId := uuid.New().String()
	logger.Info("this is a delete request info message",
		"requestId", reqId,
		"request", deleteHttpRequest,
		"response", deleteHttpResponse)
}
