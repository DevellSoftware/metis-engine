package workspace

import (
	"github.com/DevellSoftware/metis-engine/pkg/log"
	"github.com/DevellSoftware/metis-engine/pkg/module"
)

type Workspace struct {
	tracer  *log.Tracer
	modules []module.Module
}

func NewWorkspace(tracer *log.Tracer) *Workspace {
	return &Workspace{
		tracer:  tracer,
		modules: make([]module.Module, 0),
	}
}

func (w *Workspace) AddModule(module module.Module) {
	w.modules = append(w.modules, module)
}
