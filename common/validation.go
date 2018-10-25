package common

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ninjadotorg/handshake-wallet/api_response"
)

func ValidateBody(context *gin.Context, body interface{}) error {
	err := context.BindJSON(body)
	if api_response.PropagateErrorAndAbort(context, api_response.InvalidRequestBody, err) != nil {
		return err
	}

	// Validate data
	err = api_response.AbortWithRequestBodyError(context, DataValidator.Struct(body))
	if err != nil {
		return err
	}

	return nil
}

// If there is not found error, error will be returns nil
func CheckNotFound(err error) error {
	if strings.Contains(fmt.Sprintf("%s", err), "code = NotFound") {
		err = nil
	}

	return err
}
