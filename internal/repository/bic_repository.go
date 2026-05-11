// Package repository defines data access interfaces and their PostgreSQL implementations.
package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/eomaxl/swift-lookup/internal/model"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgxBICRepository struct {
	pool *pgxpool.Pool
}

func NewBICRepository(pool *pgxpool.Pool) BICRepository {
	return &pgxBICRepository{pool: pool}
}

func (r *pgxBICRepository) Search(ctx context.Context, req model.SearchRequest) ([]model.FinancialInstitution, int64, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		req.PageSize = 20
	}

	offset := (req.Page - 1) * req.PageSize

	var conditions []string
	var args []interface{}
	argIdx := 1

	if req.Prefix != "" {
		conditions = append(conditions, fmt.Sprintf("swift_code LIKE $%d", argIdx))
		args = append(args, req.Prefix+"%")
		argIdx++
	}

	if req.CountryCode != "" {
		conditions = append(conditions, fmt.Sprintf("country_code = $%d", argIdx))
		args = append(args, strings.ToUpper((req.CountryCode)))
		argIdx++
	}

	if req.InstitutionName != "" {
		conditions = append(conditions, fmt.Sprintf("institution_name ILIKE $%d", argIdx))
		args = append(args, "%"+req.InstitutionName+"%")
		argIdx++
	}

	if req.ActiveOnly {
		conditions = append(conditions, "is_active = true")
	}

	whereClause := ""

	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM financial_institutions %s`, whereClause)
	var total int64
	if err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("bic repository count : %w", err)
	}

	query := fmt.Sprintf(`
		SELECT id, swift_code, bank_code, country_code, location_code, branch_code, 
		institution_name, short_name, address, city, state_province, postal_code, 
		country_name, time_zone, phone_number, is_active, data_source, verification_status, 
		created_at, updated_at
		FROM financial_institutions
		%s
		ORDER BY swift_code ASC
		LIMIT $%d OFFSET $%d`,
		whereClause, argIdx, argIdx+1,
	)

	args = append(args, req.PageSize, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("bic repository search query : %w", &err)
	}
	defer rows.Close()

	institutions, err := scanInstitutions(rows)
	if err != nil {
		return nil, 0, fmt.Errorf("bic repository scan : %w", err)
	}

	return institutions, total, nil
}

func scanInstitutions(rows pgx.Rows) ([]model.FinancialInstitution, error) {
	var results []model.FinancialInstitution
	for rows.Next() {
		var fi model.FinancialInstitution
		err := rows.Scan()

		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return results, nil
			}
			return nil, fmt.Errorf("scanning insititutions row : %w", err)
		}
		results = append(results, fi)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating institutions rows : %w", err)
	}
	return results, nil
}
