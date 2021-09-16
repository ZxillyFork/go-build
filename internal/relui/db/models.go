// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Migration struct {
	Version int64
	Dirty   bool
}

type Task struct {
	WorkflowID uuid.UUID
	Name       string
	Finished   bool
	Result     sql.NullString
	Error      sql.NullString
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Workflow struct {
	ID        uuid.UUID
	Params    sql.NullString
	Name      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}
