package up

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tus/tusd"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Up struct {
	Stop chan os.Signal

	db     *gorm.DB
	router *gin.Engine
	tusd   *tusd.StoreComposer
}

func New(config *Config) (*Up, error) {
	router := gin.Default()

	db, err := gorm.Open(config.DbProvider, config.DbAddress)
	if err != nil {
		return nil, err
	}

	store := filestore.Filestore{Path: config.DataPath}
	tusd := tusd.NewStoreComposer()
	store.UseIn(tusd)

	up := &Up{
		Stop: make(chan os.Signal, 1),

		db:     db,
		router: router,
		tusd:   tusd,
	}

	up.route()
	return up, nil
}

func (up *Up) Run() {
	up.router.Run()
}

func (up *Up) Shutdown() {
	up.db.Close()
}
