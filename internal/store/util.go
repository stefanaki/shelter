package store

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func PostgresUUIDColumnToString(pgUUID pgtype.UUID) string {
	return uuid.UUID(pgUUID.Bytes).String()
}

func StringToPostgresUUIDColumn(uuidStr string) (pgtype.UUID, error) {
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return pgtype.UUID{}, fmt.Errorf("failed to parse UUID: %v", err)
	}

	var pgUUID pgtype.UUID
	copy(pgUUID.Bytes[:], parsedUUID[:])
	pgUUID.Valid = true
	return pgUUID, nil
}
