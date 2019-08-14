package tempconv

import (
	"math"
	"testing"
)

func TestTempConv(t *testing.T) {
	tests := []struct {
		f Fahrenheit
		c Celsius
		k Kelvin
	}{
		{68, 20, 293.15},
		{32, 0, 273.15},
		{-40, -40, 233.15},
	}
	eps := 0.0000001 // acceptable floating point error
	for _, test := range tests {
		if math.Abs(float64(CToF(test.c)-test.f)) > eps {
			t.Errorf("CToF(%s): got %s, want %s", test.c, CToF(test.c), test.f)
		}
		if math.Abs(float64(FToC(test.f)-test.c)) > eps {
			t.Errorf("FToC(%s): got %s, want %s", test.f, FToC(test.f), test.c)
		}
		if math.Abs(float64(KToC(test.k)-test.c)) > eps {
			t.Errorf("KToC(%s): got %s, want %s", test.k, KToC(test.k), test.c)
		}
		if math.Abs(float64(CToK(test.c)-test.k)) > eps {
			t.Errorf("CToK(%s): got %s, want %s", test.c, CToK(test.c), test.k)
		}
	}
}

func TestCelsius_String(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Celsius.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFahrenheit_String(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("Fahrenheit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKelvin_String(t *testing.T) {
	tests := []struct {
		name string
		k    Kelvin
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Kelvin.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
