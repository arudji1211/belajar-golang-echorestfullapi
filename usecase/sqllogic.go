package usecase

import "database/sql"

type SqlLogic struct{}

func (s SqlLogic) CreateNullString(data string) sql.NullString {
	if len(data) == 0 {
		return sql.NullString{}
	} else {
		return sql.NullString{
			String: data,
			Valid:  true,
		}
	}
}
