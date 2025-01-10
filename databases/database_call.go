package databases

import (
	"github.com/jmoiron/sqlx"
)

type (
	AppDbImplement struct {
		ConnectTkbaiDB *sqlx.DB
		Err            error
	}
)

var DbInterface *AppDbImplement
