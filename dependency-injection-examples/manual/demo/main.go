package main

import (
	"demo/business"
	"demo/database"
	"demo/service"
	"fmt"
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
	// make dependency injection manually
	db := database.NewDatabase()
	busi := business.NewBusiness(db)
	svc := service.NewService(busi)
	client := NewClient(svc)

	fmt.Println(client.MakeRequest())
}
