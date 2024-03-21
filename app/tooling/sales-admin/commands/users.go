package commands

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-json-experiment/json"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/user"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/user/stores/userdb"
	"github.com/shohinsan/SaleSphereAPI/business/data/sqldb"
	"github.com/shohinsan/SaleSphereAPI/foundation/logger"
)

// Users retrieves all users from the database.
func Users(log *logger.Logger, cfg sqldb.Config, pageNumber string, rowsPerPage string) error {
	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	page, err := strconv.Atoi(pageNumber)
	if err != nil {
		return fmt.Errorf("converting page number: %w", err)
	}

	rows, err := strconv.Atoi(rowsPerPage)
	if err != nil {
		return fmt.Errorf("converting rows per page: %w", err)
	}

	core := user.NewCore(log, nil, userdb.NewStore(log, db))

	users, err := core.Query(ctx, user.QueryFilter{}, user.DefaultOrderBy, page, rows)
	if err != nil {
		return fmt.Errorf("retrieve users: %w", err)
	}

	return json.MarshalWrite(os.Stdout, users, json.FormatNilSliceAsNull(true))
}