package backend

import (
	"os"

	"github.com/Sebastian-Soto-M/price-history/backend/database"
	"github.com/Sebastian-Soto-M/price-history/backend/store"
)

func CleanDatabase() {
	os.Remove(database.DBNAME)
	database.Migrate()
	storebe.Initialize()
}
