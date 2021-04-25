package homework

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

// 题目：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

// 解答：dao层遇到sql.ErrNoRows不应该Wrap错误，而是将错误处理并且返回nil，因为dao属于数据传输层，提供的服务就是数据查询，调用此服务器的上层可能是前端页面，他们要的只是结果而不是一个error对象

// 为此不应该向上抛而是直接将此error在dao层处理

// 举例：

type MockClass struct {
}

func mockQuery(params string) (MockClass, error) {
	mockSql := "SELECT * from T WHERE condition=?"
	var ans MockClass
	var db *sql.DB
	err := db.QueryRow(mockSql, params).Scan(&ans)
	if err != nil && err != sql.ErrNoRows {
		return ans, errors.WithMessage(err, fmt.Sprintf("sql: %v occurred an error\n", mockSql))
	}
	return ans, nil
}

func main() {
	ans, err := mockQuery("")
	if err != nil {
		fmt.Printf("&v\n", err)
		return
	}
	// do something
	fmt.Printf("%T\n", ans)
}
