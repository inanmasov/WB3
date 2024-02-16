package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Args ...
type Args struct {
	flags   Flags
	pattern string
	files   []string
}

// Flags ...
type Flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	matchCount int
}

func main() {
	afterPtr := flag.Int("A", 0, "Print +N lines after match")
	beforePtr := flag.Int("B", 0, "Print +N lines before match")
	contextPtr := flag.Int("C", 0, "Print ±N lines around match")
	countPtr := flag.Bool("c", false, "Print count of matching lines")
	ignoreCasePtr := flag.Bool("i", false, "Ignore case")
	invertPtr := flag.Bool("v", false, "Invert match")
	fixedPtr := flag.Bool("F", false, "Fixed string match")
	lineNumPtr := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	pattern := flag.Arg(0)
	if len(flag.Args()) <= 1 {
		log.Fatalln(fmt.Errorf("укажите файл"))
	}
	files := flag.Args()[1:]

	args := &Args{
		flags: Flags{
			after:      *afterPtr,
			before:     *beforePtr,
			context:    *contextPtr,
			count:      *countPtr,
			ignoreCase: *ignoreCasePtr,
			invert:     *invertPtr,
			fixed:      *fixedPtr,
			lineNum:    *lineNumPtr,
		},
		pattern: pattern,
		files:   files,
	}

	res, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range res {
		printLines(v)
	}
}

func processFile(args *Args) ([][]string, error) {
	var res [][]string
	for _, file := range args.files {
		lines, err := scanFile(file)
		if err != nil {
			return nil, err
		}
		filteredLines := grep(args, lines)
		res = append(res, filteredLines)
	}

	return res, nil
}

func scanFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func printLines(filteredLines []string) {
	for _, line := range filteredLines {
		fmt.Println(line)
	}
}

func grep(args *Args, lines []string) []string {
	var filteredLines []string

	for i, line := range lines {
		// flag i and F
		match := findPattern(args, line)

		if (args.flags.invert && !match) || (!args.flags.invert && match) {
			if args.flags.count {
				args.flags.matchCount++
			} else {
				if args.flags.after > 0 {
					filteredLines = append(filteredLines, flagLineNum(args, line, i, ":"))
					flagAfter(args, &filteredLines, lines, args.flags.after, i)
				} else if args.flags.before > 0 {
					flagBefore(args, &filteredLines, lines, args.flags.before, i)
					filteredLines = append(filteredLines, flagLineNum(args, line, i, ":"))
				} else if args.flags.context > 0 {
					flagBefore(args, &filteredLines, lines, args.flags.context, i)
					filteredLines = append(filteredLines, flagLineNum(args, line, i, ":"))
					flagAfter(args, &filteredLines, lines, args.flags.context, i)
				} else {
					filteredLines = append(filteredLines, flagLineNum(args, line, i, ":"))
				}
			}
		}
	}

	if args.flags.count {
		filteredLines = append(filteredLines, strconv.Itoa(args.flags.matchCount))
	}

	return filteredLines
}

func findPattern(args *Args, line string) bool {
	var match bool
	pattern := args.pattern
	if args.flags.ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if args.flags.fixed {
		return strings.Contains(line, args.pattern)
	}

	re := regexp.MustCompile(pattern)
	match = re.MatchString(line)

	return match
}

func flagAfter(args *Args, filteredLines *[]string, lines []string, count, i int) {
	for j := i + 1; j < len(lines); j++ {
		if count <= 0 {
			*filteredLines = append(*filteredLines, "--")
			break
		}
		if findPattern(args, lines[j]) {
			break
		}
		*filteredLines = append(*filteredLines, flagLineNum(args, lines[j], j, "-"))
		count--
	}

}

func flagBefore(args *Args, filteredLines *[]string, lines []string, count, i int) {
	var res []string
	for j := i - 1; j >= 0; j-- {
		if count <= 0 {
			res = append(res, "--")
			break
		}
		if findPattern(args, lines[j]) {
			break
		}

		res = append(res, flagLineNum(args, lines[j], j, "-"))
		count--
	}

	for j := len(res) - 1; j >= 0; j-- {
		*filteredLines = append(*filteredLines, res[j])
	}
}

func flagLineNum(args *Args, line string, i int, de string) string {
	if args.flags.lineNum {
		return fmt.Sprintf("%d%s%s", i+1, de, line)
	}
	return line
}
