package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/beevik/etree"
)

func findChangedFiles() []string {
	output, err := exec.Command("git", "status", "-s").Output()
	if err != nil {
		log.Fatal(err)
	}

	var changed_files []string

	for changed_file := range strings.Lines(string(output)) {
		parts := strings.Split(strings.TrimSpace(changed_file), " ")

		if parts[0] == "M" && (strings.Contains(parts[1], "/") || strings.Contains(parts[1], "\\")) {
			changed_files = append(changed_files, parts[1])
		}
	}

	return changed_files
}

func getArtifactName(file string) string {
	doc := etree.NewDocument()
	err := doc.ReadFromFile(file)
	if err != nil {
		log.Fatal(err)
	}
	project := doc.SelectElement("project")
	return project.SelectElement("artifactId").Text()
}

func determineModule(file string) (*string, error) {
	var pomFile string
	index := strings.Index(file, "src")
	if index != -1 {
		pomFile = filepath.Join(file[:index-1], "pom.xml")
		if isModule(pomFile) {
			artifactId := getArtifactName(pomFile)
			return &artifactId, nil
		}
	}

	return nil, exec.ErrNotFound
}

func isModule(file string) bool {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return !strings.Contains(string(content), "<packaging>pom</packaging>")
}

func main() {

	changed_files := findChangedFiles()
	var modules string
	for _, file := range changed_files {

		pom, err := determineModule(file)
		if err == nil {
			if len(modules) == 0 {
				modules = ":" + *pom
			} else {
				modules = modules + ",:" + *pom
			}
		}
	}

	for _, mod := range modules {
		fmt.Println(mod)
	}

	var arguments = os.Args[1:]
	arguments = append(arguments, "-pl", modules)
	cmd := exec.Command("mvn", arguments...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	os.Stdout.ReadFrom(stdout)
}
