package main

import (
	"github.com/gin-gonic/gin"
	//	"github.com/ziutek/mymysql/mysql"
	//"github.com/ziutek/mymysql/mysql"
	//_ "github.com/ziutek/mymysql/native" // Native engine
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		_mainDb := new(MainDb)
		err := _mainDb.Init("test")
		if err != nil {
			c.String(401, "Init: "+err.Error())

		} else {
			defer _mainDb.Dispose()
		}

		rows, err := _mainDb.ExecuteQuery("SELECT * FROM employee")
		if err != nil {
			c.String(401, "Execute: "+err.Error())
		}
		for _, row := range rows {
			c.JSON(200, gin.H{"Id": row.Str(0), "Name": row.Str(1)})
		}

		//This also works
		//err1 := init_("temp")
		//if err1 != nil {
		//	c.String(401, err1.Error())
		//}
		//rows, err := ExecuteQuery_("SELECT * FROM employee")
		//if err != nil {
		//	c.String(401, "Error while executing query")
		//}
		//for _, row := range rows {
		//	c.JSON(200, gin.H{"Id": row.Str(0), "Name": row.Str(1)})
		//}

		//Works.................................................................
		//......................................................................
		//db := mysql.New("tcp", "", "127.0.0.1:3306", "", "", "test")
		//err := db.Connect()
		//if err != nil {
		//	panic(err)
		//}
		//defer db.Close()
		////rows, res, err := db.Query("SELECT * FROM employee where Id = %d", 2)
		//rows, _, err := db.Query("SELECT * FROM employee")
		//if err != nil {
		//	panic(err)
		//}

		//for _, row := range rows {
		//	c.JSON(200, gin.H{"Id": row.Str(0), "Name": row.Str(1)})
		//}

	})
	r.Run(":8080")
}
