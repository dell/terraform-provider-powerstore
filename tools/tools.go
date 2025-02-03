//go:build tools

/*
Copyright (c) 2025 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://mozilla.org/MPL/2.0/


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// getCopyrightYear get copyright year from git log and file modify time.
func getCopyrightYear(filePath string) (string, error) {
	// Sanitize the filePath
	filePath = filepath.Clean(filePath)
	currYear := fmt.Sprintf("%d", time.Now().Year())
	cmd := exec.Command("bash", "-c", "git log --follow --format=%cd --date=format:%Y "+filePath+" | sort -u") // #nosec G204
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(filePath, "git-log: (", string(output), ") ")
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	// if newly created
	if lines[0] == "" {
		return currYear, nil
	}
	startYear := lines[0]
	if len(lines) == 1 {
		return startYear, nil
	}
	endYear := lines[len(lines)-1]
	return startYear + "-" + endYear, nil
}

type Facts struct {
	Note       string
	ExampleVar string
}

// main function to traveser docs folder and update copyright year.
func main() {

	err := filepath.Walk("docs/", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		path = filepath.Clean(path)
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		year, err := getCopyrightYear(path)
		if err != nil {
			return err
		}

		replacedFile := strings.Replace(string(file), "<copyright-year>", year, 1)

		println("Copyright Years: " + year + " " + path)

		// get the file and directory name by splitting by '/' and getting last 2
		pathHierarchy := strings.Split(path, "/")
		fileName := strings.TrimSuffix(pathHierarchy[len(pathHierarchy)-1], ".md")
		dirName := pathHierarchy[len(pathHierarchy)-2]

		var fnote, exampleEnding string
		// if dir is datasource
		if dirName == "data-sources" {

			fmt.Println("Got filename: " + fileName)
			// if note exist
			if note, ok := datasourceFacts[fileName]; ok {
				// add note
				if note.Note != "" {
					fnote = "\n\n" + note.Note
				}
				// add example var
				if note.ExampleVar != "" {
					exampleEnding = "\n\nAfter the successful execution of above said block," +
						" We can see the output by executing `terraform output` command." +
						" Also, we can fetch information via the variable: `" +
						note.ExampleVar +
						"` where attribute_name is the attribute which user wants to fetch."
				}
			}
		}

		// if dir is resource
		if dirName == "resources" {
			// if note exist
			if note, ok := resourceFacts[fileName]; ok {
				// add note
				if note.Note != "" {
					fnote = "\n\n" + note.Note
				}
				// add example var
				if note.ExampleVar != "" {
					exampleEnding = "\n\nAfter the execution of above resource block, " +
						note.ExampleVar +
						" would have been created on the PowerStore array. For more information, Please check the terraform state file."
				}
			}
		}

		// replace <note>
		replacedFile = strings.ReplaceAll(replacedFile, "<note>", fnote)
		// replace <example-ending>
		replacedFile = strings.ReplaceAll(replacedFile, "<example-ending>", exampleEnding)

		err = os.WriteFile(path, []byte(replacedFile), 0600)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
}
