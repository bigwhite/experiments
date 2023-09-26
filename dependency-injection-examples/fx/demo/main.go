package main

import (
	"demo/business"
	"demo/database"
	"demo/service"
	"fmt"

	"go.uber.org/fx"
)

// Client struct
type Client struct {
	service service.Service
}

// Constructor
func NewClient(service service.Service) *Client {
	return &Client{service: service}
}

// Call service
func (c Client) MakeRequest() string {
	return "Client request: " + c.service.HandleRequest()
}

func main() {
	app := fx.New(
		fx.Provide(
			fx.Annotate(
				service.NewService,
				fx.As(new(service.Service)),
			),
		),
		fx.Provide(
			fx.Annotate(
				business.NewBusiness,
				fx.As(new(business.BusinessLogic)),
			),
		),
		fx.Provide(
			fx.Annotate(
				database.NewDatabase,
				fx.As(new(database.DatabaseAccess)),
			),
		),

		fx.Invoke(func(svc service.Service) {
			client := NewClient(svc)
			fmt.Println(client.MakeRequest())
		}),
		fx.NopLogger, // no fx log output
	)

	app.Run()
}
