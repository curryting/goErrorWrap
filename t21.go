// Q: 在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// A: 个人认为方法有很多，不能限定某一种方法，具体需要根据公司架构来做选择。
//下面演示使用wrap方法的demo

package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var NotFound = errors.New("not found")

func Biz() error {
	err := Dao("")
	if errors.Is(err, NotFound) {
		// 根据业务角度进行处理
		return nil
	}
	if err != nil {
		// 数据库查询报错
	}
	return nil
}

func Dao(query string) error {
	err := mockError()
	if err == sql.ErrNoRows {
		return errors.Wrapf(NotFound, fmt.Sprintf("data not found, sql: %s", query))
	}
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("db query system error, sql: %s", query))
	}

	return nil
}
