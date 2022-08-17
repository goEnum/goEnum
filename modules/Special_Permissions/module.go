package special_permissions

import (
	"github.com/goEnum/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"SUID and GUID Files",
		Prereqs,
		Enumeration,
		Report,
	)
}
