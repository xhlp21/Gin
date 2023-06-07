package Model

import (
	"fmt"
	"gin/Config"
	"log"
)

type Test struct {
	Id   int    `json:"id"`
	Test string `json:"test"`
}

// 查询单条数据
func (t *Test) Info() (test Test) {

	//调用连接数据库方法，并使用 QueryRow 获取单挑数据并用 Scan 遍历出来
	err := datebles.InitDB().QueryRow("select * from test where id = ?", t.Id).Scan(
		&test.Id, &test.Test)
	if err != nil {
		return Test{}
	}
	return
}

// 查询所有的数据
func (t *Test) All() (test []Test) {

	//调用连接数据库方法，并使用 Query 获取全部数据进行循环遍历数据
	rows, err := datebles.InitDB().Query("select * from test")
	if err != nil {
		log.Fatal(err)
	}

	// Next的作用用来判断是否存在下一行数据，返回布尔类型 true 或 false
	for rows.Next() {
		t := Test{}
		err := rows.Scan(&t.Id, &t.Test)
		if err != nil {
			log.Fatal(err)
		}

		//append 的作用是将数据追加到切片中
		test = append(test, t)
	}

	// Close 关闭查询结果集，防止资源泄露
	rows.Close()
	return
}

// 执行新增操作
func (t *Test) Insert() int {

	//调用连接数据库方法，并使用 Exec 进行新增操作
	rs, err := datebles.InitDB().Exec("insert into test(test) value(?)", t.Test)
	if err != nil {
		log.Fatal(err)
	}

	// LastInsertId 的作用是返回新增这条数据的 ID
	res, err := rs.LastInsertId()
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
	return t.Id
}

// 执行删除操作
func (t *Test) Delete() int64 {

	//调用连接数据库方法，并使用 Exec 进行删除操作
	rs, err := datebles.InitDB().Exec("delete from test where id = ?", t.Id)
	if err != nil {
		log.Fatal(err)
	}

	// RowsAffected 的作用是返回处结果后影响了几条数据
	res, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return res
}

// 执行修改操作
func (t *Test) Update() int64 {

	//调用连接数据库方法，并使用 Exec 进行修改操作
	rs, err := datebles.InitDB().Exec("update test set test = ? where id = ?", t.Test, t.Id)
	if err != nil {
		log.Fatal(err)
	}

	// RowsAffected 的作用是返回处结果后影响了几条数据
	res, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return res
}
