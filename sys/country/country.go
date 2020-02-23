package country

import "github.com/pariz/gountries"

func GetSymbol() {
	gountries.New().FindAllCountries()
}
