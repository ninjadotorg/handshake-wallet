package api_response

import "net/http"

const Success = "Success"
const UnexpectedError = "UnexpectedError"
const ResourceNotFound = "ResourceNotFound"
const TokenInvalid = "TokenInvalid"
const FirebaseError = "FirebaseError"
const SendEmailError = "SendEmailError"
const InvalidRequestBody = "InvalidRequestBody"
const InvalidRequestParam = "InvalidRequestParam"
const InvalidQueryParam = "InvalidQueryParam"
const ExternalApiFailed = "ExternalApiFailed"
const InvalidNumber = "InvalidNumber"
const InvalidConfig = "InvalidConfig"

const GetDataFailed = "GetDataFailed"
const AddDataFailed = "AddDataFailed"
const UpdateDataFailed = "UpdateDataFailed"
const DeleteDataFailed = "DeleteDataFailed"

const InvalidGiftCardCode = "InvalidGiftCardCode"
const ExpiredGiftCardCode = "ExpiredGiftCardCode"
const RedeemGiftCardCodeFailed = "RedeemCodeFailed"
const GiftCardCodeRedeemed = "GiftCardCodeRedeemed"
const BlockNotFound = "BlockNotFound"

var CodeMessage = map[string]struct {
	StatusCode int
	Code       int
	Message    string
}{
	Success:             {http.StatusOK, 1, "Success"},
	UnexpectedError:     {http.StatusInternalServerError, -1, "Unexpected error"},
	ResourceNotFound:    {http.StatusNotFound, -1, "Resource not found"},
	FirebaseError:       {http.StatusInternalServerError, -1, "Unexpected error"},
	SendEmailError:      {http.StatusInternalServerError, -1, "Unexpected error"},
	TokenInvalid:        {http.StatusUnauthorized, -3, "Token is invalid"},
	InvalidRequestBody:  {http.StatusBadRequest, -4, "Request body is invalid"},
	InvalidRequestParam: {http.StatusBadRequest, -5, "Request param is invalid"},
	InvalidQueryParam:   {http.StatusBadRequest, -6, "Query param is invalid"},
	ExternalApiFailed:   {http.StatusBadRequest, -7, "External API failed"},
	InvalidNumber:       {http.StatusBadRequest, -8, "Invalid number"},
	InvalidConfig:       {http.StatusBadRequest, -9, "Invalid config"},

	GetDataFailed:    {http.StatusBadRequest, -201, "Get data failed"},
	AddDataFailed:    {http.StatusBadRequest, -202, "Add data failed"},
	UpdateDataFailed: {http.StatusBadRequest, -203, "Update data failed"},
	DeleteDataFailed: {http.StatusBadRequest, -204, "Delete data failed"},

	InvalidGiftCardCode:      {http.StatusBadRequest, -300, "Invalid code or code is used"},
	ExpiredGiftCardCode:      {http.StatusBadRequest, -301, "Code has been expired"},
	RedeemGiftCardCodeFailed: {http.StatusBadRequest, -302, "Could not redeem this code at this time"},
	GiftCardCodeRedeemed:     {http.StatusBadRequest, -303, "This code has been redeemed"},

	BlockNotFound: {http.StatusBadRequest, -600, "Last block not found"},
}
