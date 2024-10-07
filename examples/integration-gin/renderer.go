package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/senforsce/tndr"
)

type TemplRender struct {
	Code int
	Data t1.Component
}

func (t TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Data != nil {
		return t.Data.Render(context.Background(), w)
	}
	return nil
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *TemplRender) Instance(name string, data interface{}) render.Render {
	if t1Data, ok := data.(t1.Component); ok {
		return &TemplRender{
			Code: http.StatusOK,
			Data: t1Data,
		}
	}
	return nil
}
