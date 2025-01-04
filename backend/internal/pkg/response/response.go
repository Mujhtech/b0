package response

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
)

type Response struct {
	StatusCode int `json:"-"`
}

type ServerResponse struct {
	Response
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
}

func (res Response) Render(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	render.Status(r, res.StatusCode)
	return nil
}

func BadRequest(w http.ResponseWriter, r *http.Request, err error) error {
	_ = render.Render(w, r, ServerResponse{
		Response: Response{
			StatusCode: http.StatusBadRequest,
		},
		Message: "Bad Request",
		Error:   err.Error(),
	})

	return nil
}

func InternalServerError(w http.ResponseWriter, r *http.Request, err error) error {
	_ = render.Render(w, r, ServerResponse{
		Response: Response{
			StatusCode: http.StatusInternalServerError,
		},
		Message: "Internal Server Error",
		Error:   err.Error(),
	})

	return nil
}

func Ok(w http.ResponseWriter, r *http.Request, message string, data interface{}) error {
	_ = render.Render(w, r, ServerResponse{
		Response: Response{
			StatusCode: http.StatusOK,
		},
		Message: message,
		Data:    data,
	})

	return nil
}

func Created(w http.ResponseWriter, r *http.Request, message string, data interface{}) error {
	_ = render.Render(w, r, ServerResponse{
		Response: Response{
			StatusCode: http.StatusCreated,
		},
		Message: message,
		Data:    data,
	})

	return nil
}

func Unauthorized(w http.ResponseWriter, r *http.Request, err error) error {
	_ = render.Render(w, r, ServerResponse{
		Response: Response{
			StatusCode: http.StatusUnauthorized,
		},
		Message: "Unauthorized",
		Error:   err.Error(),
	})

	return nil
}

func Redirect(w http.ResponseWriter, r *http.Request, u string, status int, ignoreUrl bool) error {
	if !ignoreUrl {
		parsedUrl, err := url.ParseRequestURI(u)
		if err != nil || parsedUrl.Host != "" {
			return BadRequest(w, r, fmt.Errorf("invalid redirect URL"))
		}
	}

	http.Redirect(w, r, u, status)
	return nil
}
