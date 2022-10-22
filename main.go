package main

import (
	db "github.com/kanounzied/mini-Go-project/db"
	load "github.com/kanounzied/mini-Go-project/load"
	"github.com/kanounzied/mini-Go-project/tester"
	treater "github.com/kanounzied/mini-Go-project/treat"
)

func main() {

	// call data base to load data
	db := db.GetDB()
	customersPurchaces := load.LoadData()
	//tester.TestDataLoad(customersPurchaces)

	// treat data
	custormersMap := treater.GenerateCustomersCAMap(customersPurchaces)
	quantiles := treater.GenerateQuantiles(custormersMap)
	tester.TestDataTreat(quantiles) // afficher le contenu de quantiles

	// export data

	db.Close()
}
