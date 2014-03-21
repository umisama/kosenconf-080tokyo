package main

import (
	"crypto/md5"
	"crypto/sha1"
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"time"
)

type User struct {
	Name       string `db:"id" json:"name"`
	ScreenName string `db:"screen_name" json:"screen_name"`
	Password   string `db:"password" json:"-"`
}

func hash(val, uname string) string {
	// Create salt
	salth := md5.New()
	io.WriteString(salth, uname)
	salt := string(salth.Sum(nil))

	// Create hashed password
	h := sha1.New()
	io.WriteString(h, val+salt)
	return string(h.Sum(nil))
}

func (u *User) SetPassword(val, uname string) {
	u.Password = hash(val, uname)
}

func (u *User) IsMatchPassword(val, uname string) bool {
	return (u.Password == hash(val, uname))
}

type Status struct {
	Id        int       `db:"id" json:"-"`
	UserName  string    `db:"user_id" json:"user"`
	Content   string    `db:"content" json:"content"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
}

func CreateUser(name, screen_name, password string) (err error) {
	u := &User{
		Name:       name,
		ScreenName: screen_name,
	}
	u.SetPassword(password, name)

	err = dbmap.Insert(u)
	return
}

func GetUser(name, password string) (user *User, err error) {
	user_raw, err := dbmap.Get(&User{}, name)
	if err != nil || user_raw == nil {
		return
	}
	user = user_raw.(*User)

	if !user.IsMatchPassword(password, name) {
		user = nil
		return
	}
	return
}

func GetUserDetail(name string) (screen_name string, err error) {
	user_raw, err := dbmap.Get(&User{}, name)
	if err != nil || user_raw == nil {
		return
	}
	user := user_raw.(*User)

	return user.ScreenName, nil
}

func CreateStatus(user, text string) (err error) {
	err = dbmap.Insert(&Status{
		UserName:  user,
		Content:   text,
		Timestamp: time.Now(),
	})
	return
}

func GetStatuses(count int) (ret []Status, err error) {
	ret = make([]Status, 0)
	_, err = dbmap.Select(&ret, `select * from "STATUS" order by "timestamp" desc limit ?`, count)
	return
}

func SearchStatuses(q string, count int) (ret []Status, err error) {
	ret = make([]Status, 0)
	_, err = dbmap.Select(&ret, `select * from "STATUS" where "content" like ? order by "timestamp" desc limit ?`, "%"+q+"%", count)
	return
}

func connectToDb() (dbmap *gorp.DbMap, err error) {
	sqlite, err := sql.Open("sqlite3", "twittor.db")
	if err != nil {
		return
	}

	dbmap = &gorp.DbMap{
		Db:      sqlite,
		Dialect: gorp.SqliteDialect{},
	}

	dbmap.AddTableWithName(User{}, "USER").SetKeys(false, "id")
	dbmap.AddTableWithName(Status{}, "STATUS").SetKeys(true, "id")

	err = dbmap.CreateTablesIfNotExists()
	return
}
