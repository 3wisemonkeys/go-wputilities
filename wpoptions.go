package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your required driver
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}
type Option struct {
    OptionId    string `orm:"pk"`
    OptionName  string
    OptionValue string
    Autoload    string
}

func (u *Option) TableName() string {
    return "wp_options"
}

func init() {
    // register model
    orm.RegisterModel(new(Option))

    // set default database
    orm.RegisterDataBase("default", "mysql", "rc2:joatmos@/wp?charset=utf8", 30)
}

func main() {
    o := orm.NewOrm()
    var options = []*Option{}
    num, err := o.QueryTable("wp_options").Filter("option_name__contains", "site").All(&options)
    if err == nil {
        fmt.Printf("%d options read\n", num)
        for _, option := range options {
            fmt.Println("*******************************************************")
            fmt.Printf("Id: %s, Name: %s Value:%s\n", option.OptionId, option.OptionName, option.OptionValue)
            fmt.Println("*******************************************************")
        }
    }

}
