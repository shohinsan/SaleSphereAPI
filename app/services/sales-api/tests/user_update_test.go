package tests

import (
	"fmt"
	"net/http"

	"github.com/shohinsan/SaleSphereAPI/app/services/sales-api/handlers/usergrp"
	"github.com/shohinsan/SaleSphereAPI/business/data/dbtest"
	"github.com/shohinsan/SaleSphereAPI/business/web/errs"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func userUpdate200(sd seedData) []tableData {
	table := []tableData{
		{
			name:       "basic",
			url:        fmt.Sprintf("/v1/users/%s", sd.users[0].ID),
			token:      sd.users[0].token,
			method:     http.MethodPut,
			statusCode: http.StatusOK,
			model: &usergrp.AppUpdateUser{
				Name:            dbtest.StringPointer("Jack Kennedy"),
				Email:           dbtest.StringPointer("jack@ardanlabs.com"),
				Roles:           []string{"ADMIN"},
				Department:      dbtest.StringPointer("IT"),
				Password:        dbtest.StringPointer("123"),
				PasswordConfirm: dbtest.StringPointer("123"),
			},
			resp: &usergrp.AppUser{},
			expResp: &usergrp.AppUser{
				Name:       "Jack Kennedy",
				Email:      "jack@ardanlabs.com",
				Roles:      []string{"ADMIN"},
				Department: "IT",
				Enabled:    true,
			},
			cmpFunc: func(x interface{}, y interface{}) string {
				resp := x.(*usergrp.AppUser)
				expResp := y.(*usergrp.AppUser)

				if _, err := uuid.Parse(resp.ID); err != nil {
					return "bad uuid for ID"
				}

				if resp.DateCreated == "" {
					return "missing date created"
				}

				if resp.DateUpdated == "" {
					return "missing date updated"
				}

				expResp.ID = resp.ID
				expResp.DateCreated = resp.DateCreated
				expResp.DateUpdated = resp.DateUpdated

				return cmp.Diff(x, y)
			},
		},
	}

	return table
}

func userUpdate400(sd seedData) []tableData {
	table := []tableData{
		{
			name:       "bad-input",
			url:        fmt.Sprintf("/v1/users/%s", sd.users[0].ID),
			token:      sd.users[0].token,
			method:     http.MethodPut,
			statusCode: http.StatusBadRequest,
			model: &usergrp.AppUpdateUser{
				Email:           dbtest.StringPointer("bill@"),
				PasswordConfirm: dbtest.StringPointer("jack"),
			},
			resp: &errs.Response{},
			expResp: &errs.Response{
				Error:  "data validation error",
				Fields: map[string]string{"email": "email must be a valid email address", "passwordConfirm": "passwordConfirm must be equal to Password"},
			},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
		{
			name:       "bad-role",
			url:        fmt.Sprintf("/v1/users/%s", sd.users[0].ID),
			token:      sd.users[0].token,
			method:     http.MethodPut,
			statusCode: http.StatusBadRequest,
			model: &usergrp.AppUpdateUser{
				Roles: []string{"BAD ROLE"},
			},
			resp: &errs.Response{},
			expResp: &errs.Response{
				Error: "parse: invalid role \"BAD ROLE\"",
			},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
	}

	return table
}

func userUpdate401(sd seedData) []tableData {
	table := []tableData{
		{
			name:       "emptytoken",
			url:        fmt.Sprintf("/v1/users/%s", sd.users[0].ID),
			token:      "",
			method:     http.MethodPut,
			statusCode: http.StatusUnauthorized,
			resp:       &errs.Response{},
			expResp:    &errs.Response{Error: "Unauthorized"},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
		{
			name:       "badsig",
			url:        fmt.Sprintf("/v1/users/%s", sd.users[0].ID),
			token:      sd.users[0].token + "A",
			method:     http.MethodPut,
			statusCode: http.StatusUnauthorized,
			resp:       &errs.Response{},
			expResp:    &errs.Response{Error: "Unauthorized"},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
		{
			name:       "wronguser",
			url:        fmt.Sprintf("/v1/users/%s", sd.admins[0].ID),
			token:      sd.users[0].token,
			method:     http.MethodPut,
			statusCode: http.StatusUnauthorized,
			model: &usergrp.AppUpdateUser{
				Name:            dbtest.StringPointer("Bill Kennedy"),
				Email:           dbtest.StringPointer("bill@ardanlabs.com"),
				Roles:           []string{"ADMIN"},
				Department:      dbtest.StringPointer("IT"),
				Password:        dbtest.StringPointer("123"),
				PasswordConfirm: dbtest.StringPointer("123"),
			},
			resp:    &errs.Response{},
			expResp: &errs.Response{Error: "Unauthorized"},
			cmpFunc: func(x interface{}, y interface{}) string {
				return cmp.Diff(x, y)
			},
		},
	}

	return table
}
