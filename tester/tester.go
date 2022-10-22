package tester

import (
	"fmt"

	model "github.com/kanounzied/mini-Go-project/model"
	treater "github.com/kanounzied/mini-Go-project/treat"
)

func TestDataLoad(data map[int64][]model.EventData) {
	fmt.Println("\n[Tester] Launching data load test ...")

	for key, value := range data {
		fmt.Println("----------------")
		fmt.Println("Customer id : ", key)
		fmt.Println("----------------")
		for _, event := range value {
			fmt.Printf("|__ Purchase id : %v\n|__ Quantity : %v\n|__ Price in %v : %v \n \n", event.Id, event.Quantity, event.Content.Currency, event.Content.Price)
		}
		fmt.Println()
	}
	fmt.Println("[Tester] Test load done!")

}

func TestDataTreat(quantiles map[int]treater.Quantile) {

	fmt.Println("\n[Tester] Launching data treat test ...")
	for key, value := range quantiles {
		fmt.Println("----------------")
		fmt.Printf("Quantile number %v\n", key)
		fmt.Println("----------------")
		fmt.Println("|__ Number of customers : ", value.Nb)
		fmt.Println("|__ Chiffre d'affaire max : ", value.CAmax)
		fmt.Println("|__ Chiffre d'affaire min : ", value.CAmin)
		fmt.Println()
	}
	fmt.Println("[Tester] Test treat done!")
}
