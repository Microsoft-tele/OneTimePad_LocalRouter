package Otp

import (
	"database/sql"
	"fmt"
)

var (
	err error
)

type Otp struct {
	Id    string
	Url   string
	Intro string
	Uname string
	Pwd   string
	Email string
	Db    *sql.DB
}

func (o *Otp) InitMysql() {
	o.Db, err = sql.Open("mysql", "root:660967@tcp(192.168.88.131:3306)/users") // @ TODO 记得修改数据库地址，最好是固定一个ip
	if err != nil {
		fmt.Println("初始化数据库失败:", err)
	}
}

func (o *Otp) AddOtp(url string, intro string, uname string, pwd string, email string) {
	o.InitMysql()
	prepare, err := o.Db.Prepare("insert into PwdTable(Url,Intro,Uname,Pwd,Email) values (?,?,?,?,?)")
	if err != nil {
		fmt.Println("预编译sql失败:", err)
		return
	}
	_, err = prepare.Exec(url, intro, uname, pwd, email)
	if err != nil {
		fmt.Println("插入数据库失败:", err)
		return
	}
	fmt.Println("插入数据库成功")
}

func (o *Otp) SelectOtp(email string) ([]Otp, error) {
	var OtpObjs []Otp
	prepare, err := o.Db.Prepare("select * from PwdTable where email=?")
	if err != nil {
		fmt.Println("预编译sql失败:", err)
		return nil, err
	}
	query, err := prepare.Query(email)
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			fmt.Println("关闭query失败:", err)
		}
	}(query)
	if err != nil {
		fmt.Println("执行sql是失败:", err)
		return nil, err
	}
	for query.Next() {
		var tmp Otp
		err = query.Scan(&tmp.Id, &tmp.Url, &tmp.Intro, &tmp.Uname, &tmp.Pwd, &tmp.Email)
		OtpObjs = append(OtpObjs, tmp)
	}
	return OtpObjs, nil
}
