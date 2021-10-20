package main

/**
 * orm => gorm
 * 参考 https://gorm.io/zh_CN
 */
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Book struct {
	Id     int64
	Url    string
	Remark string
	Type   uint
}

func main() {
	dsn := "root:Ysj666tY#iM@tcp(110.42.206.132:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`User`表为`t_users`
			SingularTable: true,  // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	var book Book
	db.First(&book)
	fmt.Printf("book=[id: %d, url: %s, remark: %s, type: %d]",
		book.Id, book.Url, book.Remark, book.Type)
}
