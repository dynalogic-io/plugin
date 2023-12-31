package v1

import (
	"github.com/dynalogic-io/plugin/v1/core"
)

type Plugin interface {
	New(core core.Core)
	Info() core.Manifest
	Process(value interface{}) interface{}
}
