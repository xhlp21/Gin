package Controller

import (
	"gin/App/Model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 查询单条数据
func Info(c *gin.Context) {

	//获取 url 传递的参数
	ids := c.Request.FormValue("id")

	//进行字符串到整数的转换
	id, _ := strconv.Atoi(ids)

	//将ID传递到模型层的结构体中
	t := Model.Test{
		Id: id,
	}

	//调用方法
	Info := t.Info()

	//将数据进行json的打印
	c.JSON(http.StatusOK, Info)
}

// 查询所有的数据
func All(c *gin.Context) {
	t := Model.Test{}
	All := t.All()
	c.JSON(http.StatusOK, All)
}

// 执行新增操作
func Insert(c *gin.Context) {
	t := Model.Test{
		Test: "呵呵",
	}
	t.Insert()
	c.JSONP(http.StatusOK, gin.H{"msg": "添加成功"})
}

// 执行删除操作
func Delete(c *gin.Context) {

	//获取路由传递过来的参数
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	t := Model.Test{
		Id: id,
	}
	t.Delete()
	c.JSON(http.StatusOK, "删除成功")
}

// 执行修改操作
func Update(c *gin.Context) {
	t := Model.Test{
		Id:   1,
		Test: "波风水门",
	}
	t.Update()
	c.JSON(http.StatusOK, "修改成功")
}
