package input

import (
	"bufio"
	"errors"
	"os"
	"ph1-assembly/constants"
	"ph1-assembly/decoder"
	"ph1-assembly/pherror"
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
	Label      string
	Name       string
	Operand    string
	LineNumber int
	Address    int
}

// Errorf retorna uma instância de ErrorType contendo as informações da linha
func (line *SourceLine) Errorf(msg interface{}, args ...interface{}) *pherror.ErrorType {
	pherr := &pherror.ErrorType{LineNumber: line.LineNumber}
	if errType, ok := msg.(*pherror.ErrorType); ok {
		pherr = pherror.Join(pherr, errType)
	} else if errMsg, ok := msg.(string); ok {
		pherr.Message = errMsg
	}
	return pherror.Format(pherr, args...)
}

// Source contém as informações do código fonte
type Source struct {
	Filename           string
	Text               []*SourceLine
	Data               []*SourceLine
	CurrentTextAddress int
	CurrentDataAddress int
}

// AppendText adiciona uma nova linha na seção text, atribuindo
// seu endereço e validando a instrução
func (src *Source) AppendText(line *SourceLine) {
	_, size := decoder.DecodeText(line.Name)
	line.Address = src.CurrentTextAddress
	src.Text = append(src.Text, line)
	src.CurrentTextAddress += size
}

// AppendData adiciona uma nova linha na seção data, atribuindo
// seu endereço e validando o tipo de dado
func (src *Source) AppendData(line *SourceLine) {
	size := decoder.DecodeData(line.Name)
	line.Address = src.CurrentDataAddress
	src.Data = append(src.Data, line)
	src.CurrentDataAddress += size
}

// ReadSource lê o codigo fonte e transforma para o tipo Source
func ReadSource(filename string) (source *Source) {
	contents, err := read(filename)

	// Verifica se houve algum erro na leitura do arquivo
	if err != nil {
		var perr *os.PathError
		if errors.As(err, &perr) {
			panic(pherror.Format(pherror.FileNotFound, filename))
		} else {
			panic(pherror.Format(pherror.CannotOpenFile, filename))
		}
	}

	source = &Source{
		Filename:           filename,
		CurrentTextAddress: constants.TextSectionAddress,
		CurrentDataAddress: constants.DataSectionAddress,
	}

	// Seção atual do código (text ou data)
	var section string

	// Cria um erro relacionado ao arquivo
	fileError := &pherror.ErrorType{Filename: filename}

	// Lẽ cada linha salvanda na seção atual
	for lineNumber, line := range contents {
		// Atualiza a linha atual do erro de arquivo
		fileError.LineNumber = lineNumber + 1

		// Faz uma validação previa da linha
		if validateSourceLine(line) {

			// Extrai os dados da linha atual
			sourceLine := parseSourceLine(line)
			sourceLine.LineNumber = lineNumber + 1

			// Verifica a seção atual
			if strings.ToUpper(sourceLine.Name) == strings.ToUpper(constants.TextSection) {
				section = constants.TextSection
			} else if strings.ToUpper(sourceLine.Name) == strings.ToUpper(constants.DataSection) {
				section = constants.DataSection
			} else if section == constants.TextSection {
				source.AppendText(sourceLine)
			} else if section == constants.DataSection {
				source.AppendData(sourceLine)
			} else {
				panic(sourceLine.Errorf(pherror.DecoratorNotFound, sourceLine.Name))
			}
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
		panic(pherror.InvalidOperandCount)
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
