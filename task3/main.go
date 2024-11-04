package main

import (
	"sync"
	"fmt"
)


/* Этап 1: Верификация заказа: проверка правильности данных. +
 Этап 2: Оплата: обработка платежа. + 
 Этап 3: Отгрузка: добавление заказа в очередь на доставку.
 Этап 4: Каждый этап требует разного времени обработки, и система может обрабатывать несколько заказов одновременно. Необходимо разработать систему на Go, которая будет:

Обрабатывать каждый заказ в отдельной горутине.
Гарантировать, что заказ переходит к следующему этапу только после завершения предыдущего.
Использовать синхронизацию, чтобы корректно завершить все горутины перед завершением программы. */

type Order struct {
	ID int
	Name string
	Quantity int
	PaymentStatus bool
}

func(o *Order) Verify() bool {
	var trueOrFalse bool
	if o.ID >= 0 && o.Name != "" && o.Quantity > 0 {
		trueOrFalse = true
	}

	return trueOrFalse
}

func (o *Order) Pay() bool {
    // Обработка оплаты заказа
    o.PaymentStatus = true
	return o.PaymentStatus
}

func (o *Order) processOrder(order Order, orderCh chan Order, doneCh chan <- struct{}, wg sync.WaitGroup) {
	defer wg.Done()
	
	if !order.Verify() {
		fmt.Printf("Order %d failed verification\n", order.ID)
		return
	}

	// Этап 2: Обработка платежа
	if !order.Pay() {
		fmt.Printf("Order %d failed payment processing\n", order.ID)
		return
	}

	orderCh <- order
}

func main() {

}