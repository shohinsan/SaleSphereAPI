package hackgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/shohinsan/SaleSphereAPI/business/web/v1/response"
	"github.com/shohinsan/SaleSphereAPI/foundation/web"
)

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if n := rand.Intn(100) % 2; n == 0 {
		// panic("OH MY GOODNESS WE PANICKED")
		return response.NewError(errors.New("TRUST ERROR"), http.StatusBadRequest)
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
