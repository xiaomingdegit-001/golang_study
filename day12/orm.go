package myorm

/**
 * orm => gorm
 * 参考 https://gorm.io/zh_CN
 */
import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
)

type Book struct {
	Id     int64  `json:"id,omitempty"`
	Url    string `json:"url,omitempty"`
	Remark string `json:"remark,omitempty"`
	Type   uint   `json:"type,omitempty"`
}

func connMysql() (*gorm.DB, error) {
	dsn := "root:Ysj666tY#iM@tcp(110.42.206.132:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`User`表为`t_users`
			SingularTable: true,  // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	})
	return db, err
}

func QueryBook(w http.ResponseWriter, r *http.Request) {
	data, err := Query2Json()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, fmtErr := fmt.Fprintln(w, data)
	if fmtErr != nil {
		fmt.Println(fmtErr)
		return
	}
}

func Query() ([]Book, error) {
	db, err := connMysql()
	if err != nil {
		return nil, err
	}
	var books []Book
	db.Find(&books)
	return books, err
}

func Query2Json() (string, error) {
	books, err := Query()
	if err != nil {
		return "", err
	}
	j, err := json.Marshal(books)
	return string(j), err
}

func Start(handler func(http.ResponseWriter, *http.Request)) error {
	http.HandleFunc("/book", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
