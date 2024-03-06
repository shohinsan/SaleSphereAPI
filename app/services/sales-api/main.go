package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/shohinsan/services/foundation/logger"
	"github.com/shohinsan/services/foundation/web"
)

/* Polymorphism means that a piece of code
changes its behavior solely depending upon
the concrete data it is operating on.
*/

// Interfaces in go mean a method set of behavior

// A polimorphic function , interface type ^

// as long as it exhibits the same behavior

// Generics in go allow us to write polymorphic functions

// Difference, it's no longer runtime time, it's compile
// we now know what concrete type we are working with

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT *******")
			return
		},
	}

	traceIDFunc := func(ctx context.Context) string {
		return web.GetTraceID(ctx)
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES-API", traceIDFunc, events)

	// -------------------------------------------------

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "start", "msg", err)
		return
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "main", "startup", "started")

	// -------------------------------------------------
	// GOMAXPROCS

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// -------------------------------------------------

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	sig := <-shutdown

	log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)
	defer log.Info(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

	return nil
}
