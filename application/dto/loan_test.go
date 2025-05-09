package dto

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUpdateLoanRequest_Validate(t *testing.T) {
	type fields struct {
		LoanGUID           *uuid.UUID
		Status             string
		PictureProof       string
		EmployeeGUID       uuid.UUID
		DateOfApproval     time.Time
		DateOfDisbursement time.Time
		AgreementLetter    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "When request status is invalid, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "invalid",
				PictureProof:       "",
				EmployeeGUID:       uuid.Nil,
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
				AgreementLetter:    "",
			},
			wantErr: true,
		},
		{
			name: "When request status is approved and pictureProof is empty, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "approved",
				PictureProof:       "",
				EmployeeGUID:       uuid.Nil,
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
				AgreementLetter:    "",
			},
			wantErr: true,
		},
		{
			name: "When request status is approved and pictureProof is not image, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "approved",
				PictureProof:       "test",
				EmployeeGUID:       uuid.Nil,
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
				AgreementLetter:    "",
			},
			wantErr: true,
		},
		{
			name: "When request status is approved and employeeGUID is empty, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "approved",
				PictureProof:       "test.jpg",
				EmployeeGUID:       uuid.Nil,
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
				AgreementLetter:    "",
			},
			wantErr: true,
		},
		{
			name: "When request status is approved and dateOfApproval is empty, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "approved",
				PictureProof:       "test",
				EmployeeGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
				AgreementLetter:    "",
			},
			wantErr: true,
		},
		{
			name: "When request status is disbursed and agreementLetter is empty, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "disbursed",
				AgreementLetter:    "",
				PictureProof:       "",
				EmployeeGUID:       uuid.Nil,
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
			},
			wantErr: true,
		},
		{
			name: "When request status is disbursed and agreementLetter extension is not pdf, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "disbursed",
				AgreementLetter:    "test",
				PictureProof:       "",
				EmployeeGUID:       uuid.Nil,
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
			},
			wantErr: true,
		},
		{
			name: "When request status is disbursed and employee guid is empty, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "disbursed",
				AgreementLetter:    "test.pdf",
				PictureProof:       "",
				EmployeeGUID:       uuid.Nil,
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
			},
			wantErr: true,
		},
		{
			name: "When request status is disbursed and date of disbursement is empty, Then return error",
			fields: fields{
				LoanGUID:           nil,
				Status:             "disbursed",
				AgreementLetter:    "",
				PictureProof:       "",
				EmployeeGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				DateOfApproval:     time.Time{},
				DateOfDisbursement: time.Time{},
			},
			wantErr: true,
		},
		{
			name: "When request status is approved and all information is valid, Then return success",
			fields: fields{
				LoanGUID:           nil,
				Status:             "approved",
				PictureProof:       "test.jpg",
				EmployeeGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				DateOfApproval:     time.Now(),
				DateOfDisbursement: time.Time{},
				AgreementLetter:    "",
			},
			wantErr: false,
		},
		{
			name: "When request status is disbursed and all information is valid, Then return success",
			fields: fields{
				LoanGUID:           nil,
				Status:             "disbursed",
				AgreementLetter:    "test.pdf",
				EmployeeGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				DateOfDisbursement: time.Now(),
				PictureProof:       "",
				DateOfApproval:     time.Time{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := UpdateLoanRequest{
				LoanGUID:           tt.fields.LoanGUID,
				Status:             tt.fields.Status,
				PictureProof:       tt.fields.PictureProof,
				EmployeeGUID:       tt.fields.EmployeeGUID,
				DateOfApproval:     tt.fields.DateOfApproval,
				DateOfDisbursement: tt.fields.DateOfDisbursement,
				AgreementLetter:    tt.fields.AgreementLetter,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLoanRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateLoanRequest_Validate(t *testing.T) {
	type fields struct {
		BorrowerGUID       uuid.UUID
		PrincipalAmount    int64
		Rate               float64
		ReturnOfInvestment int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "When borrowerGUID is null, Then return error",
			fields: fields{
				BorrowerGUID:       uuid.Nil,
				PrincipalAmount:    0,
				Rate:               0,
				ReturnOfInvestment: 0,
			},
			wantErr: true,
		},
		{
			name: "When principalAmount is 0, Then return error",
			fields: fields{
				BorrowerGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				PrincipalAmount:    0,
				Rate:               0,
				ReturnOfInvestment: 0,
			},
			wantErr: true,
		},
		{
			name: "When Rate is 0, Then return error",
			fields: fields{
				BorrowerGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				PrincipalAmount:    20000000,
				Rate:               0,
				ReturnOfInvestment: 0,
			},
			wantErr: true,
		},
		{
			name: "When returnOfInvestment is 0, Then return error",
			fields: fields{
				BorrowerGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				PrincipalAmount:    20000000,
				Rate:               0.25,
				ReturnOfInvestment: 0,
			},
			wantErr: true,
		},
		{
			name: "When all information is valid, Then return Success",
			fields: fields{
				BorrowerGUID:       uuid.MustParse("b3a6d7e2-2f8b-4e9e-9c0e-6f3e2e1e1e1e"),
				PrincipalAmount:    20000000,
				Rate:               0.25,
				ReturnOfInvestment: 30000000,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CreateLoanRequest{
				BorrowerGUID:       tt.fields.BorrowerGUID,
				PrincipalAmount:    tt.fields.PrincipalAmount,
				Rate:               tt.fields.Rate,
				ReturnOfInvestment: tt.fields.ReturnOfInvestment,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CreateLoanRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
