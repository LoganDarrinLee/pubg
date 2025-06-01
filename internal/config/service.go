package config

import (
  "log"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

  "github.com/LoganDarrinLee/pubg/internal/common"
)


// easier http service management
type HttpService interface {

	// mount the http service on the main router. 
	Mount(
    env *EnvConfig, 
    logger *common.BasicLogger, 
    pool *pgxpool.Pool,
    r *chi.Mux,
  ) error 
}

// Iterate through the provided services, and mount them on the main router.
func ConfigureServices(
  env *EnvConfig, 
  logger *common.BasicLogger,
  pool *pgxpool.Pool, 
  router *chi.Mux, 
  xs ...HttpService,
) {
  for i := len(xs) - 1; i >= 0; i-- {
    if err := xs[i].Mount(env, logger, pool, router); err != nil {
      log.Fatal(err)
    }
  }
}

