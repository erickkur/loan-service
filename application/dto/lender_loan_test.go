package dto

import (
	"testing"

	"github.com/google/uuid"
)

func TestCreateLenderLoanRequest_Validate(t *testing.T) {
	type fields struct {
		LoanGUID       uuid.UUID
		LenderGUID     uuid.UUID
		InvestedAmount int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "When loan guid is invalid, Then return Error",
			fields: fields{
				LoanGUID:       uuid.Nil,
				LenderGUID:     uuid.Nil,
				InvestedAmount: 0,
			},
			wantErr: true,
		},
		{
			name: "When lender guid is invalid, Then return Error",
			fields: fields{
				LoanGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				LenderGUID:     uuid.Nil,
				InvestedAmount: 0,
			},
			wantErr: true,
		},
		{
			name: "When invested amount is 0, Then return Error",
			fields: fields{
				LoanGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				LenderGUID:     uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e11"),
				InvestedAmount: 0,
			},
			wantErr: true,
		},
		{
			name: "When all information is valid, Then return Success",
			fields: fields{
				LoanGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				LenderGUID:     uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e11"),
				InvestedAmount: 200000,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CreateLenderLoanRequest{
				LoanGUID:       tt.fields.LoanGUID,
				LenderGUID:     tt.fields.LenderGUID,
				InvestedAmount: tt.fields.InvestedAmount,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CreateLenderLoanRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
