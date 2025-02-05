package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "arym",
		Price: 100.00,
		SKU:   "adsf-dvd-fdfd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
