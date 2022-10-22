package treat

import (
	"fmt"
	"math"
	"sort"

	model "github.com/kanounzied/mini-Go-project/model"
)

const (
	CURRENCY_EUR = 1    // consensus (facteur == 1)
	CURRENCY_USD = 1.03 // on peut ajouter un service pour un appel api et mettre a jour les currency
	QUANTILE     = 2.5  // quantile fixéé à 2.5% d'aprés le cahier de charge
)

var (
	keys_for_sort []int64 // this array is to be used for map sorting for customers CA
)

type Quantile struct {
	Nb    int
	CAmax float32
	CAmin float32
}

// this function transforms all currencies to a signle consensus (eur in this case because its unit factor is 1)
func handleCurrency(content model.EventContent) (converted_price float32) {
	switch content.Currency {
	case "eur":
		converted_price = content.Price * CURRENCY_EUR
	case "usd":
		converted_price = content.Price * CURRENCY_USD
	}
	return
}

func calculCA(purchaces []model.EventData) (CA float32) {
	CA = 0
	for _, purchace := range purchaces {
		CA += float32(purchace.Quantity) * handleCurrency(purchace.Content)
	}
	return
}

func GenerateCustomersCAMap(customersPurchaces map[int64][]model.EventData) (res map[int64]float32) {
	keys_for_sort = make([]int64, 0, len(customersPurchaces))
	fmt.Println("[Treater] Generating Customers CA map ...")
	fmt.Println("-----------------\n")

	// counteur pour logger 10 entrée de map
	logCounter := 0
	res = make(map[int64]float32, len(customersPurchaces))
	for customerID, purchaces := range customersPurchaces {
		keys_for_sort = append(keys_for_sort, customerID) // filling keys here to optimize generation of the array
		logCounter++
		CA := calculCA(purchaces)
		res[customerID] = CA
		if logCounter <= 10 {
			fmt.Printf("[Treater] Customer %v => %v \n", customerID, CA)
		}
	}

	fmt.Println("-----------------")
	fmt.Println("\n[Treater] Done generating!")
	return
}

func prepareQuantile(quantile Quantile, memberCA float32) Quantile {
	if quantile.Nb == 0 {
		quantile.CAmax = memberCA
		quantile.CAmin = memberCA
	} else {
		if memberCA > quantile.CAmax {
			quantile.CAmax = memberCA
		}
		if memberCA < quantile.CAmin {
			quantile.CAmin = memberCA
		}
	}
	quantile.Nb++
	return quantile
}

func GenerateQuantiles(customerCaMap map[int64]float32) map[int]Quantile {

	fmt.Println("\n[Treater] Generating QUantiles map ...")
	// sort keys array in descendent order following customerCA
	sort.SliceStable(keys_for_sort, func(i, j int) bool {
		return customerCaMap[keys_for_sort[i]] > customerCaMap[keys_for_sort[j]]
	})

	// N : number of clients per quantile
	// nbQuantile : nb of total quantiles
	var (
		N          int
		nbQuantile int
	)
	// if nbCust < 100 each quantile has N=1 customers
	if len(customerCaMap) < 100 {
		N = 1
		nbQuantile = len(customerCaMap)
	} else {
		nbQuantile = int(math.Ceil(100.0 / QUANTILE))
		N = int(math.Ceil(float64(len(customerCaMap)) / float64(nbQuantile)))
	}
	// generate quantiles and insert them in map
	var (
		quantileIndex int              = 0
		quantilesMap  map[int]Quantile = make(map[int]Quantile, nbQuantile)
	)
	for _, custID := range keys_for_sort {
		if quantilesMap[quantileIndex].Nb >= N {
			quantileIndex++
		}
		quantilesMap[quantileIndex] = prepareQuantile(quantilesMap[quantileIndex], customerCaMap[custID])
	}

	fmt.Println("[Treater] Done generating!")
	return quantilesMap
}
