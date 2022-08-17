package protected_files

import (
	"github.com/goEnum/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Protected Files",
		Prereqs,
		Enumeration,
		Report,
	)
}
