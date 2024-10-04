package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PushCurrent(remote string) {
	current := GetCurrentBranch()
	cmd := exec.Command("git", "push", remote, current)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		message := fmt.Errorf("Could not push current branch. Cause: %s", err.Error())
		fmt.Println(message)
		os.Exit(1)
	}
}

func GetCurrentBranch() string {
	cmd := exec.Command("git", "branch", "--show-current")
	stdout, err := cmd.Output()
	if err != nil {
		message := fmt.Errorf("Could not fetch current branch. Cause: %s", err.Error())
		fmt.Println(message)
		os.Exit(1)
	}
	return strings.TrimSpace(string(stdout))
}

func GetBranches() []BranchName {
	stdout, err := getBranchesCommand()
	if err != nil {
		return []BranchName{{
			Name:    err.Error(),
			Current: false,
		}}
	}
	return parseBranchString(stdout)
}

func getBranchesCommand() (string, error) {
	cmd := exec.Command("git", "branch", "-a")
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}

func parseBranchString(value string) []BranchName {
	parts := strings.Split(value, "\n")
	branches := make([]BranchName, len(parts))
	for index, part := range parts {
		name := part
		current := strings.Contains(name, "*")
		if current {
			name = name[strings.Index(name, "*")+1:]
		}
		branches[index] = BranchName{
			Name:    strings.TrimSpace(name),
			Current: current,
		}

	}
	return branches
}

type BranchName struct {
	Name    string
	Current bool
}
