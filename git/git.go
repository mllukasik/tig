package git

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile("^([*])?\\s{1,2}(\\S*)\\s*(\\S*)\\s(\\[\\S*\\])?\\s?(.*)$")

func PruneBranchAll() {
	fmt.Println("Removing current branch is not supported yet")
	PruneBranch()
}

func PruneBranch() {
	branches := GetBranches()
	for _, branch := range branches {
		if branch.Current {
			continue
		}
		DeleteBranch(branch.Name)
	}
}

func DeleteBranch(branch string) {
	remote := false
	if strings.HasPrefix(branch, "remotes/") {
		remote = true
		branch = branch[8:]
	}

	args_length := 3
	args := [4]string{"branch", "-D", branch}
	if remote {
		args[3] = "--remote"
		args_length = 4
	}

	cmd := exec.Command("git", args[:args_length]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

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
	cmd := exec.Command("git", "branch", "-a", "-vv")
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(stdout)), nil
}

func parseBranchString(value string) []BranchName {
	lines := strings.Split(value, "\n")
	branches := make([]BranchName, len(lines))
	for index, line := range lines {
		branches[index] = parseBranchLine(line)
	}
	return branches
}

func parseBranchLine(value string) BranchName {
	groups := pattern.FindStringSubmatch(value)
	return BranchName{
		Name:              groups[2],
		Current:           len(groups[1]) == 1,
		RemoteTracking:    len(groups[4]) != 0,
		RemoteName:        groups[4],
		LastCommitMessage: groups[3],
		LastCommitHash:    groups[5],
	}
}

type BranchName struct {
	Name           string
	Current        bool
	RemoteTracking bool
	RemoteName     string
	//todo: create commit struct
	LastCommitMessage string
	LastCommitHash    string
}
