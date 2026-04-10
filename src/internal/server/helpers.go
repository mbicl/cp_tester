package server

import (
	"fmt"
	"os"
)

func createFolders(path string) error {
	// Create tests folder if it doesn't exist
	if _, err := os.Stat(path + "/.tests"); os.IsNotExist(err) {
		err := os.MkdirAll(path+"/.tests", os.ModePerm)
		if err != nil {
			return err
		}
	}
	// Create bin folder if it doesn't exist
	if _, err := os.Stat(path + "/.bin"); os.IsNotExist(err) {
		err := os.MkdirAll(path+"/.bin", os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveTests(tests []Test, path, problemId string) error {
	path += "/.tests"
	for i, test := range tests {
		inputFile := fmt.Sprintf("%s/%s%02d.in", path, problemId, i)
		ansFile := fmt.Sprintf("%s/%s%02d.ans", path, problemId, i)
		if err := os.WriteFile(inputFile, []byte(test.Input), 0644); err != nil {
			return err
		}
		if err := os.WriteFile(ansFile, []byte(test.Output), 0644); err != nil {
			return err
		}
	}
	return nil
}

func createSolutionFile(config *Config, problemContent *ProblemContent, problemId string) error {
	solutionFile := fmt.Sprintf("%s/%s.%s", config.CPPath, problemId, config.Language)
	// return if the solution file already exists
	if _, err := os.Stat(solutionFile); err == nil {
		return nil
	}
	var tmpFileByte []byte
	if _, err := os.Stat(config.Template); err == nil {
		tmpFileByte, err = os.ReadFile(config.Template)
		if err != nil {
			return err
		}
	}

	tmpHeader := GetTemplate(config.Language, problemContent)
	return os.WriteFile(solutionFile, append([]byte(tmpHeader+"\n\n"), tmpFileByte...), 0644)

}
