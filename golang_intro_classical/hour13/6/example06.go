package temperature

import "testing"

type (
	temperatureTest struct {
		i        float64
		expected Temperature
	}
)

func (t temperatureTest) I() float64 {
	return t.i
}

func (t temperatureTest) Expected() Temperature {
	return t.expected
}

func newTemperatureTest(i float64) *temperatureTest {
	return &temperatureTest{i: i}
}

var CtoFTests = []temperatureTest{
	{4.1, 39.38},
	{10, 50},
	{-10, 14},
}

var FtoCTests = []temperatureTest{
	{32, 0},
	{50, 10},
	{5, -15},
}

func TestCtoF(t *testing.T) {
	for _, tt := range CtoFTests {
		actual := CtoF(tt.i)
		if actual != tt.expected {
			t.Errorf("expected %v, actual %v", tt.expected, actual)
		}
	}
}

func CtoF(i float64) interface{} {

	return nil
}

func TestFtoC(t *testing.T) {
	for _, tt := range FtoCTests {
		actual := FtoC(tt.i)
		if actual != tt.expected {
			t.Errorf("expected %v, actual %v", tt.expected, actual)
		}
	}
}

func FtoC(i float64) interface{} {

	return nil
}
