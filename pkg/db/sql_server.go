package db

import (
	"go-demo-unit-test/domain/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func GetDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
}

func Migrate() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, _ := GetDB()
	g.UseDB(db) // reuse your gorm db

	db.AutoMigrate(

		entity.Customer{},
		entity.Order{},
		entity.Product{},
		entity.OrderDetail{},
		// &model.DOAJobLevel{},
		// &model.Table{},
		// &model.Node{},

		// &model.NodeDOAJobLevel{},
		// &model.TableDOAJobLevel{},

	)
	g.ApplyInterface(func(Querier) {},
		entity.Customer{},
		entity.Order{},
		entity.Product{},
		entity.OrderDetail{},
	)

	// Generate the code
	g.Execute()

}
