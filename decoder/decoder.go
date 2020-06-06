package decoder

type Operation struct {
	Code string
	Size int
}

var OperationTable = map[string]*Operation{
	"LDR": &Operation{"10", 1},
}

// Decode traduz o mnemônico de uma instrução e retorna
// seu opcode e tamanho
func Decode(name string) (op *Operation) {
	return
}
