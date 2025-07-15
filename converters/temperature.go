package converters

import "fmt"

func ConvertTemperature(value float64, from, to string) (float64, error) {
	if from == to {
		return value, nil
	}

	var celsius float64
	switch from {
	case "C":
		celsius = value
	case "F":
		celsius = (value - 32) * 5 / 9
	case "K":
		celsius = value - 273.15
	default:
		return 0, fmt.Errorf("invalid temperature unit: %s", from)
	}

	switch to {
	case "C":
		return celsius, nil
	case "F":
		return celsius*9/5 + 32, nil
	case "K":
		return celsius + 273.15, nil
	default:
		return 0, fmt.Errorf("invalid remperature unit: %s", to)
	}
}
