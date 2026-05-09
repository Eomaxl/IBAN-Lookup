package repository

import (
	"context"

	"github.com/eomaxl/iban-lookup/internal/model"
)

type BICReader interface {
	Search(ctx context.Context, req model.SearchRequest) ([]model.FinancialInstitution, int64, error)
	FindByIbanCode(ctx context.Context, ibanCode string) (*model.FinancialInstitution, error)
}

type BICWriter interface {
	BulkUpsert(ctx context.Context, institutions []model.FinancialInstitution) error
}

type BICRepository interface {
	BICReader
	BICWriter
	LoadTopPrefixes(ctx context.Context, limit int) ([]string, error)
	LoadByPrefixes(ctx context.Context, prefix string) ([]model.FinancialInstitution, error)
}

type HistoryRepository interface {
	LogChange(ctx context.Context, entry model.HistoryEntry) error
	BulkLogChanges(ctx context.Context, entries []model.HistoryEntry) error
}

type SyncRepository interface {
	CreateBatch(ctx context.Context, source model.DataSource) (*model.SyncBatch, error)
	UpdateBatch(ctx context.Context, batch *model.SyncBatch) error
	GetLastSync(ctx context.Context, source model.DataSource) (*model.SyncBatch, error)
	LoadString(ctx context.Context, records []model.StagingRecord, batchID string) error
	ClearStaging(ctx context.Context, batchID string) error
}
