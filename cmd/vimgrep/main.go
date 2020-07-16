package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("File extension and search term are required")
		fmt.Println("vimgrep [FILE EXTENSION] SEARCH TERM")
		os.Exit(1)
	}
	ending := os.Args[1]

	searchterm := strings.Join(os.Args[2:], " ")

	location := "."
	cmd := exec.Command("grep", "-H", "-n", "-R", "-I", fmt.Sprintf("--include=*.%s", ending), searchterm, location)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Grep command failed: %s\n", cmd.String())
		fmt.Println("This is probably because the string could not be found")
		os.Exit(1)
	}
	list := append([]string{"Done"}, strings.Split(strings.Trim(string(out), " \n"), "\n")...)
	if len(list) == 1 {
		fmt.Println("No result found")
		os.Exit(1)
	}
	var selected []string
	for i := 0; i < 20; i++ {
		prompt := promptui.Select{
			Label: "Select file to open",
			Items: list,
			Size:  40,
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Prompt failed.")
			os.Exit(1)
		}

		if result == "Done" {
			break
		}

		for i, v := range list {
			if v == result {
				list = append(list[:i], list[i+1:]...)
				break
			}
		}

		parts := strings.Split(result, ":")
		if len(parts) < 2 {
			continue
		}

		selected = append(selected, fmt.Sprintf("%s:%s", parts[0], parts[1]))
		if len(list) == 1 {
			break
		}
	}
	if len(selected) == 0 {
		os.Exit(0)
	}
	selected = append([]string{"-p"}, selected...)
	cmd = exec.Command("vim", selected...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error logging in to AWS, exiting")
		os.Exit(1)
	}

}
