package decoder

import "log"

type instruction struct {
	opCode string
	size   int
}

// Mapeia todos os mnemônicos disponíveis para seus opCodes e número de endereços
var operations = map[string]*instruction{
	"NOP": &instruction{opCode: "00", size: 0},
	"LDR": &instruction{opCode: "01", size: 1},
	"STR": &instruction{opCode: "02", size: 1},
	"ADD": &instruction{opCode: "03", size: 1},
	"SUB": &instruction{opCode: "04", size: 1},
	"MUL": &instruction{opCode: "05", size: 1},
	"DIV": &instruction{opCode: "06", size: 1},
	"NOT": &instruction{opCode: "07", size: 1},
	"AND": &instruction{opCode: "08", size: 0},
	"OR":  &instruction{opCode: "09", size: 1},
	"XOR": &instruction{opCode: "A0", size: 1},
	"JMP": &instruction{opCode: "B0", size: 1},
	"JEQ": &instruction{opCode: "C0", size: 1},
	"JG":  &instruction{opCode: "D0", size: 1},
	"JL":  &instruction{opCode: "E0", size: 1},
	"HLT": &instruction{opCode: "F0", size: 0},
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
