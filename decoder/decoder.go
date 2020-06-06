package decoder

import "log"

type metaInstruction struct {
	opCode string
	size   int
}

// Mapeia todos os mnemônicos disponíveis para seus opCodes e número de endereços
var operations = map[string]*metaInstruction{
	"NOP": &metaInstruction{opCode: "00", size: 0},
	"LDR": &metaInstruction{opCode: "01", size: 1},
	"STR": &metaInstruction{opCode: "02", size: 1},
	"ADD": &metaInstruction{opCode: "03", size: 1},
	"SUB": &metaInstruction{opCode: "04", size: 1},
	"MUL": &metaInstruction{opCode: "05", size: 1},
	"DIV": &metaInstruction{opCode: "06", size: 1},
	"NOT": &metaInstruction{opCode: "07", size: 1},
	"AND": &metaInstruction{opCode: "08", size: 0},
	"OR":  &metaInstruction{opCode: "09", size: 1},
	"XOR": &metaInstruction{opCode: "A0", size: 1},
	"JMP": &metaInstruction{opCode: "B0", size: 1},
	"JEQ": &metaInstruction{opCode: "C0", size: 1},
	"JG":  &metaInstruction{opCode: "D0", size: 1},
	"JL":  &metaInstruction{opCode: "E0", size: 1},
	"HLT": &metaInstruction{opCode: "F0", size: 0},
}

// Decode traduz o mnemônico de uma instrução e retorna
// seu opcode e tamanho
func Decode(name string) (string, int) {
	instruction := operations[name]

	if instruction == nil {
		log.Panicf("None instruction found for this name: %s", name)
	}

	return instruction.opCode, instruction.size
}
