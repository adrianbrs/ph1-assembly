package pherror

import (
	"fmt"
	"ph1-assembly/constants"
)

var inputFileName string

// Setup inicializa os valores para os logs de erro
func Setup(filename string) {
	inputFileName = filename
}

// ErrorType é utilizado para erros relacionados ao montador
type ErrorType struct {
	Message    string
	Code       int
	Filename   string
	LineNumber int
}

// ErrorType método da interface de erro
func (err *ErrorType) Error() string {
	var msg string = "Error"
	var args []interface{}

	// Verifica se existe um código de erro
	if err.Code != 0 {
		msg += fmt.Sprintf(" [%%0%dd]", constants.MaxErrorCodeDigits)
		args = append(args, err.Code)
	}

	// Verifica se há um nome de arquivo no erro
	if err.Filename != constants.Empty || err.LineNumber != 0 {
		var filename string

		if err.Filename != constants.Empty {
			filename = err.Filename
		} else if inputFileName != constants.Empty {
			filename = inputFileName
		}

		if filename != constants.Empty {
			msg += " in %s"
			args = append(args, filename)

			// Verifica se há um número de linha informado no erro
			if err.LineNumber != 0 {
				msg += ":%d"
				args = append(args, err.LineNumber)
			}
		}
	}

	// Formata a mensagem e retorna o texto formatado do erro
	args = append(args, err.Message)
	return fmt.Sprintf(msg+": %s", args...)
}

// Join unifica vários ErrorType em uma única cópia de ErrorType
func Join(errTypes ...*ErrorType) *ErrorType {
	if len(errTypes) == 0 {
		return nil
	}

	retErr := *errTypes[0]
	for _, err := range errTypes[1:] {
		if err.Code != 0 {
			retErr.Code = err.Code
		}
		if err.Filename != constants.Empty {
			retErr.Filename = err.Filename
		}
		if err.LineNumber != 0 {
			retErr.LineNumber = err.LineNumber
		}
		if err.Message != constants.Empty {
			retErr.Message = err.Message
		}
	}

	return &retErr
}

// Format formata qualquer erro em um ErrorType
func Format(msg interface{}, args ...interface{}) *ErrorType {
	var errType ErrorType

	// Verifica se o erro passado é do tipo ErrorType
	if pherr, ok := msg.(*ErrorType); ok {
		if len(args) == 0 {
			return pherr
		}
		errType = *pherr
		errType.Message = fmt.Sprintf(errType.Message, args...)
	} else if err, ok := msg.(error); ok {
		// Cria um novo ErrorType a partir do erro genérico
		errType = ErrorType{Message: fmt.Sprintf(err.Error(), args...)}
	} else if msgstr, ok := msg.(string); ok {
		// Formata a string se o erro for uma string
		errType = ErrorType{Message: fmt.Sprintf(msgstr, args...)}
	} else {
		// Define como mensagem a representação do valor informado
		errType = ErrorType{Message: fmt.Sprintf("%#v", msg)}
	}
	return &errType
}
