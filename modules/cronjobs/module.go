package cronjobs

import (
	"github.com/goEnum/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Cronjobs with Writable Executable",
		Prereqs,
		Enumeration,
		Report,
	)
}
