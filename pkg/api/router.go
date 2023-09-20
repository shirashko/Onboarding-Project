package api

import (
	"github.com/shirrashko/BuildingAServer-step2/config"
	healthAPI "github.com/shirrashko/BuildingAServer-step2/pkg/api/health"
	profileAPI "github.com/shirrashko/BuildingAServer-step2/pkg/api/profile"
	healthBL "github.com/shirrashko/BuildingAServer-step2/pkg/bl/health"
	profileBL "github.com/shirrashko/BuildingAServer-step2/pkg/bl/profile"
	"github.com/shirrashko/BuildingAServer-step2/pkg/db"
	healthDB "github.com/shirrashko/BuildingAServer-step2/pkg/repository/health"
	profileDB "github.com/shirrashko/BuildingAServer-step2/pkg/repository/profile"
)

type Handlers struct {
	handlers []IRoutesHandler
}

// check that sub services handlers are implementing IRoutesHandler interface.
// not really needed since in the return of the Router function we initiate instances of these types, but good practice.
var _ IRoutesHandler = (*profileAPI.Handler)(nil)
var _ IRoutesHandler = (*healthAPI.Handler)(nil)

func Router(conf config.Config) (Handlers, error) {
	// Hierarchy chain: handler-> service -> repo -> clientDB

	dbClient, err := db.NewDBClient(conf.DBConfig)
	if err != nil {
		return Handlers{}, err
	}
	if err != nil {
		return Handlers{}, err
	}

	// profile service
	profileRepo := profileDB.NewRepository(dbClient)
	profileService := profileBL.NewService(&profileRepo)
	profileHandler := profileAPI.NewHandler(&profileService)

	// health service
	healthRepo := healthDB.NewRepository(dbClient)
	healthService := healthBL.NewService(&healthRepo)
	healthHandler := healthAPI.NewHandler(&healthService)

	return Handlers{handlers: []IRoutesHandler{profileHandler, healthHandler}}, nil
}
