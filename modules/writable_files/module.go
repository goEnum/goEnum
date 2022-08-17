package writable_files

import (
	"github.com/goEnum/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Mispermissioned Files (readable)",
		Prereqs,
		Enumeration,
		Report,
	)
}
