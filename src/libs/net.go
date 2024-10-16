package libs

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// RequestError represents the parameter error in the HTTP request
type requestError struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Follow the same semantics of cornice
type errorResponse struct {
	Status string         `json:"status"`
	Errors []requestError `json:"errors"`
}

type successResponse struct {
	DataType string      `json:"type"`
	Result   interface{} `json:"solution"`
}

// Request object represents the base application request
type Request struct {
	errors      []requestError
	request     *http.Request
	response    http.ResponseWriter
	errorStatus int
	validator   *validator.Validate
	payloadType string
	payloadData interface{}
}

// NewRequest creates a new request
func NewRequest(res http.ResponseWriter, req *http.Request) *Request {
	appReq := &Request{
		request:   req,
		response:  res,
		validator: validator.New(),
	}
	return appReq
}

// HasErrors returns true if there are any errors in the request
func (req *Request) HasErrors() bool {
	return len(req.errors) > 0
}

// AddError adds the errors that the request has encountered
func (req *Request) AddError(status int, location, name, desc string) {
	req.errorStatus = status
	err := requestError{Location: location, Name: name, Description: desc}
	req.errors = append(req.errors, err)
}

// WriteErrorResponse writes error response to the supplied http.ResponseWriter
func (req *Request) WriteErrorResponse() {
	req.response.Header().Set("Content-Type", "application/json")
	if req.errorStatus == 0 {
		req.response.WriteHeader(http.StatusInternalServerError)
	} else {
		req.response.WriteHeader(req.errorStatus)
	}
	if len(req.errors) <= 0 {
		return
	}
	resp := errorResponse{Status: "error", Errors: req.errors}
	response, err := json.Marshal(resp)
	if err != nil {
		// We can't do much about it here
		log.Println("Error while constructing json error response")
	} else {
		req.response.Write(response)
	}
}

// Validate will deserialize the data from the http request and set
// appropriate errors if there are any errors on the data and return
// whether the validation of data was successful or not.
func (req *Request) Validate(i interface{}) bool {
	// For POST requests
	err := json.NewDecoder(req.request.Body).Decode(i)
	if err != nil {
		req.AddError(http.StatusBadRequest, "body", "decode", "Unable to decode json body")
		return false
	}
	vErr := req.validator.Struct(i)
	if vErr == nil {
		return true
	}
	var validationErrors validator.ValidationErrors
	errors.As(vErr, &validationErrors)
	for _, err := range validationErrors {
		req.AddError(http.StatusBadRequest, "body",
			fmt.Sprintf("%s", err.Field()),
			fmt.Sprintf("Validation failed with tag: %s", err.Tag()))
	}
	return false
}

// AddResponsePayload adds the payload data for success response
func (req *Request) AddResponsePayload(retType string, data interface{}) {
	req.payloadType = retType
	req.payloadData = data
}

// WriteSuccessResponse writes the success response and attaches
// payload if there are any.
func (req *Request) WriteSuccessResponse() {
	req.response.Header().Set("Content-Type", "application/json")
	req.response.WriteHeader(http.StatusOK)
	if req.payloadData == nil {
		return
	}
	resp := successResponse{DataType: req.payloadType, Result: req.payloadData}
	response, err := json.Marshal(resp)
	if err != nil {
		// We can't do much about it here
		log.Println("Error while construction json success response")
	} else {
		req.response.Write(response)
	}
}
