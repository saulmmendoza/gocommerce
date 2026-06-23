package models

import (
	"strings"
	"testing"
)

func TestAddressRequest_Validate(t *testing.T) {
	tests := []struct {
		name        string
		address     AddressRequest
		wantErr     bool
		errContains []string
	}{
		{
			name: "Valid address",
			address: AddressRequest{
				Name:     "John Doe",
				Address1: "123 Main St",
				Country:  "USA",
				City:     "New York",
				Zip:      "10001",
			},
			wantErr: false,
		},
		{
			name: "Valid address using FirstName and LastName",
			address: AddressRequest{
				FirstName: "John",
				LastName:  "Doe",
				Address1:  "123 Main St",
				Country:   "USA",
				City:      "New York",
				Zip:       "10001",
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			address: AddressRequest{
				Address1: "123 Main St",
				Country:  "USA",
				City:     "New York",
				Zip:      "10001",
			},
			wantErr:     true,
			errContains: []string{"name"},
		},
		{
			name: "Missing multiple fields",
			address: AddressRequest{
				Name: "John Doe",
			},
			wantErr:     true,
			errContains: []string{"address", "country", "city", "zip"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.address.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && len(tt.errContains) > 0 {
				errStr := err.Error()
				for _, reqField := range tt.errContains {
					if !strings.Contains(errStr, reqField) {
						t.Errorf("Validate() error = %v, should contain missing field %q", errStr, reqField)
					}
				}
			}
		})
	}
}
