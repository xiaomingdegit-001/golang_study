package main

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

func queryBook(w http.ResponseWriter, r *http.Request) {
	data, err := query()
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

func query() (string, error) {
	db, err := connMysql()
	if err != nil {
		return "", err
	}
	var book Book
	db.First(&book)
	fmt.Printf("book=[id: %d, url: %s, remark: %s, type: %d]\n",
		book.Id, book.Url, book.Remark, book.Type)
	j, err := json.Marshal(book)
	return string(j), err
}

func main() {
	http.HandleFunc("/getBook", queryBook)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
