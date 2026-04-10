package server

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var slashTemplate = `
// Created by cp_tester on %s
// Author: %s
// Problem: %s
// URL: %s
`

var hashTemplate = `
# Created by cp_tester on %s
# Author: %s
# Problem: %s
# URL: %s
`

var languageTemplates = map[string]string{
	"c":    slashTemplate,
	"cpp":  slashTemplate,
	"cs":   slashTemplate,
	"go":   slashTemplate,
	"java": slashTemplate,
	"rs":   slashTemplate, // Rust uses line comments with slashes
	"js":   slashTemplate, // JavaScript uses line comments with slashes
	"ts":   slashTemplate, // TypeScript uses line comments with slashes
	"hs":   hashTemplate,  // Haskell uses hash comments
	"rb":   hashTemplate,  // Ruby uses hash comments
	"py":   hashTemplate,  // Python uses hash comments
}

func GetTemplate(language string, problemContent *ProblemContent) string {
	if template, exists := languageTemplates[language]; exists {
		// Get OS username as author
		author := "Unknown"
		if runtime.GOOS == "windows" {
			author = fmt.Sprintf("%s\\%s", os.Getenv("USERDOMAIN"), os.Getenv("USERNAME"))
		} else {
			author = os.Getenv("USER")
		}
		return fmt.Sprintf(template, time.Now().Format("2006-01-02 15:04:05"), author, problemContent.Name, problemContent.URL)
	}
	return ""
}
