//go:build wireinject
// +build wireinject

package main

// wire.go

import (
	"demo/business"
	"demo/database"
	"demo/service"

	"github.com/google/wire"
)

func InitializeService() service.Service {
	wire.Build(service.NewService,
		wire.Bind(new(service.Service), new(*service.ServiceImpl)),
		business.NewBusiness,
		wire.Bind(new(business.BusinessLogic), new(*business.Business)),
		database.NewDatabase,
		wire.Bind(new(database.DatabaseAccess), new(*database.Database)),
	)
	return nil
}
