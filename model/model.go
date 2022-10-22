package model

type Customer struct {
	CustomerID int64
	// ChannelValue string
}

func NewCustomer(id int64) (customer Customer) {
	customer = Customer{CustomerID: id}
	return
}

type EventData struct { // type qui rassemble les attributs necessaires pour le calcul du CA
	Id         int64
	ContentID  int
	CustomerID int64
	Quantity   int16
	Content    EventContent
}

func NewEventData(id int64, contentID int, customerID int64, quantity int16) (eventData EventData) {
	eventData = EventData{Id: id, ContentID: contentID, CustomerID: customerID, Quantity: quantity}
	return
}

type EventContent struct {
	Id       int
	Price    float32
	Currency string
}

func NewEventContent(id int, price float32, currency string) (eventContent EventContent) {
	eventContent = EventContent{Id: id, Price: price, Currency: currency}
	return
}
