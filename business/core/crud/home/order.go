package home

import "github.com/shohinsan/SaleSphereAPI/business/web/v1/order"

// DefaultOrderBy represents the default way we sort.
var DefaultOrderBy = order.NewBy(OrderByID, order.ASC)

// Set of fields that the results can be ordered by.
const (
	OrderByID     = "home_id"
	OrderByType   = "type"
	OrderByUserID = "user_id"
)
