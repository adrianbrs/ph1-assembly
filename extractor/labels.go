package extractor

import (
	"ph1-assembly/input"
)

// ExtractLabels efetua a primeira passagem no código, guardando os rótulos
// e seus endereços em um map que irá retornar
func ExtractLabels(src *input.Source) map[string]int {
	labels := map[string]int{}

	// Encontra as labels
	for _, srcText := range src.Text {
		if srcText.Label != "" {
			labels[srcText.Label] = srcText.Address
		}
	}

	for _, srcData := range src.Data {
		if srcData.Label != "" {
			labels[srcData.Label] = srcData.Address
		}
	}

	return labels
}
