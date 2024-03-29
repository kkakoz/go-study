package test // 自定义支持类型
import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Timestamp struct {
	timestamp int64
}

func (t *Timestamp) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%d", t.timestamp)
	return buf.Bytes(), nil
}

// 实现sql.Scanner接口
func (t *Timestamp) Scan(val interface{}) (err error) {
	t.timestamp = val.(int64)
	return nil
}

type Point struct {
	X float64 `json:"lat"`
	Y float64 `json:"lng"`
}

// 实现driver.Valuer接口
func (p *Point) Value() (driver.Value, error) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "(%f %f)", p.X, p.Y)
	return buf.Bytes(), nil
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v %v)", p.X, p.Y)
}

// 实现sql.Scanner接口
func (p *Point) Scan(val interface{}) (err error) {
	if bb, ok := val.([]uint8); ok {
		tmp := bb[1 : len(bb)-1]
		coors := strings.Split(string(tmp[:]), ",")
		if p.X, err = strconv.ParseFloat(coors[0], 64); err != nil {
			return err
		}
		if p.Y, err = strconv.ParseFloat(coors[1], 64); err != nil {
			return err
		}
	}
	return nil
}

type Agent struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	CS_NO string `json:"cs_no"`
	Channel_id int `json:"channel_id"`
	Address string `json:"address"`
	Coordinate *Point `json:"point"`
}

func main() {
	pgurl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "postgres", "123456", "localhost", "5432", "people")
	db, err := sql.Open("postgres", pgurl)
	if err != nil {
		panic(fmt.Errorf("连接数据库出错:%v", err))
	}
	querySql := `SELECT * FROM t_agent`
	rows, err := db.Query(querySql)
	if err != nil {
		panic(fmt.Errorf("查询数据出错:%v", err))
	}
	for rows.Next() {
		agent := &Agent{Coordinate: &Point{}}
		err = rows.Scan(&agent.Id, &agent.Name, &agent.Code, &agent.CS_NO, &agent.Channel_id, &agent.Address, agent.Coordinate)
		fmt.Println(agent, err)
	}
	var id int
	err = db.QueryRow("INSERT INTO t_agent (name, code, cs_no, address, coordinate) VALUES($1,$2,$3,$4,$5) RETURNING id", "test1", "123457", "2", "111", "(12,43)").Scan(&id)
	fmt.Println("id:", id, "err:", err)
}
