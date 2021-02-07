package reader

import (
	"log"
	"time"

	"github.com/bigwhite/testdboper/pkg/db"
	"github.com/bigwhite/testdboper/pkg/model"
)

func dumpEmployee() {
	var rs []model.Employee
	d := db.DB()
	d.Find(&rs)
	log.Println(rs)
}

func Run(quit <-chan struct{}) {
	tk := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-tk.C:
			dumpEmployee()

		case <-quit:
			return
		}
	}
}
