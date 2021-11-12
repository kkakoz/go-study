package relate

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("root:admin@(localhost:3306)/test")),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		log.Fatalln(err)
	}
	m.Run()
}

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	ClassId   int64      `json:"class_id"`
	Class     Class      `json:"class"`
	Languages []Language `json:"language" gorm:"many2many:user_languages"`
}

type Language struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Class struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Rooms []Room `json:"rooms"`
}

type Room struct {
	ID      int64  `json:"id"`
	ClassId int64  `json:"class_id"`
	Name    string `json:"name"`
}

func TestMany(t *testing.T) {
	err := db.AutoMigrate(&User{}, &Class{}, &Room{})
	if err != nil {
		t.Fatal(err)
	}
	//db.Create(&User{
	//	Name: "zhangsan",
	//	Class: Class{
	//		Name: "class1",
	//		Rooms: []Room{{
	//			Name: "room1",
	//		}, {
	//			Name: "room2",
	//		}},
	//	},
	//})
	user := &User{}
	err = db.Preload("Class").Preload("Class.Rooms", "name like ?", "%room%").First(user).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}

func TestManyMany(t *testing.T) {
	err := db.AutoMigrate(&User{}, &Language{})
	if err != nil {
		t.Fatal(err)
	}
	//db.Create(&User{
	//	Name: "lisi",
	//	Class: Class{
	//		Name: "class2",
	//	},
	//	Language: []Language{
	//		{
	//			Name: "language1",
	//		},
	//		{
	//			Name: "language2",
	//		},
	//	},
	//})
	user := &User{}
	err = db.Preload("Languages").Find(user, 2).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}

func TestJsonMap(t *testing.T) {
	m := map[string]interface{}{
		"1": "1",
		"2": nil,
		"3": map[string]interface{}{
			"1": 1,
			"2": nil,
		},
	}
	bytes, _ := json.Marshal(m)
	fmt.Println(string(bytes))
}