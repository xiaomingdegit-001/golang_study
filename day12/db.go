package dbtool

/**
 * orm => gorm
 * 参考 https://gorm.io/zh_CN
 */
import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Book struct {
	Id     int64  `json:"id,omitempty"`
	Url    string `json:"url,omitempty"`
	Remark string `json:"remark,omitempty"`
	Type   uint   `json:"type,omitempty"`
}

func ConnMysql() (*gorm.DB, error) {
	dsn := "root:Ysj666tY#iM@tcp(110.42.206.132:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`User`表为`t_users`
			SingularTable: true,  // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	return db, err
}

func Query(db *gorm.DB) []Book {
	var books []Book
	db.Find(&books)
	return books
}

func Query2Json(db *gorm.DB) (string, error) {
	books := Query(db)
	j, err := json.Marshal(books)
	return string(j), err
}
