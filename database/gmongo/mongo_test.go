package gmongo

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/shawnwyckoff/gpkg/sys/gclock"
	"time"

	//"github.com/shawnwyckoff/finance"
	"testing"
)

/*
func TestNewConn(t *testing.T) {
	conn, err := Dial("mongodb://192.168.9.11:27717")
	if err != nil {
		t.Error(err)
		return
	}

	now := time.Now()
	kl := finance.Kline{}
	kl.Time = now
	if err != nil {
		t.Error(err)
		return
	}
	kl.High = 500
	kl.Low = 1.00000123
	err = conn.DB("coins").C("kline").Insert(kl)
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second)
	kl.Low = 1.000001234567
	err = conn.DB("coins").C("kline").Upsert(bson.M{"Time":kl.Time}, bson.M{"$set": kl})
	if err != nil {
		t.Error(err)
		return
	}

	kl.Time = kl.Time.Add(time.Second)
	err = conn.DB("coins").C("kline").Upsert(bson.M{"Time":kl.Time}, bson.M{"$set": kl})
	if err != nil {
		t.Error(err)
		return
	}

	maxTime, _, err := conn.DB("coins").C("kline").MaxTime("Time")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(maxTime)
}
*/
func TestColl_Find(t *testing.T) {
	conn, err := Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}

	sz, err := conn.DB("Details2").C("123").Count(nil)
	fmt.Println(sz, err)

	cur, err := conn.DB("Details").C("123").Find(nil)
	fmt.Println(cur.Next(), err)
}

func TestColl_UpsertEntireDoc_UpsertFields(t *testing.T) {
	conn, err := Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}

	type TestItemA struct {
		Name  string `bson:"_id"`
		Age   int    `bson:"Age"`
		Class string `bson:"Class"`
	}

	type TestItemB struct {
		Name string `bson:"_id"`
		Age  int    `bson:"Age"`
		Step string `bson:"Step"`
	}

	type TestItemC struct {
		Age  int    `bson:"Age"`
		Step string `bson:"Step"`
	}

	itemA := TestItemA{Name: "WangJiaChen", Age: 18, Class: "#1"}
	itemRetA := TestItemA{}
	if err = conn.DB("test").C("test").UpsertEntireDoc(itemA.Name, itemA); err != nil {
		t.Error(err)
		return
	}
	ok, err := conn.DB("test").C("test").FindId(itemA.Name, &itemRetA)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Errorf("must FindId")
		return
	}
	fmt.Println(itemRetA)
	if itemRetA != itemA {
		t.Error("UpsertEntireDoc error")
		return
	}

	itemB := TestItemB{Name: "WangJiaChen", Age: 19, Step: "#2"}
	itemRetB := TestItemB{}
	if err = conn.DB("test").C("test").UpsertEntireDoc(itemB.Name, itemB); err != nil {
		t.Error(err)
		return
	}
	ok, err = conn.DB("test").C("test").FindId(itemB.Name, &itemRetB)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Errorf("must FindId")
		return
	}
	fmt.Println(itemRetB)
	if itemRetB != itemB {
		t.Error("UpsertEntireDoc error")
		return
	}

	itemC := TestItemC{Age: 20, Step: "#3"}
	itemRetC := TestItemB{}
	if err = conn.DB("test").C("test").UpsertFields(itemRetB.Name, itemC); err != nil {
		t.Error(err)
		return
	}
	ok, err = conn.DB("test").C("test").FindId(itemB.Name, &itemRetC)
	if err != nil {
		t.Error(err)
		return
	}
	if !ok {
		t.Errorf("must FindId")
		return
	}
	fmt.Println(itemRetC)
	if itemRetC.Name != "WangJiaChen" || itemRetC.Age != 20 || itemRetC.Step != "#3" {
		t.Error("UpsertEntireDoc error")
		return
	}
}

type tmpKline struct {
	Time      time.Time `json:"Time" bson:"_id"`
	Open      float64   `json:"Open" bson:"Open"` // open price in USD
	Close     float64   `json:"Close" bson:"Close"`
	High      float64   `json:"High" bson:"High"`
	Low       float64   `json:"Low" bson:"Low"`
	Volume    float64   `json:"Volume" bson:"Volume"`                           // volume in USD
	MarketCap float64   `json:"MarketCap,omitempty" bson:"MarketCap,omitempty"` // exchangeName cap in USD CoinMarketCap接口中用到
}

func TestColl_FindCmp(t *testing.T) {
	conn, err := Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}

	date, err := gclock.NewDate(2018, 11, 1)
	if err != nil {
		t.Error(err)
		return
	}

	cur, err := conn.DB("TradeHistory-YF").C("s.nasdaq.aapl").FindCmp("_id", CmpGTE, bsonx.Time(date.ToTimeUTC()))
	item := tmpKline{}
	total := 0
	for cur.Next() {
		if err := cur.Decode(&item); err != nil {
			t.Error(err)
			return
		} else {
			t.Log(item)
			total++
		}
	}
	t.Log(total)
}

func TestColl_RemoveCmp(t *testing.T) {
	conn, err := Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		t.Error(err)
		return
	}

	date, err := gclock.NewDate(2019, 1, 1)
	if err != nil {
		t.Error(err)
		return
	}

	ds, err := conn.DB("TradeHistory-CMC").C("c.0chain").RemoveCmp("_id", CmpGTE, bsonx.Time(date.ToTimeUTC()))
	t.Log(ds, err)
	t.Log(conn.DB("TradeHistory-CMC").C("c.0chain").MaxTime("_id"))
}
