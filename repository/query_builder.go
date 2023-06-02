package repository

import (
	sq "github.com/Masterminds/squirrel"
)

func qb() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
