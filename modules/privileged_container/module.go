package privileged_container

import (
	"github.com/goEnum/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Priviledged Container",
		Prereqs,
		Enumeration,
		Report,
	)
}
