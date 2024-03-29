package tests

import (
	"fmt"
	"net/http"

	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/handlers/homegrp"
	"github.com/shohinsan/SaleSphereAPI/business/core/crud/user"
	"github.com/shohinsan/SaleSphereAPI/business/web/page"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func homeQuery200(sd seedData) []tableData {
	total := len(sd.admins[0].homes) + len(sd.users[0].homes)
	usrsMap := make(map[uuid.UUID]user.User)

	for _, adm := range sd.admins {
		usrsMap[adm.ID] = adm.User
	}
	for _, usr := range sd.users {
		usrsMap[usr.ID] = usr.User
	}

	table := []tableData{
		{
			name:       "basic",
			url:        "/v1/homes?page=1&rows=10&orderBy=user_id,DESC",
			token:      sd.admins[0].token,
			statusCode: http.StatusOK,
			method:     http.MethodGet,
			resp:       &page.Document[homegrp.AppHome]{},
			expResp: &page.Document[homegrp.AppHome]{
				Page:        1,
				RowsPerPage: 10,
				Total:       total,
				Items:       toAppHomes(append(sd.admins[0].homes, sd.users[0].homes...)),
			},
			cmpFunc: func(x interface{}, y interface{}) string {
				resp := x.(*page.Document[homegrp.AppHome])
				exp := y.(*page.Document[homegrp.AppHome])

				var found int
				for _, r := range resp.Items {
					for _, e := range exp.Items {
						if e.ID == r.ID {
							found++
							break
						}
					}
				}

				if found != total {
					return "number of expected homes didn't match"
				}

				return ""
			},
		},
	}

	return table
}

func homeQueryByID200(sd seedData) []tableData {
	table := []tableData{
		{
			name:       "basic",
			url:        fmt.Sprintf("/v1/homes/%s", sd.users[0].homes[0].ID),
			token:      sd.users[0].token,
			statusCode: http.StatusOK,
			method:     http.MethodGet,
			resp:       &homegrp.AppHome{},
			expResp:    toAppHomePtr(sd.users[0].homes[0]),
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
	}

	return table
}
