package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreditCard_Valid(t *testing.T) {
	tests := []struct {
		name    string
		Number  string
		Date    string
		Name    string
		Code    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "valid card",
			Number:  "5062821234567892",
			Date:    "12/99",
			Name:    "IVAN IVANOV",
			Code:    "999",
			wantErr: assert.NoError,
		},
		{
			name:   "wrong number sum",
			Number: "1234567890123456",
			Date:   "01/99",
			Name:   "VASYA A",
			Code:   "231",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "card number")
			},
		},
		{
			name:   "wrong number record",
			Number: "1234-5678-9012-3456",
			Date:   "01/99",
			Name:   "VASYA A",
			Code:   "231",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "card number")
			},
		},
		{
			name:   "wrong date",
			Number: "5062821234567892",
			Date:   "3/99",
			Name:   "VASYA PUPKIN",
			Code:   "231",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "card date")
			},
		},
		{
			name:   "wrong name",
			Number: "5062821234567892",
			Date:   "13/99",
			Name:   "VASYA999 $%^",
			Code:   "231",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "card client")
			},
		},
		{
			name:   "wrong code",
			Number: "5062821234567892",
			Date:   "03/99",
			Name:   "VASYA PUPKIN",
			Code:   "2314",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "card code")
			},
		},
		{
			name:   "wrong all",
			Number: "1234567890123456",
			Date:   "13/99",
			Name:   "VASYA999 $%^",
			Code:   "1",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "card number") &&
					assert.ErrorContains(t, err, "card date") &&
					assert.ErrorContains(t, err, "card client") &&
					assert.ErrorContains(t, err, "card code")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreditCard{
				Number: tt.Number,
				Date:   tt.Date,
				Name:   tt.Name,
				Code:   tt.Code,
			}
			tt.wantErr(t, c.Valid())
		})
	}
}
