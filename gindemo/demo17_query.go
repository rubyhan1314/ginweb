package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
)

//定义User类型结构
type User struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

//定义一个getALL函数用于回去全部的信息
func getAll() (users []User, err error) {

	//1.操作数据库
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()

	//2.查询
	rows, err := db.Query("SELECT id, username, password FROM user_info")
	if err != nil {
		log.Fatal(err.Error())
	}

	for rows.Next() {
		var user User
		//遍历表中所有行的信息
		rows.Scan(&user.Id, &user.Username, &user.Password)
		//将user添加到users中
		users = append(users, user)
	}
	//最后关闭连接
	defer rows.Close()
	return
}

//插入数据
func add(user User) (Id int, err error) {

	//1.操作数据库
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO user_info(username, password) VALUES (?, ?)")
	if err != nil {
		return
	}
	//执行插入操作
	rs, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return
	}
	//返回插入的id
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	//将id类型转换
	Id = int(id)
	defer stmt.Close()
	return
}

//修改数据
func update(user User) (rowsAffected int64, err error) {

	//1.操作数据库
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()
	stmt, err := db.Prepare("UPDATE  user_info SET username=?, password=? WHERE id=?")
	if err != nil {
		return
	}
	//执行修改操作
	rs, err := stmt.Exec(user.Username, user.Password, user.Id)
	if err != nil {
		return
	}
	//返回插入的id
	rowsAffected, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	return
}

//通过id删除
func del(id int) (rows int, err error) {
	//1.操作数据库
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM user_info WHERE id=?")
	if err != nil {
		log.Fatalln(err)
	}

	rs, err := stmt.Exec(id)
	if err != nil {
		log.Fatalln(err)
	}
	//删除的行数
	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows = int(row)
	return
}

func main() {

	//创建一个路由Handler
	router := gin.Default()

	//get方法的查询
	router.GET("/user", func(c *gin.Context) {

		users, err := getAll()
		if err != nil {
			log.Fatal(err)
		}
		//H is a shortcut for map[string]interface{}
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})
	})

	//利用post方法新增数据
	router.POST("/add", func(c *gin.Context) {
		var u User
		err := c.Bind(&u)
		if err != nil {
			log.Fatal(err)
		}
		Id, err := add(u)
		fmt.Print("id=", Id)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s 插入成功", u.Username),
		})
	})

	//利用put方法修改数据
	router.PUT("/update", func(c *gin.Context) {
		var u User
		err := c.Bind(&u)
		if err != nil {
			log.Fatal(err)
		}
		num, err := update(u)
		fmt.Print("num=", num)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("修改id: %d 成功", u.Id),
		})
	})

	//利用DELETE请求方法通过id删除
	router.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")

		Id, err := strconv.Atoi(id)

		if err != nil {
			log.Fatalln(err)
		}
		rows, err := del(Id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("delete rows ", rows)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})

	router.Run(":8080")

}
