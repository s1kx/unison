package client

import (
	"strings"

	"github.com/s1kx/unison"
)

type commandRegistry map[string]*unison.Command

// Check if a message content string is a valid command by it's prefix "!" or bot mention
func identifiesAsCommand(content, prefix string) (status bool, updatedContent string) {
	if strings.HasPrefix(content, prefix) {
		// remove duplicate command prefixes
		result := removePrefix(content, prefix)
		for {
			if !strings.HasPrefix(result, prefix) {
				break
			} else {
				result = removePrefix(result, prefix)
			}
		}

		// make sure there is content after removing the command prefix
		if len(result) > 0 {
			return true, result
		}
		return false, result
	}

	return false, content
}

// Removes a substring from the string and cleans up leading & trailing spaces.
func removePrefix(str, remove string) string {
	result := strings.TrimPrefix(str, remove)
	result = strings.TrimSpace(result)

	if str[0] == ' ' {
		return result[1:len(result)]
	}
	return result
}
