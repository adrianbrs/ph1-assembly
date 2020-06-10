package pherror

var (
	//MissingInputFile erro de arquivo de entrada não informado
	MissingInputFile = &ErrorType{
		Code:    1,
		Message: "No input file given",
	}

	//FileNotFound erro arquivo não encontrado
	FileNotFound = &ErrorType{
		Code:    2,
		Message: "File \"%s\" not found",
	}

	//CannotOpenFile erro genérico ao falhar tentando abrir um arquivo
	CannotOpenFile = &ErrorType{
		Code:    3,
		Message: "Error opening file %s",
	}

	//LabelNotFound erro de label não encontrada
	LabelNotFound = &ErrorType{
		Code:    4,
		Message: "Label \"%s\" not found",
	}

	//NoneInstructionFound erro de nome de instrução não encontrado
	NoneInstructionFound = &ErrorType{
		Code:    5,
		Message: "None instruction found for \"%s\"",
	}

	//InvalidOperandCount erro de quantidade de operadores invalidas
	InvalidOperandCount = &ErrorType{
		Code:    6,
		Message: "Invalid operand count",
	}

	//DecoratorNotFound erro retornado ao não encontrar nenhuma instrução ou
	// decorador com o nome informado
	DecoratorNotFound = &ErrorType{
		Code:    7,
		Message: "Decorator or instruction not found for \"%s\"",
	}
)
