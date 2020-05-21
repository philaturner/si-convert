package unit_test

import (
	unit "si-convert/convert"
	"testing"
)

func TestUnsupportedConversion(t *testing.T) {
	siUnit := unit.SI{Value: 100, Option: "b"}
	_, err := siUnit.Convert("a")
	if err == nil {
		t.Error()
	}
}

var meterTests = []struct {
	input    float64 // input
	to       string  // conversion to
	expected float64 // expected result
}{
	{100, "cm", 10000},
	{1.5, "cm", 150},
	{100, "km", 0.1},
	{1.5, "km", 0.0015},
	{1, "yd", 1.094},
	{1.5, "yd", 1.641},
	{1, "ft", 3.281},
	{10, "ft", 32.81},
}

func TestMeterConversions(t *testing.T) {
	for _, tm := range meterTests {
		siUnit := unit.SI{Value: tm.input, Option: "m"}
		actual, err := siUnit.Convert(tm.to)
		if err != nil {
			t.Errorf("Convert(%g): to %s,  expected %g, actual %g", tm.input, tm.to, tm.expected, actual)
		}
		if actual != tm.expected {
			t.Errorf("Convert(%g): to %s, expected %g, actual %g", tm.input, tm.to, tm.expected, actual)
		}
	}
}

var centimeterTests = []struct {
	input    float64 // input
	to       string  // conversion to
	expected float64 // expected result
}{
	{2.54, "in", 1},
	{22.5, "in", 8.858267716535433},
	{100, "m", 1},
	{1, "m", 0.01},
	{100, "km", 0.001},
	{100.5, "km", 0.001005},
	{1.5, "yd", 0.016404199475065617},
	{1500, "yd", 16.404199475065617},
	{20, "ft", 0.6561679790026247},
	{2003, "ft", 65.71522309711285},
}

func TestCentimeterConversions(t *testing.T) {
	for _, tcm := range centimeterTests {
		siUnit := unit.SI{Value: tcm.input, Option: "cm"}
		actual, err := siUnit.Convert(tcm.to)
		if err != nil {
			t.Errorf("Convert(%g): to %s,  expected %g, actual %g", tcm.input, tcm.to, tcm.expected, actual)
		}
		if actual != tcm.expected {
			t.Errorf("Convert(%g): to %s, expected %g, actual %g", tcm.input, tcm.to, tcm.expected, actual)
		}
	}
}

var inchTests = []struct {
	input    float64 // input
	to       string  // conversion to
	expected float64 // expected result
}{
	{1, "cm", 2.54},
	{10.4, "cm", 26.416},
	{1, "m", 0.025400050800101603},
	{13.4, "m", 0.3403606807213615},
	{100, "km", 0.00254000508001016},
	{100.5, "km", 0.0025527051054102107},
	{1.5, "yd", 0.041666666666666664},
	{1500, "yd", 41.666666666666664},
	{20, "ft", 1.6666666666666667},
	{3005.2, "ft", 250.4333333333333},
}

func TestInchConversions(t *testing.T) {
	for _, tinc := range inchTests {
		siUnit := unit.SI{Value: tinc.input, Option: "in"}
		actual, err := siUnit.Convert(tinc.to)
		if err != nil {
			t.Errorf("Convert(%g): to %s,  expected %g, actual %g", tinc.input, tinc.to, tinc.expected, actual)
		}
		if actual != tinc.expected {
			t.Errorf("Convert(%g): to %s, expected %g, actual %g", tinc.input, tinc.to, tinc.expected, actual)
		}
	}
}

var footTests = []struct {
	input    float64 // input
	to       string  // conversion to
	expected float64 // expected result
}{
	{1, "in", 12},
	{56.5, "in", 678},
	{1, "cm", 30.48},
	{10.4, "cm", 316.992},
	{1, "m", 0.30478512648582745},
	{13.4, "m", 4.084120694910088},
	{100, "km", 0.03047851264858275},
	{100.5, "km", 0.030630905211825665},
	{202, "yd", 67.33333333333333},
	{1234.5, "yd", 411.5},
}

func TestFootConversions(t *testing.T) {
	for _, tft := range footTests {
		siUnit := unit.SI{Value: tft.input, Option: "ft"}
		actual, err := siUnit.Convert(tft.to)
		if err != nil {
			t.Errorf("Convert(%g): to %s,  expected %g, actual %g", tft.input, tft.to, tft.expected, actual)
		}
		if actual != tft.expected {
			t.Errorf("Convert(%g): to %s, expected %g, actual %g", tft.input, tft.to, tft.expected, actual)
		}
	}
}

var kilometerTests = []struct {
	input    float64 // input
	to       string  // conversion to
	expected float64 // expected result
}{
	{1, "cm", 100000},
	{122.4, "cm", 12240000},
	{10, "m", 10000},
	{122.4, "m", 122400},
	{1, "ft", 3281},
	{55.1, "ft", 180783.1},
	{1, "yd", 1094},
	{12.4, "yd", 13565.6},
}

func TestKilometerConversion(t *testing.T) {
	for _, kt := range kilometerTests {
		siUnit := unit.SI{Value: kt.input, Option: "km"}
		actual, err := siUnit.Convert(kt.to)
		if err != nil {
			t.Errorf("Convert(%g): to %s,  expected %g, actual %g", kt.input, kt.to, kt.expected, actual)
		}
		if actual != kt.expected {
			t.Errorf("Convert(%g): to %s, expected %g, actual %g", kt.input, kt.to, kt.expected, actual)
		}
	}
}
