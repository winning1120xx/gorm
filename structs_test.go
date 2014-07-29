package gorm_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/jinzhu/gorm"

	"reflect"
	"time"
)

type Company struct {
	Id   int64
	Name string
}

type Role struct {
	Name string
}

func (role *Role) Scan(value interface{}) error {
	role.Name = string(value.([]uint8))
	return nil
}

func (role Role) Value() (driver.Value, error) {
	return role.Name, nil
}

func (role Role) IsAdmin() bool {
	return role.Name == "admin"
}

type Language struct {
	Id   int
	Name string
}

type Ignored struct {
	Name string
}

type Num int64

func (i *Num) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
	case int64:
		*i = Num(s)
	default:
		return errors.New("Cannot scan NamedInt from " + reflect.ValueOf(src).String())
	}
	return nil
}

type User struct {
	Id                int64
	Age               int64
	UserNum           Num
	Name              string        `sql:"size:255"`
	Birthday          time.Time     // Time
	CreatedAt         time.Time     // CreatedAt: Time of record is created, will be insert automatically
	UpdatedAt         time.Time     // UpdatedAt: Time of record is updated, will be updated automatically
	Emails            []Email       // Embedded structs
	Ignored           Ignored       `sql:"-"`
	BillingAddress    Address       // Embedded struct
	BillingAddressId  sql.NullInt64 // Embedded struct's foreign key
	ShippingAddress   Address       // Embedded struct
	ShippingAddressId int64         // Embedded struct's foreign key
	CreditCard        CreditCard
	Latitude          float64
	CompanyId         int64
	Company
	Role
	PasswordHash      []byte
	IgnoreMe          int64    `sql:"-"`
	IgnoreStringSlice []string `sql:"-"`
}

type UserCompany struct {
	Id        int64
	UserId    int64
	CompanyId int64
}

func (t UserCompany) TableName() string {
	return "user_companies"
}

type CreditCard struct {
	Id        int8
	Number    string
	UserId    sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Email struct {
	Id        int16
	UserId    int
	Email     string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Address struct {
	Id        int
	Address1  string
	Address2  string
	Post      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Product struct {
	Id                    int64
	Code                  string
	Price                 int64
	CreatedAt             time.Time
	UpdatedAt             time.Time
	AfterFindCallTimes    int64
	BeforeCreateCallTimes int64
	AfterCreateCallTimes  int64
	BeforeUpdateCallTimes int64
	AfterUpdateCallTimes  int64
	BeforeSaveCallTimes   int64
	AfterSaveCallTimes    int64
	BeforeDeleteCallTimes int64
	AfterDeleteCallTimes  int64
}

type Animal struct {
	Counter   int64 `primaryKey:"yes"`
	Name      string
	From      string //test reserved sql keyword as field name
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Details struct {
	Id   int64
	Bulk gorm.Hstore
}

type Category struct {
	Id   int64
	Name string
}

type Post struct {
	Id             int64
	CategoryId     sql.NullInt64
	MainCategoryId int64
	Title          string
	Body           string
	Comments       []Comment
	Category       Category
	MainCategory   Category
}

type Comment struct {
	Id      int64
	PostId  int64
	Content string
	Post    Post
}

type Order struct {
}

type Cart struct {
}

func (c Cart) TableName() string {
	return "shopping_cart"
}

type BigEmail struct {
	Id           int64
	UserId       int64
	Email        string
	UserAgent    string
	RegisteredAt time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (b BigEmail) TableName() string {
	return "emails"
}

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}
	nt.Time, nt.Valid = value.(time.Time), true
	return nil
}

func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

type NullValue struct {
	Id      int64
	Name    sql.NullString `sql:"not null"`
	Age     sql.NullInt64
	Male    sql.NullBool
	Height  sql.NullFloat64
	AddedAt NullTime
}
