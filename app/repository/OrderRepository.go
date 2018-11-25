package repository

import (
	. "../helper"
	. "../model"
	. "../model/response"
	"github.com/kataras/golog"
)

type IOrderRepository interface {
	FindAllOrdersByBorrowerId(id int) (orders []OrderResponse, errorCode uint)
	DoesOrderIdExist(id int) bool
	FindOrderById(id int) Order
	UpdateOrderStatus(order Order, orderStatus string)
	SaveOrder(order Order) bool
}

type OrderRepository struct {
	IDatabaseConnectionHelper
}

func (orderRepository *OrderRepository) FindAllOrdersByBorrowerId(id int) (orders []OrderResponse, errorCode uint) {
	var orderResponses []OrderResponse
	db := orderRepository.OpenDatabaseConnection()
	defer db.Close()
	rows, err := db.Raw("SELECT orders.id, products.name AS product_name, products.image_name AS image_url, orders.total_price, orders.order_status, orders.return_time "+
		"FROM orders, products "+
		"WHERE orders.borrower_id = ? AND orders.product_id = products.id", id).Rows()
	defer rows.Close()
	if err != nil {
		golog.Warn("Error raw SQL selecting all orders " + err.Error())
		return orderResponses, GET_ALL_ORDERS_FAILED_SQL_ERROR
	} else {
		for rows.Next() {
			var orderResponse OrderResponse
			db.ScanRows(rows, &orderResponse)
			orderResponse.ImageUrl = IMAGE_URL_PREFIX + orderResponse.ImageUrl
			orderResponses = append(orderResponses, orderResponse)
		}
		golog.Info("Get all orders succeed")
	}
	return orderResponses, 0
}

func (orderRepository *OrderRepository) DoesOrderIdExist(id int) bool {
	var order Order
	db := orderRepository.OpenDatabaseConnection()
	defer db.Close()
	return !db.Where("id = ?", id).Find(&order).RecordNotFound()
}

func (orderRepository *OrderRepository) FindOrderById(id int) Order {
	var order Order
	db := orderRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Where("id = ?", id).Find(&order)
	return order
}

func (orderRepository *OrderRepository) UpdateOrderStatus(order Order, orderStatus string) {
	db := orderRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Model(&order).Update("order_status", orderStatus)
}

func (orderRepository *OrderRepository) SaveOrder(order Order) bool {
	db := orderRepository.OpenDatabaseConnection()
	defer db.Close()
	db.Create(&order)
	return !db.NewRecord(order) // if the `order` object is not a new record, it means that it's successfully saved to database
}