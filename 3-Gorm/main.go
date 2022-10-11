package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func main() {

	// 	By Pasakorn Limchuchua

	//	Main Reference
	// 	-> https://github.com/codebangkok/golang
	//	-> https://www.youtube.com/watch?v=yZQRPATbASM&ab_channel=CodeBangkok

	//	GORM
	//	-> https://gorm.io/docs/
	// 	-> https://pkg.go.dev/gorm.io/gorm@v1.23.8

	//	Driver
	//	-> https://pkg.go.dev/gorm.io/driver/mysql
	//	-> https://pkg.go.dev/github.com/go-sql-driver/mysql / https://github.com/go-sql-driver/mysql#usage

	//	Docker
	//	-> https://hub.docker.com/_/mysql

	//	Extensions
	// 	-> Go, Error Lens, MySQL

	// 	Commands
	// 	go mod init [module_name]		-> create modules
	// 	go run [file_name].go			-> run file
	//	go run .						-> run file
	//	go get [url_path]				-> install package
	//	docker pull [image_name]		-> install image

	//	install package					-> go get gorm.io/gorm
	//									-> go get gorm.io/driver/mysql
	//	install image					-> docker pull mysql
	//	use docker (with yml)			-> docker-compose up -d
	//	connection						-> Host : 127.0.0.1 / port : 3306 / username : root / password : pass
	//									-> execute testdb

	/* 	--------------- Gorm Library ---------------	*/

	// 	Gorm Convention (https://gorm.io/docs/conventions.html)
	// 	1. ID -> primary key
	// 	2. Pluralized Table Name + snake_case
	//	3. Column Name + snake_case

	// 	dsn = data source name (https://github.com/go-sql-driver/mysql#dsn-data-source-name)
	dsn := "root:pass@tcp(127.0.0.1:3306)/testdb?parseTime=True"
	//  dialector = driver
	dial := mysql.Open(dsn)

	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{}, // -> logs sql command
		DryRun: false,        // -> false = really execute in database / true = otherwise
	})
	if err != nil {
		panic(err)
	}

	// db.Migrator().CreateTable(Test{})

	db.AutoMigrate(&Gender{}, &Test{}, Customer{})

	CreateGender("male")
	CreateGender("female")
	CreateGender("xxxx")
	GetGenders()
	GetGender(1)
	GetGenderByName1("male")
	GetGenderByName2("female")
	UpdateGender1(3, "yyyy")
	UpdateGender2(3, "zzzz")
	DeleteGender(3)

	CreateTest(0, "Test1")
	CreateTest(0, "Test2")
	CreateTest(0, "Test3")
	DeleteTest(3)

	CreateCustomer("Best", 1)
	CreateCustomer("Noon", 2)
	GetCustomers()
}

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n====================================\n", sql)
}

type Test struct { // -> can config gorm convention by setting field tags `gorm:` (https://gorm.io/docs/models.html)
	gorm.Model        // -> embed ID, CreatedAt, UpdatedAt, DeletedAt
	Code       uint   `gorm:"comment:This is code"`
	Name       string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
}

func (t Test) TableName() string { // can change tablename in return (https://gorm.io/docs/conventions.html#TableName)
	return "MyTest" // -> change from Test to MyTest
}

type Gender struct {
	ID   int `gorm:"unique;size(10)"`
	Name string
}

// 	========================= CRUD =========================

// 	Create -> (https://gorm.io/docs/create.html)

func CreateGender(name string) { // INSERT INTO `genders` (`name`) VALUES ('female')
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

//	Query -> (https://gorm.io/docs/query.html)

func GetGenders() { // SELECT * FROM `genders` ORDER BY id
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGender(id uint) { // SELECT * FROM `genders` WHERE `genders`.`id` = 1 ORDER BY `genders`.`id` LIMIT 1
	gender := Gender{}
	tx := db.First(&gender, id) // -> where by id (using ,)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

func GetGenderByName1(name string) { // SELECT * FROM `genders` WHERE name='male' ORDER BY id
	genders := Gender{}
	tx := db.Order("id").Find(&genders, "name=?", name) // -> where clause
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

func GetGenderByName2(name string) { // SELECT * FROM `genders` WHERE name='female'
	genders := Gender{}
	tx := db.Where("name=?", name).Find(&genders) // -> where clause
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

//	update -> (https://gorm.io/docs/update.html)

func UpdateGender1(id uint, name string) { // SELECT * FROM `genders` WHERE `genders`.`id` = 3 ORDER BY `genders`.`id` LIMIT 1
	//									   // UPDATE `genders` SET `name`='yyyy' WHERE `id` = 3
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

func UpdateGender2(id uint, name string) { // UPDATE `genders` SET `name`='zzzz' WHERE id=3
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender) // -> beware zero value !! (will not work)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

//	delete -> (https://gorm.io/docs/delete.html)

func DeleteGender(id uint) { // DELETE FROM `genders` WHERE `genders`.`id` = 3
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
}

//	soft delete vs hard delete -> If model includes a gorm.DeletedAt field (which is included in "gorm.Model"), it will get soft delete ability automatically!
//							   -> (https://gorm.io/docs/delete.html#Soft-Delete)

func CreateTest(code uint, name string) { // INSERT INTO `MyTest` (`created_at`,`updated_at`,`deleted_at`,`code`,`myname`) VALUES ('2022-09-26 14:37:13.139','2022-09-26 14:37:13.139',NULL,0,'Test3')
	test := Test{Code: code, Name: name}
	db.Create(&test)
}

func GetTests() {
	tests := []Test{}
	db.Find(&tests)
	for _, t := range tests {
		fmt.Printf("%v|%v\n", t.ID, t.Name)
	}
}

func DeleteTest(id uint) { // DELETE FROM `MyTest` WHERE `MyTest`.`id` = 3
	// db.Delete(&Test{}, id) -> Soft delete (will mark as delete_at but not really delete)
	db.Unscoped().Delete(&Test{}, id) // -> Delete permanently
}

// Real Case : Customer

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender // -> associate with genders table
	GenderID uint
}

func CreateCustomer(name string, genderID uint) { // INSERT INTO `customers` (`name`,`gender_id`) VALUES ('Noon',2)
	customer := Customer{Name: name, GenderID: genderID}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

func GetCustomers() { // SELECT * FROM `genders` WHERE `genders`.`id` IN (1,2)
	//				  // SELECT * FROM `customers`
	customers := []Customer{}
	tx := db.Preload(clause.Associations).Find(&customers) // -> use preload to pull associate table (https://gorm.io/docs/preload.html#Preload-All)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Printf("%v|%v|%v\n", customer.ID, customer.Name, customer.Gender.Name)
	}
}

// 	However we can raw query normally through gorm (https://gorm.io/docs/sql_builder.html)
