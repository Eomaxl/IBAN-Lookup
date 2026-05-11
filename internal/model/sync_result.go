package model

import "time"

type SyncBatch struct {
	ID               string     `json:"id" db:"id"`
	Source           DataSource `json:"source" db:"source"`
	Status           SyncStatus `json:"status" db:"status"`
	StartedAt        time.Time  `json:"started_at" db:"started_at"`
	CompletedAt      *time.Time `json:"completed_at,omitempty" db:"completed_at"`
	FilePath         string     `json:"file_path,omitempty" db:"file_path"`
	RecordsProcessed int        `json:"records_processed" db:"records_processed"`
	RecordsAdded     int        `json:"records_added" db:"records_added"`
	RecordsUpdated   int        `json:"records_updated" db:"records_updated"`
	RecordsFailed    int        `json:"records_failed" db:"records_failed"`
	ErrorMessage     string     `json:"error_message,omitempty" db:"error_message"`
}

type SyncResult struct {
	BatchID  string        `json:"batch_id"`
	Source   DataSource    `json:"source"`
	Added    int           `json:"added"`
	Updated  int           `json:"updated"`
	Failed   int           `json:"failed"`
	Duration time.Duration `json:"duration_ms"`
	Error    error         `json:"error,omitempty"`
}

type StagingRecord struct {
	SwiftCode       string     `json:"swift_code"`
	InstitutionName string     `json:"institution_name"`
	CountryCode     string     `json:"country_code"`
	Address         string     `json:"address,omitempty"`
	City            string     `json:"city,omitempty"`
	StateProvince   string     `json:"state_province,omitempty"`
	PostalCode      string     `json:"postal_code,omitempty"`
	PhoneNumber     string     `json:"phone_number,omitempty"`
	Source          DataSource `json:"source"`
	RowNumber       int        `json:"row_number"`
}

type HistoryEntry struct {
	ID         int64     `json:"id" db:"id"`
	SwiftCode  string    `json:"swift_code" db:"swift_code"`
	ChangeType string    `json:"change_type" db:"change_type"`
	OldValues  string    `json:"old_values,omitempty" db:"old_values"`
	NewValues  string    `json:"new_values,omitempty" db:"new_values"`
	ChangedBy  string    `json:"changed_by" db:"changed_by"`
	BatchID    string    `json:"batch_id,omitempty" db:"batch_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
