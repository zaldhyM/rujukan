package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func OKList(c *gin.Context, data any, total, limit, offset int) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"total":   total,
		"limit":   limit,
		"offset":  offset,
		"data":    data,
	})
}

func OKMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": message})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": data})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"success": false, "message": message})
}

func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message)
}

func Conflict(c *gin.Context, message string) {
	Error(c, http.StatusConflict, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message)
}

func InternalError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message)
}

// ValidationError formats binding/validation errors into field-level messages.
func ValidationError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		fields := make(map[string]string, len(ve))
		for _, fe := range ve {
			fields[strings.ToLower(fe.Field())] = fieldMessage(fe)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validasi input gagal",
			"errors":  fields,
		})
		return
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		BadRequest(c, "Format JSON tidak valid")
		return
	}

	var unmarshalErr *json.UnmarshalTypeError
	if errors.As(err, &unmarshalErr) {
		BadRequest(c, fmt.Sprintf("Tipe data tidak sesuai untuk field '%s'", unmarshalErr.Field))
		return
	}

	BadRequest(c, err.Error())
}

func fieldMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "wajib diisi"
	case "min":
		return fmt.Sprintf("minimal %s karakter", fe.Param())
	case "max":
		return fmt.Sprintf("maksimal %s karakter", fe.Param())
	case "len":
		return fmt.Sprintf("harus tepat %s karakter", fe.Param())
	case "email":
		return "harus berupa alamat email yang valid"
	case "oneof":
		return fmt.Sprintf("harus salah satu dari: %s", strings.ReplaceAll(fe.Param(), " ", ", "))
	case "numeric":
		return "harus berupa angka"
	case "alpha":
		return "hanya boleh berisi huruf"
	case "alphanum":
		return "hanya boleh berisi huruf dan angka"
	}
	return fmt.Sprintf("validasi '%s' gagal", fe.Tag())
}
