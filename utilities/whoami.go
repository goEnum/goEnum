package utilities

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/goEnum/goEnum/structs"
	"github.com/goEnum/goEnum/utilities/color"
)

func Whoami() *structs.Whoami {
	whoami := structs.NewWhoami()

	currUser(whoami)
	home(whoami)
	path(whoami)
	groups(whoami)

	return whoami
}

func groups(whoami *structs.Whoami) {
	groups, err := exec.Command("id").Output()
	if err == nil {
		id := strings.Split(strings.TrimSpace(string(groups)), " ")
		whoami.UID = strings.Split(id[0], "=")[1]
		whoami.GID = strings.Split(id[1], "=")[1]
		whoami.Groups = strings.Split(id[2][7:], ",")
	}
}

func path(whoami *structs.Whoami) {
	whoami.Path = os.Getenv("PATH")
}

func home(whoami *structs.Whoami) {
	whoami.Home = os.Getenv("HOME")
}

func currUser(whoami *structs.Whoami) {
	user, err := exec.Command("whoami").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	whoami.User = strings.TrimSpace(string(user))
}

func WhoamiString(whoami *structs.Whoami) string {
	var output bytes.Buffer

	color.Fprintf(color.Bold, &output, "====== User ======\n")

	if whoami.User != "" {

		fmt.Fprintf(&output, "User:   %v\n", whoami.User)
	}

	if whoami.Home != "" {
		fmt.Fprintf(&output, "Home:   %v\n", whoami.Home)
	}

	if whoami.Path != "" {
		fmt.Fprintf(&output, "Path:   %v\n", whoami.Path)
	}

	if whoami.UID != "" {
		fmt.Fprintf(&output, "UID:    %v\n", whoami.UID)
	}

	if whoami.GID != "" {
		fmt.Fprintf(&output, "GID:    %v\n", whoami.GID)
	}

	if len(whoami.Groups) != 0 {
		fmt.Fprint(&output, ListPrint("Groups", whoami.Groups))
	}

	return output.String()
}
