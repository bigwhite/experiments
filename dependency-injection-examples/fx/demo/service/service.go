package service

import "demo/business"

// Service interface
type Service interface {
	HandleRequest() string
}

// ServiceImpl struct
type ServiceImpl struct {
	logic business.BusinessLogic
}

// Constructor
func NewService(logic business.BusinessLogic) *ServiceImpl {
	return &ServiceImpl{logic: logic}
}

// Implement HandleRequest()
func (s ServiceImpl) HandleRequest() string {
	return "Handled request: " + s.logic.ProcessData()
}
