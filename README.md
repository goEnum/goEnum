<img src="/assets/goenum_banner.png">
<h1  align="center" style="border-bottom">
  Modular and System-Agnostic Enumeration Framework
</h1>

<div align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/goEnum/goEnum?color=green">
  <img src="https://img.shields.io/github/release/goEnum/goEnum.svg">
  <img src="https://img.shields.io/github/license/goEnum/goEnum">
  <img src="https://img.shields.io/github/stars/goEnum/goEnum?color=yellow">
  <img src="https://img.shields.io/github/contributors-anon/goEnum/goEnum">
  <img src="https://img.shields.io/github/issues-raw/goEnum/goEnum">
  <img src="https://img.shields.io/github/downloads/goEnum/goEnum/total">
</div>


## Usage

goEnum is a standalone CLI tools which no dependancies, this means all you will ever need it the binary itself

goEnum also has a robust help interface (thanks to [Cobra](https://github.com/spf13/cobra)!) for if you have any questions on what goEnum is doing

### Examples
 `goEnum --help`

 ```
System-Agnostic and Modular Enumeration Framework by Maxwell Fusco

Usage:
  goEnum [flags]
  goEnum [command]

Available Commands:
  all         run all available modules
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  modules     display all available modules
  none        runs no modules
  ssh         execute goEnum over ssh remote connection

Flags:
  -f, --format string   output format [json, markdown]
  -h, --help            help for goEnum
  -c, --no-color        disable color output
  -o, --output string   output file
  -v, --verbose         verbose output

Use "goEnum [command] --help" for more information about a command.
 ```
 
***

`goEnum modules`

```
====== Modules ======

[+] services              => Insecure Services and Utilized Binaries
[+] unquoted-service-path => Unquoted Service Paths
[+] cve-2021-3156         => CVE-2021-3156
[+] writable-files        => Mispermissioned Files (readable)
[+] special-perms         => SUID and GUID Files
[+] priv-container        => Priviledged Container
[+] block-devices         => Block Devices in Containers
[+] protected-files       => Protected Files
[+] cronjobs              => Cronjobs with Writable Executable
[+] readable-files        => Mispermissioned Files (readable)
[+] docker-sock           => Container with Docker Socket
```

***

`goEnum cve-2021-3156`

```
====== CVE-2021-3156 ======
[+] Prereqs: Passed
[+] Enumeration: Succeeded
[*] Reporting: Skipping

====== Report ======

=== CVE-2021-315 ===
Description: Vulnerable version of sudo allowing for Heap-Based Buffer Overflow Privelege Escalation
Files: /snap/bin/sudo
     - /snap/bin/sudoedit
```
** Results for modules will vary based on the system executed on

## Contributing

Pull requests, forks, and issues are all welcome! Make sure to make open a new branch, pull request, and issues for any submitted changes and they will be reviewed!

## License

[BSD-3](https://opensource.org/licenses/BSD-3-Clause)