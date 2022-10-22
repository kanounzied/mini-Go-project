package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	model "github.com/kanounzied/mini-Go-project/model"
)

var (
	dbUser string
	dbPass string
	dbUrl  string
	dbName string
	Db     *sql.DB
)

func loadEnvVariables() error {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return err
	}
	dbUser = os.Getenv("DBUSER")
	dbPass = os.Getenv("PASSWORD")
	dbUrl = os.Getenv("DBURL")
	dbName = os.Getenv("DBNAME")
	fmt.Println("\n[DB] Environment variables loaded successfully!")
	return nil
}

func Connect() {

	fmt.Println("\n[DB] Connecting to database")
	err := loadEnvVariables()
	if err != nil {
		fmt.Println("could not load env variables!")
	}
	d, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbUrl+":3306)/"+dbName)
	if err != nil {
		log.Fatal("Unable to open connection to DB: ", err)
	}
	Db = d
	fmt.Println("\n[DB] Success!")
}

func GetDB() *sql.DB {
	if Db == nil {
		Connect()
	}
	return Db
}

func GetAllCustomers() (customers []model.Customer) {

	db := GetDB()
	fmt.Println("\n[DB] Scanning customers table...")
	res, err := db.Query("select CustomerID from Customer")
	if err != nil {
		log.Fatal(err)
	}
	customers = make([]model.Customer, 0)
	for res.Next() {
		var custID int64
		err = res.Scan(&custID)
		customers = append(customers, model.NewCustomer(custID))
	}
	fmt.Println("\n[DB]  Done!")
	return
}

func GetContent(contentID int) model.EventContent {
	db := GetDB()
	var (
		Id       int
		Price    float32
		Currency string
	)
	err := db.QueryRow("select C.ContentID, Price, Currency from Content C, ContentPrice CP where C.ContentID = ? and CP.ContentID = C.ContentID", contentID).Scan(&Id, &Price, &Currency)
	if err != nil {
		log.Fatal(err)
	}
	return model.NewEventContent(Id, Price, Currency)
}

func GetAllPurchaces() (purchaces []model.EventData) { // purchaces c à d type = 6

	db := GetDB()
	res, err := db.Query("select EventDataID, ContentID, CustomerID, Quantity from CustomerEventData where EventTypeID = 6")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n[DB] Scanning purchaces table ...")
	purchaces = make([]model.EventData, 0)
	for res.Next() {
		var (
			id         int64
			contentID  int
			customerID int64
			quantity   int16
		)
		res.Scan(&id, &contentID, &customerID, &quantity)
		eventData := model.NewEventData(id, contentID, customerID, quantity)
		content := GetContent(contentID)
		eventData.Content = content
		purchaces = append(purchaces, eventData)
	}
	fmt.Println("\n[DB] Done!")
	return
}
