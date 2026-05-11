package model

import "time"

type DataSource string

const (
	DataSourcePayoneer DataSource = "payoneer"
	DataSourcePayPal   DataSource = "paypal"
	DataSourceManual   DataSource = "manual"
)

type SyncStatus string

const (
	SyncStatusRunning   SyncStatus = "running"
	SyncStatusCompleted SyncStatus = "completed"
	SyncStatusFailed    SyncStatus = "failed"
	SyncStatusPartial   SyncStatus = "partial"
)

type VerificationStatus string

const (
	VerificationStatusVerified   VerificationStatus = "verified"
	VerificationStatusUnverified VerificationStatus = "unverified"
	VerificationStatusPending    VerificationStatus = "pending"
)

type FinancialInstitution struct {
	ID                 int64              `json:"id" db:"id"`
	SwiftCode          string             `json:"swift_code" db:"swift_code"`
	BankCode           string             `json:"bank_code" db:"bank_code"`
	CountryCode        string             `json:"country_code" db:"country_code"`
	LocationCode       string             `json:"location_code" db: "location_code"`
	BranchCode         string             `json:"branch_code,omitempty" db:"branch_code"`
	InstitutionName    string             `json:"institution_name" db:"institution_name"`
	ShortName          string             `json:"short_name, omitempty" db:"short_name"`
	Address            string             `json:"address,omitempty" db:"address"`
	City               string             `json:"city,omitempty" db:"city"`
	StateProvince      string             `json:"state_province,omitempty" db:"state_province"`
	PostalCode         string             `json:"postal_code,omitempty" db:"postal_code"`
	CountryName        string             `json:"country_name,omitempty" db:"country_name"`
	TimeZone           string             `json:"time_zone,omitempty" db:"time_zone"`
	PhoneNumber        string             `json:"phone_number,omitempty" db:"phone_number"`
	IsActive           bool               `json:"is_active" db:"is_active"`
	DataSource         DataSource         `json:"data_source" db:"data_source"`
	VerificationStatus VerificationStatus `json:"verification_status" db:"verification_status"`
	CreatedAt          time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" db:"updated_at"`
}

func (fi *FinancialInstitution) IsPrimaryOffice() bool {
	return fi.BranchCode == "" || fi.BranchCode == "XXX"
}

func (fi *FinancialInstitution) Is11Char() bool {
	return len(fi.SwiftCode) == 11
}
