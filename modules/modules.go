package modules

import (
	"github.com/goEnum/goEnum/modules/CVE_2021_3156"
	"github.com/goEnum/goEnum/modules/container_block_devices"
	"github.com/goEnum/goEnum/modules/cronjobs"
	"github.com/goEnum/goEnum/modules/docker_socket"
	"github.com/goEnum/goEnum/modules/privileged_container"
	"github.com/goEnum/goEnum/modules/protected_files"
	"github.com/goEnum/goEnum/modules/readable_files"
	"github.com/goEnum/goEnum/modules/services"
	"github.com/goEnum/goEnum/modules/special_permissions"
	"github.com/goEnum/goEnum/modules/unquoted_service_path"
	"github.com/goEnum/goEnum/modules/writable_files"
	"github.com/goEnum/goEnum/structs"
)

var (
	Modules map[string]*structs.Module
	Padding int
)

func init() {
	Modules = make(map[string]*structs.Module)

	Modules["protected-files"] = protected_files.Module
	Modules["cve-2021-3156"] = CVE_2021_3156.Module
	Modules["cronjobs"] = cronjobs.Module
	Modules["readable-files"] = readable_files.Module
	Modules["writable-files"] = writable_files.Module
	Modules["special-perms"] = special_permissions.Module
	Modules["priv-container"] = privileged_container.Module
	Modules["docker-sock"] = docker_socket.Module
	Modules["block-devices"] = container_block_devices.Module
	Modules["services"] = services.Module
	Modules["unquoted-service-path"] = unquoted_service_path.Module

	padding()
}

func padding() {
	for key := range Modules {
		if len(key) > Padding {
			Padding = len(key)
		}
	}
}
