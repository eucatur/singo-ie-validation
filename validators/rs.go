package validators

// RS struct - Rio Grande do Sul
// Implements the Validator interface
type RS struct {
}

// IsValid func
func (v RS) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 10) {
		return false
	}

	base := rule.GetBaseValue(insc, 9)

	weights := rule.GetWeight(2, 9)
	total := rule.CalculateTotal(insc, 9, weights)
	var digit = rule.GetDigit(total, 11)

	return insc == base+digit
}
