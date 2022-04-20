package storebe

import (
	"github.com/Sebastian-Soto-M/price-history/backend/database"
	"github.com/Sebastian-Soto-M/price-history/exceptions"
	"github.com/Sebastian-Soto-M/price-history/models"
)

var stPanic = exceptions.Panic("Store")

// func (repo StoreRepository) getById(id int) (store models.Store) {
//   return nil
// }
func GetAll() []models.Store {
	var stores []models.Store
	db := database.Connect()
	query := db.Find(&stores)
	result, err := query.Rows()
	defer result.Close()
	stPanic(err)
	return stores
}

func Initialize() {
	db := database.Connect()
	db.Create(models.Amazon.Store())
	db.Create(models.Ebay.Store())
	db.Create(models.Wish.Store())
}
