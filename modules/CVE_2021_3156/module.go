package CVE_2021_3156

import (
	"github.com/goEnum/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"CVE-2021-3156",
		Prereqs,
		Enumeration,
		Report,
	)
}
