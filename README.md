# go_exercises
A workspace to practice writing small programs in Go

## Structure

New exercises should be implemented in their own subdirectories. Each subdirectory should contain its own `main` package and `main` func. 

Programs should be run from the root directory via `go run ./subdir_name`

## Contents

- `/render_html_template`: passing a struct to an HTML template and serving the rendered page.
- `/parse_yaml_frontmatter`: an implementation of how to parse YAML frontmatter from a markdown document.
- `/reader_interface`: Working through how to read in the contents of a file by using the Reader interface.
- `/simple_web_server`: A super simple way to implement a hello world web server.