package converters

import "fmt"

func ConvertWeight(value float64, from, to string) (float64, error) {
	conversions := map[string]float64{
		"mg":    0.001,
		"g":     1.0,
		"kg":    1000.0,
		"ounce": 28.3495,
		"pound": 453.592,
	}

	fromFactor, ok1 := conversions[from]
	toFactor, ok2 := conversions[to]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("invalid unit conversion: %s or %s", from, to)
	}

	grams := value * fromFactor
	converted := grams / toFactor

	return converted, nil
}
