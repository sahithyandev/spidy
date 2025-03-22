package crawler

import (
	"strings"
)

// ParseRobotsTxt parses the robots.txt file and returns a map of disallowed URLs
// for the given robotName.
func ParseRobotsTxt(body string, robotName string) []string {
	disallowedUrls := []string{}
	cursorIndex := 0

	currentKey := ""
	currentValue := ""
	currentTarget := "key"
	isForRobot := false

	normalizedRobotName := strings.ToLower(robotName)
	body = body + "\n"
	for cursorIndex < len(body) {
		switch body[cursorIndex] {
		case '\n':
			if (currentKey == "" && currentValue == "") || (currentKey == "user-agent" && currentValue == "*" && isForRobot) {
				isForRobot = false
			} else if currentKey == "user-agent" &&
				(currentValue == normalizedRobotName || (len(disallowedUrls) == 0 && currentValue == "*")) {
				isForRobot = true
			} else if isForRobot && currentKey == "disallow" {
				disallowedUrls = append(disallowedUrls, currentValue)
			}

			currentKey = ""
			currentValue = ""
			currentTarget = "key"
		case ':':
			cursorIndex += 2
			currentTarget = "value"
			continue
		default:
			if currentTarget == "key" {
				currentKey += strings.ToLower(string(body[cursorIndex]))
				if currentKey == "#" {
					// it's a comment. skip to next line
					for cursorIndex < len(body) && body[cursorIndex] != '\n' {
						currentKey = ""
						cursorIndex++
					}
				}
			} else {
				currentValue += strings.ToLower(string(body[cursorIndex]))
			}
		}
		cursorIndex++
	}

	return disallowedUrls
}
