// Package taxes provides tax calculations based on the country and postal code.
//
// This code is intended for demonstrating how module versions and imports work. Please don't
// use it in production code.
package simpletax

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

var usTaxTable = map[string]decimal.Decimal{
	"12345": decimal.NewFromFloat(.08),
	"19805": decimal.NewFromFloat(0),
	"95472": decimal.NewFromFloat(0.09),
	"23456": decimal.NewFromFloat(0.06),
	"45678": decimal.NewFromFloat(0.0725),
}

// ForZip takes in a US 5-digit zip code as a string and returns the sales tax within the zip code
// as a decimal.Decimal. If the zip is not valid, an error is returned.
func ForZip(zip string) (decimal.Decimal, error) {
	tax, ok := usTaxTable[zip]
	if !ok {
		return decimal.Decimal{}, fmt.Errorf("unknown zip: %s", zip)
	}
	return tax, nil
}

func forCAPostalCode(postalCode string) (decimal.Decimal, error) {
	if len(postalCode) == 0 {
		return decimal.Decimal{}, errors.New("no postal code supplied")
	}
	switch strings.ToUpper(postalCode)[0] {
	// Newfoundland and Labrador
	case 'A':
		return decimal.NewFromFloat(0.15), nil
	// Nova Scotia
	case 'B':
		return decimal.NewFromFloat(0.15), nil
	// Prince Edward Island (PEI)
	case 'C':
		return decimal.NewFromFloat(0.15), nil
	// New-Brunswick
	case 'E':
		return decimal.NewFromFloat(0.15), nil
	// Quebec
	case 'G', 'H', 'J':
		return decimal.NewFromFloat(0.14975), nil
	// Ontario
	case 'K', 'L', 'M', 'N', 'P':
		return decimal.NewFromFloat(0.13), nil
	// Manitoba
	case 'R':
		return decimal.NewFromFloat(0.12), nil
	// Saskatchewan
	case 'S':
		return decimal.NewFromFloat(0.11), nil
	// Alberta
	case 'T':
		return decimal.NewFromFloat(0.05), nil
	// British Columbia
	case 'V':
		return decimal.NewFromFloat(0.12), nil
	// Nunavut, Northwest Territories
	case 'X':
		return decimal.NewFromFloat(0.05), nil
	// Yukon
	case 'Y':
		return decimal.NewFromFloat(0.05), nil
	default:
		return decimal.Decimal{}, fmt.Errorf("unknown postal code: %s", postalCode)

	}

}

var countryTaxes = map[string]func(string) (decimal.Decimal, error){
	"us": ForZip,
	"ca": forCAPostalCode,
}

// ForCountryPostalCode takes in two parameters, an ISO 3166-1 alpha-2 country code and a valid postal code for that
// country. It returns a decimal.Decimal with the sales tax rate. If the country code is invalid or the postal code
// cannot be processed, an error is returned.
func ForCountryPostalCode(countryCode string, postalCode string) (decimal.Decimal, error) {
	taxTable, ok := countryTaxes[strings.ToLower(countryCode)]
	if !ok {
		return decimal.Decimal{}, fmt.Errorf("unknown country code: %s", countryCode)
	}
	return taxTable(postalCode)
}
