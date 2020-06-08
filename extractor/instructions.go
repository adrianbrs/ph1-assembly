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
	// Contador de endereço
	var address int = 0
	for _, srcLine := range contents {
		opCode, size, err := decoder.Decode(srcLine.Name)

		if err != nil {
			panic(err)
		}

		// Cria instrução sem operando
		instruction := &Instruction{
			Address: address,
			OpCode:  opCode,
		}

		// Verifica se o valor retornado da decodificação para aquela instrução é 0 ou 1
		if size == 1 {
			instruction.HasOperand = true
			// Busca nos labels o valor do operando
			operandValue := labelMap[srcLine.Operand]

			instruction.Data.Value = operandValue

			// O endereço do operando é sempre um após sua instrução, para manter a contuinidade da contagem
			// é preciso incrementar um
			instruction.Data.Address = address + 1
			address++
		}

		// Adiciona a instrução na lista e executa o laço novamente
		instructions = append(instructions, *instruction)

	}
	return
}
