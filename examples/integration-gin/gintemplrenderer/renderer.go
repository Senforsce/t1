package gintemplrenderer

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin/render"
	"github.com/senforsce/tndr"
)

var Default = &Renderer{}

func New(ctx context.Context, status int, component t1.Component) *Renderer {
	return &Renderer{
		Ctx:       ctx,
		Status:    status,
		Component: component,
	}
}

type Renderer struct {
	Ctx       context.Context
	Status    int
	Component t1.Component
}

func (t Renderer) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Status)
	if t.Component != nil {
		return t.Component.Render(t.Ctx, w)
	}
	return nil
}

func (t Renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *Renderer) Instance(name string, data any) render.Render {
	t1Data, ok := data.(t1.Component)
	if !ok {
		return nil
	}
	return &Renderer{
		Ctx:       context.Background(),
		Status:    http.StatusOK,
		Component: t1Data,
	}
}
