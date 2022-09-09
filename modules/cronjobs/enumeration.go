//go:build !windows

package cronjobs

import (
	"github.com/goEnum/goEnum/utilities"
)

func Enumeration() ([]string, bool) {
	var files []string

	utilities.Crontab("/etc/crontab", &files)
	utilities.CrontabDirectory("/etc/cron.d/", &files)
	utilities.ShellDirectory("/etc/cron.daily/", &files)
	utilities.ShellDirectory("/etc/cron.hourly/", &files)
	utilities.ShellDirectory("/etc/cron.monthly/", &files)
	utilities.ShellDirectory("/etc/cron.weekly/", &files)

	return files, len(files) != 0
}
