package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/ninjadotorg/handshake-wallet/api_response"
)

type ContextServiceValidator interface {
	CheckNotFound() bool
	CheckValidate() string
	CheckError() error
	GetStatusKey() string
	ContextValidate(context *gin.Context) bool
}

type SimpleContextError struct {
	NotFound  bool
	Error     error
	StatusKey string
}

func (r SimpleContextError) CheckNotFound() bool {
	return r.NotFound
}

func (r SimpleContextError) CheckValidate() string {
	return r.StatusKey
}

func (r SimpleContextError) CheckError() error {
	return r.Error
}

func (r SimpleContextError) GetStatusKey() string {
	return r.StatusKey
}

func (r SimpleContextError) ContextValidate(context *gin.Context) (invalid bool) {
	if r.CheckNotFound() {
		api_response.AbortNotFound(context)
		return true
	}
	if err := r.CheckError(); err != nil {
		api_response.PropagateErrorAndAbort(context, r.GetStatusKey(), err)
		return true
	}
	if statusKey := r.CheckValidate(); statusKey != "" {
		api_response.AbortWithValidateErrorSimple(context, statusKey)
		return true
	}

	return
}

func (r *SimpleContextError) SetErrorOnly(err error) bool {
	// Only set to error and status key if there is really error
	if err != nil {
		r.StatusKey = api_response.UnexpectedError
		r.Error = err
		return true
	}
	return false
}

func (r *SimpleContextError) SetStatusKey(statusKey string) {
	r.SetError(statusKey, errors.New(statusKey))
}

func (r *SimpleContextError) SetError(statusKey string, err error) bool {
	// Only set to error and status key if there is really error
	if err != nil {
		r.StatusKey = statusKey
		r.Error = err
		return true
	}
	return false
}

func (r *SimpleContextError) FeedContextErrorDefault(object ContextServiceValidator) bool {
	return r.FeedContextError(object.CheckValidate(), object)
}

func (r *SimpleContextError) FeedContextError(statusKey string, object ContextServiceValidator) bool {
	r.NotFound = object.CheckNotFound()
	return r.SetError(statusKey, object.CheckError())
}

func (r *SimpleContextError) HasError() bool {
	if r.CheckNotFound() || r.CheckValidate() != "" || r.CheckError() != nil {
		return true
	}
	return false
}

func GetError(validator ContextServiceValidator) (string, error) {
	if validator.CheckNotFound() {
		return api_response.ResourceNotFound, nil
	}
	if statusKey := validator.CheckValidate(); statusKey != "" {
		return statusKey, nil
	}
	if err := validator.CheckError(); err != nil {
		return validator.GetStatusKey(), err
	}

	return "", nil
}
