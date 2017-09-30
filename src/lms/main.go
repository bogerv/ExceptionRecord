package main

import (
	"database/sql"
	"fmt"
	_ "lms/routers"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int
	UserName string `form:"username" json:"username" binding:"required"`
	URL      string `form:"url" json:"url" bdinding:"required"`
	Age      int    `form:"age" json:"age"`
}

var db *sql.DB

func main() {
	/*Open MySql*/
	var err error
	db, err = sql.Open("mysql", "root:Mit000@tcp(127.0.0.1:3306)/bogerv?parseTime=true")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	/*Open MySql*/

	/*Config Router*/
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/user", func(c *gin.Context) {
		var usr user
		users, err := usr.getUsers()
		if err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	router.GET("/user/info/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user user
		err := db.QueryRow("SELECT id, UserName, URL,Age FROM users WHERE id=?", id).Scan(
			&user.ID, &user.UserName, &user.URL, &user.Age,
		)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"user": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/user/add", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addUser.html", gin.H{})
	})

	router.POST("/user/add", func(c *gin.Context) {
		var user user
		user.UserName = c.Request.FormValue("UserName")
		user.URL = c.Request.FormValue("Url")
		user.Age, _ = strconv.Atoi(c.Request.FormValue("Age"))
		id, err := user.addUser()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert person Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})

	router.Run(":8000")
}

func (u *user) addUser() (id int64, err error) {
	rs, err := db.Exec("INSERT INTO users(UserName, Url,Age) VALUES (?, ?, ?)", u.UserName, u.URL, u.Age)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (u *user) getUsers() (users []user, err error) {
	users = make([]user, 0)
	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var usr user
		err := rows.Scan(&usr.ID, &usr.UserName, &usr.URL, &usr.Age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, usr)
	}

	if err = rows.Err(); err != nil {
		return
	}
	return
}
