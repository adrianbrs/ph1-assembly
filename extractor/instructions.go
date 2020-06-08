package extractor

import (
	"ph1-assembly/decoder"
	"ph1-assembly/input"
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
func ExtractInstructions(contents []*input.SourceLine, labelMap map[string]int) (instructions []Instruction) {
	for index, srcLine := range contents {
		opCode, size, err := decoder.Decode(srcLine.Name)

		if err != nil {
			panic(err)
		}

		instruction := &Instruction{
			Address: index,
			OpCode:  opCode,
		}

		if size == 1 {
			instruction.HasOperand = true
			operandValue := labelMap[srcLine.Operand]

			instruction.Data.Value = operandValue
			instruction.Data.Address = index + 1
			index++
		}

		instructions = append(instructions, *instruction)

	}
	return
}
