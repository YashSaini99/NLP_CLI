package validator

import (
	"errors"
	"regexp"
	"strings"
)

// Base allowed commands.
var allowedBaseCommands = []string{
	"ls", "tar", "gzip", "zip", "cp", "mv", "df", "ps", "ping", "curl",
	"cat", "less", "more", "head", "tail", "grep", "echo", "pwd", "whoami",
	"date", "uptime", "uname", "which", "man", "wc", "sort", "uniq", "diff",
	"comm", "cut", "paste", "tr", "du", "free", "top", "htop", "dig", "nslookup",
	"host", "traceroute", "mtr", "netstat", "ss", "whois", "sleep", "time",
}

// List of allowed command patterns.
var allowedCommands = []string{}

func init() {
	// Add patterns for base commands.
	for _, cmd := range allowedBaseCommands {
		allowedCommands = append(allowedCommands, `^(Command:\s*)?`+cmd+`(\s|$)`)
	}
	// Add pattern for find with -exec using allowed commands.
	execPattern := `^(Command:\s*)?find\s+.*-exec\s+(` + strings.Join(allowedBaseCommands, "|") + `)\s+.*`
	allowedCommands = append(allowedCommands, execPattern)
}

// ValidateCommand checks if the command is allowed.
func ValidateCommand(cmd string) error {
	for _, pattern := range allowedCommands {
		matched, err := regexp.MatchString(pattern, cmd)
		if err != nil {
			return err
		}
		if matched {
			return nil
		}
	}
	return errors.New("command not allowed by whitelist")
}
