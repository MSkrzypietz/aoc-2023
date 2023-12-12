package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	genCmd := flag.NewFlagSet("gen", flag.ExitOnError)
	genDay := genCmd.Int("day", 1, "day")

	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	execDay := execCmd.Int("day", 1, "day")

	switch os.Args[1] {
	case "gen":
		if err := genCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal("error parsing the cli args: ", err)
		}
		generateDay(*genDay)
	case "exec":
		if err := execCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal("error parsing the cli args: ", err)
		}
		executeDay(*execDay)
	default:
		fmt.Println("expected 'gen' or 'exec' subcommands")
		os.Exit(1)
	}
}

func generateDay(day int) {
	fmt.Printf("Generating files for day %d...\n", day)

	dir := fmt.Sprintf("day%02d", day)
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		log.Fatal("error creating the day's directory: ", err)
	}

	source, err := os.Open("template/template.go")
	if err != nil {
		log.Fatal("error opening the template file: ", err)
	}
	defer source.Close()

	destination, err := os.Create(fmt.Sprintf("%s/code.go", dir))
	if err != nil {
		log.Fatal("error creating the destination file: ", err)
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		log.Fatal("error copying the template file to its destination: ", err)
	}

	input, err := os.Create(fmt.Sprintf("%s/input.txt", dir))
	if err != nil {
		log.Fatal("error creating the input.txt file: ", err)
	}
	defer input.Close()

	fmt.Printf("Successfully generated files for day %d!\n", day)
	fmt.Printf("Paste the contents of https://adventofcode.com/2023/day/%d/input into %s/input.txt\n", day, dir)
}

func executeDay(day int) {
	cmd := exec.Command("go", "run", fmt.Sprintf("day%02d/code.go", day))
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("DAY=%02d", day))

	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal("error getting the output from the executed day: ", err)
	}
	fmt.Println(string(stdout))
}
