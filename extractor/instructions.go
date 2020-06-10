package extractor

import (
	"ph1-assembly/constants"
	"ph1-assembly/decoder"
	"ph1-assembly/input"
	"ph1-assembly/pherror"
	"strconv"
)

//Instruction representa uma instrução
type Instruction struct {
	Address    int
	OpCode     string
	HasOperand bool
	Data       *Data
}

// Data representa os valores contidos na section data ou representa os valores
// contidos no segundo byte da instrução, dependendo do contexto
type Data struct {
	Address int
	Value   int
}

//ExtractInstructionsAndData efetua a segunda passagem no código, guardando as
// instrucoes em uma lista de struct e os dados da section Data na struct Data
func ExtractInstructionsAndData(content *input.Source, labelMap map[string]int) (*[]Instruction, *[]Data) {
	instructions := extractInstructions(content, labelMap)
	datas := extractData(content.Data)

	return instructions, datas
}

// Extrai as instruções e cria uma lista de instructions contendo endereço, opCode e os dados do segundo
// byte(caso tenha)
func extractInstructions(content *input.Source, labelMap map[string]int) *[]Instruction {
	var instructions = make([]Instruction, 0)

	dataContent := content.Data
	textContent := content.Text
	for _, srcLine := range textContent {

		if srcLine.Name == constants.TextSection || srcLine.Name == constants.DataSection {
			continue
		}

		opCode, size := decoder.DecodeText(srcLine.Name)

		// Cria instrução sem operando
		instruction := Instruction{
			Address: srcLine.Address,
			OpCode:  opCode,
		}

		// Verifica se o valor retornado da decodificação para aquela instrução é 1 ou 2
		if size == 2 {
			extractDataLabel(dataContent, instruction, labelMap, srcLine)
		}

		// Adiciona a instrução na lista e executa o laço novamente
		instructions = append(instructions, instruction)

	}

	return &instructions
}

// Procura no label map o valor do operando
func extractDataLabel(dataContent []*input.SourceLine, instruction Instruction, labelMap map[string]int,
	instructionInfo *input.SourceLine) {
	instruction.HasOperand = true

	// Busca nos labels o valor do operando
	operandValue, found := labelMap[instructionInfo.Operand]

	// Caso não encontre
	if found == false {
		var err error
		operandValue, err = strconv.Atoi(instructionInfo.Operand)

		if err != nil {
			panic(pherror.Format(pherror.LabelNotFound, instructionInfo.Operand))
		}
	}

	data := &Data{
		Value:   operandValue,
		Address: instructionInfo.Address + 1,
	}
	instruction.Data = data
}

func extractData(content []*input.SourceLine) *[]Data {
	var datas = make([]Data, 0)

	for _, data := range content {
		newData := Data{
			Value:   data.Address,
			Address: data.LineNumber,
		}

		datas = append(datas, newData)
	}

	return &datas
}
