package load

import (
	"fmt"

	db "github.com/kanounzied/mini-Go-project/db"
	model "github.com/kanounzied/mini-Go-project/model"
)

func LoadData() (preparedData map[int64][]model.EventData) {

	fmt.Println("\n[Loader] Loading data ...")
	// prepare all customers
	var customers []model.Customer = db.GetAllCustomers()
	// prepare all EventData with purchace type
	var purchaces []model.EventData = db.GetAllPurchaces()
	db.Db.Close()
	fmt.Println("\n[Loader] Done loading!")

	// organize result in map: clientId(key) => list [purchases](value)
	preparedData = make(map[int64][]model.EventData, len(customers))
	// add customers with purchaces
	for _, event := range purchaces {
		if preparedData[event.CustomerID] != nil {
			preparedData[event.CustomerID] = append(preparedData[event.CustomerID], event)
		} else {
			preparedData[event.CustomerID] = []model.EventData{event}
		}
	}
	// add customers without purchaces default is empty eventdata with content 0 eur and 0 quantity
	for _, customer := range customers {
		if preparedData[customer.CustomerID] == nil {
			preparedData[customer.CustomerID] = []model.EventData{model.EventData{0, 0, customer.CustomerID, 0, model.EventContent{}}}
		}
	}

	return
}
