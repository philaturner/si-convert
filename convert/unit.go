package unit

import "errors"

type SI struct {
	Value float64
	Option string
}

func (si SI) Convert(t string)(r float64, err error) {
	switch si.Option {
	case "cm":
		switch t {
		case "m":
			return si.Value / 100, nil
		case "km":
			return si.Value / 100000, nil
		case "yd":
			return si.Value /  91.44, nil
		case "ft":
			return si.Value / 30.48, nil
		}
	case "in":
		switch t {
		case "cm":
			return si.Value * 2.54, nil
		case "m":
			return si.Value / 39.37, nil
		case "km":
			return si.Value / 39370, nil
		case "yd":
			return si.Value /  36, nil
		case "ft":
			return si.Value / 12, nil
		}
	case "ft":
		switch t {
		case "m":
			return si.Value / 3.281, nil
		case "cm":
			return si.Value * 30.48, nil
		case "km":
			return si.Value / 3281, nil
		case "yd":
			return si.Value / 3, nil
		}
	case "m":
		switch t {
		case "cm":
			return si.Value * 100, nil
		case "km":
			return si.Value / 1000, nil
		case "yd":
			return si.Value * 1.094, nil
		case "ft":
			return si.Value * 3.281, nil
		}
	case "km":
		switch t {
		case "cm":
			return si.Value * 100000, nil
		case "m":
			return si.Value * 1000, nil
		case "yd":
			return si.Value * 1094, nil
		case "ft":
			return si.Value * 3281, nil
		}
	}

	return 0, errors.New("no support for this conversion")
}