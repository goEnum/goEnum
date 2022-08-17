package container_block_devices

import (
	"github.com/goEnum/goEnum/structs"
)

var Module *structs.Module

func init() {
	Module = structs.NewModule(
		"Block Devices in Containers",
		Prereqs,
		Enumeration,
		Report,
	)
}
