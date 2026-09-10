package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const inp = "parse_yaml_frontmatter/demo.md"

func main() {
	f, err := os.Open(inp)

	if err != nil {
		log.Fatalf("Couldn't open file: %s", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var yamlCont []string
	var separatorCount int
	inFrontmatter := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "---" {
			separatorCount++
			if separatorCount == 1 {
				inFrontmatter = true
				continue
			} else if separatorCount == 2 {
				break
			}
		}

		if inFrontmatter {
			yamlCont = append(yamlCont, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error occured while scanning: %s", err)
	}

	fmt.Print(yamlCont)
}
