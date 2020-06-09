package extractor

import (
	"ph1-assembly/decoder"
	"ph1-assembly/input"
)

// ExtractLabels efetua a primeira passagem no código, guardando os rótulos
// e seus endereços em um map que irá retornar
func ExtractLabels(contents []*input.SourceLine) (labelMap map[string]int) {

	sections := map[string]int{
		"text": 0,
		"data": 128,
	}
	currentSection := ""

	for _, srcLine := range contents {
		if _, ok := sections[srcLine.Name]; ok {
			currentSection = srcLine.Name
			continue
		}

		srcLine.Address = sections[currentSection]

		_, size, _ := decoder.Decode(srcLine.Name)

		if size == 0 {
			size = 1
		}
		sections[currentSection] += size

		// Adiciona o label no map se a label não for vazia
		if srcLine.Label != "" {
			labelMap[srcLine.Label] = srcLine.Address
		}
	}

	return
}
