package proxy

import (
	"sync"

	lsp "github.com/a-h/protocol"
)

func NewDiagnosticCache() *DiagnosticCache {
	return &DiagnosticCache{
		m:     &sync.Mutex{},
		cache: make(map[string]fileDiagnostic),
	}
}

type fileDiagnostic struct {
	t1Diagnostics    []lsp.Diagnostic
	goplsDiagnostics []lsp.Diagnostic
}

type DiagnosticCache struct {
	m     *sync.Mutex
	cache map[string]fileDiagnostic
}

func zeroLengthSliceIfNil(diags []lsp.Diagnostic) []lsp.Diagnostic {
	if diags == nil {
		return make([]lsp.Diagnostic, 0)
	}
	return diags
}

func (dc *DiagnosticCache) AddTemplDiagnostics(uri string, goDiagnostics []lsp.Diagnostic) []lsp.Diagnostic {
	goDiagnostics = zeroLengthSliceIfNil(goDiagnostics)
	dc.m.Lock()
	defer dc.m.Unlock()
	diag := dc.cache[uri]
	diag.goplsDiagnostics = goDiagnostics
	diag.t1Diagnostics = zeroLengthSliceIfNil(diag.t1Diagnostics)
	dc.cache[uri] = diag
	return append(diag.t1Diagnostics, goDiagnostics...)
}

func (dc *DiagnosticCache) ClearTemplDiagnostics(uri string) {
	dc.m.Lock()
	defer dc.m.Unlock()
	diag := dc.cache[uri]
	diag.t1Diagnostics = make([]lsp.Diagnostic, 0)
	dc.cache[uri] = diag
}

func (dc *DiagnosticCache) AddGoDiagnostics(uri string, t1Diagnostics []lsp.Diagnostic) []lsp.Diagnostic {
	t1Diagnostics = zeroLengthSliceIfNil(t1Diagnostics)
	dc.m.Lock()
	defer dc.m.Unlock()
	diag := dc.cache[uri]
	diag.t1Diagnostics = t1Diagnostics
	diag.goplsDiagnostics = zeroLengthSliceIfNil(diag.goplsDiagnostics)
	dc.cache[uri] = diag
	return append(diag.goplsDiagnostics, t1Diagnostics...)
}
