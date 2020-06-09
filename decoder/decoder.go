package decoder

import (
	"fmt"
	"ph1-assembly/constants"
)

type metaInstruction struct {
	opCode string
	size   int
}

// Mapeia todos os mnemônicos disponíveis para seus opCodes e número de endereços
var operations = map[string]*metaInstruction{
	"NOP": &metaInstruction{opCode: "00", size: 1},
	"LDR": &metaInstruction{opCode: "10", size: 2},
	"STR": &metaInstruction{opCode: "20", size: 2},
	"ADD": &metaInstruction{opCode: "30", size: 2},
	"SUB": &metaInstruction{opCode: "40", size: 2},
	"MUL": &metaInstruction{opCode: "50", size: 2},
	"DIV": &metaInstruction{opCode: "60", size: 2},
	"NOT": &metaInstruction{opCode: "70", size: 2},
	"AND": &metaInstruction{opCode: "80", size: 1},
	"OR":  &metaInstruction{opCode: "90", size: 2},
	"XOR": &metaInstruction{opCode: "A0", size: 2},
	"JMP": &metaInstruction{opCode: "B0", size: 2},
	"JEQ": &metaInstruction{opCode: "C0", size: 2},
	"JG":  &metaInstruction{opCode: "D0", size: 2},
	"JL":  &metaInstruction{opCode: "E0", size: 2},
	"HLT": &metaInstruction{opCode: "F0", size: 1},
}

// Decode traduz o mnemônico de uma instrução e retorna
// seu opcode e tamanho
func Decode(name string) (string, int, error) {
	instruction := operations[name]

	if instruction == nil {
		return "", 0, fmt.Errorf(constants.NoneInstructionFound, name)
	}

	return instruction.opCode, instruction.size, nil
}
