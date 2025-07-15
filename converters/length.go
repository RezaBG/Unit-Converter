package converters

import (
	"fmt"
)

func ConvertLength(value float64, from, to string) (float64, error) {
	conversions := map[string]float64{
		"mm":   0.001,
		"cm":   0.01,
		"m":    1.0,
		"km":   1000.0,
		"inch": 0.0254,
		"ft":   0.3048,
		"yard": 0.9144,
		"mile": 1609.34,
	}

	fromFactor, ok1 := conversions[from]
	toFactor, ok2 := conversions[to]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("invalid unit conversion: %s or %s", from, to)
	}

	meters := value * fromFactor
	converted := meters / toFactor

	return converted, nil
}
