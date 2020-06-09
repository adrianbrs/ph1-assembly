package extractor

import (
	"ph1-assembly/constants"
	"ph1-assembly/decoder"
	"ph1-assembly/input"
	"strconv"
)

//Instruction representa uma instrução
type Instruction struct {
	Address    int
	OpCode     string
	HasOperand bool
	Data       *Data
}

//Data representa o valor que a operação carrega
type Data struct {
	Address int
	Value   int
}

//ExtractInstructions efetua a segunda passagem no código, guardando as
// instrucoes em uma lista de struct
func ExtractInstructions(contents []*input.SourceLine, labelMap map[string]int) []Instruction {
	var instructions = make([]Instruction, 0)

	for _, srcLine := range contents {

		if srcLine.Name == constants.TextSection || srcLine.Name == constants.DataSection {
			break
		}

		opCode, size, err := decoder.Decode(srcLine.Name)

		if err != nil {
			panic(err)
		}

		// Cria instrução sem operando
		instruction := &Instruction{
			Address: srcLine.Address,
			OpCode:  opCode,
		}

		// Verifica se o valor retornado da decodificação para aquela instrução é 1 ou 2
		if size == 2 {
			instruction.HasOperand = true
			// Busca nos labels o valor do operando
			operandValue, found := labelMap[srcLine.Operand]

			if found == false {
				operandValue, err = strconv.Atoi(srcLine.Operand)

				if err != nil {
					panic(constants.LabelNotFound)
				}
			}

			instruction.Data.Value = operandValue

			instruction.Data.Address = srcLine.Address + 1
		}

		// Adiciona a instrução na lista e executa o laço novamente
		instructions = append(instructions, *instruction)

	}

	return instructions
}
