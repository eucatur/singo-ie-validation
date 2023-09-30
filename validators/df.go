package validators

import (
	"fmt"
	"regexp"
)

// DistroFederal struct - Distrito Federal
// Implements the Validator interface
type DistroFederal struct {
}

// IsValid func
func (v DistroFederal) IsValid(insc string) bool {

	rule := NewRule()

	re := regexp.MustCompile("[^0-9]")
	insc = re.ReplaceAllString(insc, "")

	if !rule.IsCorrectSize(insc, 13) {

		if len([]rune(insc)) > 13 {
			return false
		}

		return v.IsValid(fmt.Sprintf("%013s", insc))
	}

	base := rule.GetBaseValue(insc, 11)

	//{4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights := rule.GetWeight(4, 11)
	total := rule.CalculateTotal(insc, 11, weights)
	firstDigit := rule.GetDigit(total, 11)

	//{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weights = rule.GetWeight(5, 12)
	total = rule.CalculateTotal(base+firstDigit, 12, weights)
	secondDigit := rule.GetDigit(total, 11)

	return insc == base+firstDigit+secondDigit
}
