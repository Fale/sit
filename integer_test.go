package sit

import "testing"

type ExpectedResults struct {
	Min   float64
	Max   float64
	Minor int
	Major int
}

func TestAbs(t *testing.T) {
	ers := []ExpectedResults{
		{Min: 10, Max: 25, Minor: 1, Major: 5},
		{Min: 20, Max: 60, Minor: 2, Major: 10},
		{Min: 20, Max: 80, Minor: 5, Major: 10},
		{Min: 1, Max: 100, Minor: 5, Major: 10},
		{Min: 1, Max: 101, Minor: 5, Major: 10},
		{Min: 1, Max: 102, Minor: 10, Major: 50},
		{Min: 1, Max: 1000, Minor: 50, Major: 100},
		{Min: 1, Max: 10000, Minor: 500, Major: 1000},
		{Min: 1, Max: 100000, Minor: 5000, Major: 10000},
	}
	for _, er := range ers {
		minor, major := tickers(er.Min, er.Max)
		if minor != er.Minor {
			t.Errorf("tickers(%f,%f) minor = %d; want %d", er.Min, er.Max, minor, er.Minor)
		}
		if major != er.Major {
			t.Errorf("tickers(%f,%f) major = %d; want %d", er.Min, er.Max, major, er.Major)
		}
	}
}
