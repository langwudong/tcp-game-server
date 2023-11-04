package mysql

import (
	"database/sql"
	"game_server/src/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

const (
	sqlUrl = "root:langwudong@tcp(127.0.0.1:3306)/users"
)

func InitDB() *sql.DB {
	once.Do(func() {
		db, _ = sql.Open("mysql", sqlUrl)

		//连接数据库
		err := db.Ping()

		if err != nil {
			logrus.Error(err)
		}

		//设置数据库连接池的最大连接数目
		db.SetMaxOpenConns(100)
	})

	return db
}

func CheckUser(user data.User, isAll bool) bool {
	//初始化数据库
	InitDB()

	var sqlStr string
	var userRow data.User

	sqlStr = "SELECT password FROM user WHERE username = ?"

	row := db.QueryRow(sqlStr, user.Username)

	if !isAll && row != nil {
		return true
	} else if !isAll && row == nil {
		return false
	}

	err := row.Scan(&userRow.Password)
	if err != nil {
		logrus.Error(err)
		return false
	}

	if ComparePwd(user.Password, userRow.Password) {
		return true
	}

	return false
}

func AddUser(user data.User) {
	//初始化数据库
	InitDB()

	stmt, err := db.Prepare("INSERT INTO user (username,password,name,email) VALUES (?,?,?,?)")

	defer stmt.Close()

	password, err := EncryptPwd(user.Password)
	if err != nil {
		logrus.Error(err)
		return
	}

	_, err = stmt.Exec(user.Username, password, user.Name, user.Email)

	if err != nil {
		logrus.Error(err)
	}
}

func UpdateUser(user data.User, newPassword string) bool {
	dbw := InitDB()

	stmt, err := dbw.Prepare("UPDATE user SET password=? WHERE username=? AND password=?")

	defer stmt.Close()

	result, err := stmt.Exec(newPassword, user.Username, user.Password)

	if err != nil {
		logrus.Error(err)
		return false
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return false
	} else {
		return true
	}
}
