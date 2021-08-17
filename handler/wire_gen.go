// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package handler

import (
	"github.com/grpc-shop/order-srv/dao"
	"github.com/grpc-shop/order-srv/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitOrderHandler(db *gorm.DB) *OrderHandler {
	orderDao := dao.NewOrderImpl(db)
	orderServer := service.NewOrderServerImpl(orderDao)
	orderHandler := NewOrderHandler(orderServer)
	return orderHandler
}
