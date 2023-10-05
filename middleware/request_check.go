package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckBody(next echo.HandlerFunc, validation interface{}) echo.HandlerFunc {
	return func(echo echo.Context) error {
		// Read original request body
		var bodyBytes []byte
		if echo.Request().Body != nil {
			bodyBytes, _ = io.ReadAll(echo.Request().Body)
			// Write back to request body
			echo.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			// Try to parse json data
			reqData := validation
			err := json.Unmarshal(bodyBytes, &reqData)
			if err != nil {
				return echo.JSON(http.StatusBadRequest, "error json.")
			}
		}
		return next(echo)
	}
}
