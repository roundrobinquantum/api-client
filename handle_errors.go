package api_client

import (
	"github.com/roundrobinquantum/api-client/errors"
)

const prefix = "API-ERROR"

var (
	NotFoundError       = errors.DefineError(prefix, 1, "Account not found.")
)

var StatusCodes errors.StatusCodeList = map[string]errors.StatusCode{
	NotFoundError.Code():       {Error: NotFoundError, StatusCode: 404, ErrorCode: 1},
}