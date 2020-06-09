package input

import (
	"bufio"
	"os"
	"ph1-assembly/constants"
	"regexp"
	"strings"
)

// Regex utilizado para extrair os valores de uma linha de código assembly
var (
	lineMatchRegex = regexp.MustCompile(`^(?:(?:(\w+):)\s+)?(\w+)(?:\s+(\w+))?$`)
	commentRegex   = regexp.MustCompile(`([;].*)`)
	spacesRegex    = regexp.MustCompile(`\s+`)
)

// SourceLine contém os dados necessários de uma linha do código fonte
type SourceLine struct {
	Label   string
	Name    string
	Operand string
	Address int
}

// Source contém as informações do código fonte
type Source struct {
	Filename string
	Contents []*SourceLine
}

// ReadSource lê o codigo fonte e transforma para o tipo Source
func ReadSource(filename string) (source *Source, err error) {
	contents, err := read(filename)

	if err != nil {
		return
	}

	source = &Source{Filename: filename}

	for _, line := range contents {
		if validateSourceLine(line) {
			source.Contents = append(source.Contents, parseSourceLine(line))
		}
	}

	return
}

func validateSourceLine(line string) bool {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, ";") {
		return false
	}
	return true
}

func parseSourceLine(line string) (srcLine *SourceLine) {
	line = commentRegex.ReplaceAllString(line, "")
	line = spacesRegex.ReplaceAllString(line, " ")
	line = strings.TrimSpace(line)
	match := lineMatchRegex.FindStringSubmatch(line)

	if len(match) == 0 {
		panic(constants.InvalidOperandCount)
	}

	// Instantcia um novo SourceLine
	srcLine = &SourceLine{
		Label:   match[1],
		Name:    match[2],
		Operand: match[3],
	}

	return
}

func read(filename string) (contents []string, err error) {
	file, err := os.Open(filename)

	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return
}
