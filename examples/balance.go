package main

import (
	"database/sql"
	"flag"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/giangcoy/go-smartbatching"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbAddr       *string = flag.String("a", "user:password@/dbname", "mysql address")
	nThread      *int    = flag.Int("n", 100, "number of thread")
	nRequest     *int    = flag.Int("r", 10, "number of request per thread")
	tblName      string  = "balance"
	keyField     string  = "acc"
	balanceField string  = "amount"
)

type balanceHotspot struct {
	db *sql.DB
}

func (b *balanceHotspot) Do(key string, datas []interface{}) []interface{} {
	tx, err := b.db.Begin()
	var balance int64 = 0
	fail := func(err error) []interface{} {
		if err != nil {
			panic(err)
		}
		return datas
	}
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	if err = tx.QueryRow(fmt.Sprintf("SELECT %s from %s where %s=? for update", balanceField, tblName, keyField), key).Scan(&balance); err != nil {
		return fail(err)
	}

	for i, data := range datas {
		amount := data.(int64)
		if amount <= balance {
			datas[i] = true
			balance -= amount
		} else {
			datas[i] = false
		}

	}
	if _, err = tx.Exec(fmt.Sprintf("Update %s  SET %s=? WHERE %s=?", tblName, balanceField, keyField), balance, key); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return datas
}
func main() {
	flag.Parse()
	keyTest := "A"
	db, err := sql.Open("mysql", *dbAddr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(%s VARCHAR(256) PRIMARY KEY, %s BIGINT(20))", tblName, keyField, balanceField)
	if _, err = db.Exec(query); err != nil {
		panic(err)
	}
	query = fmt.Sprintf("INSERT IGNORE INTO %s VALUES('%s',1000000)", tblName, keyTest)
	if _, err = db.Exec(query); err != nil {
		fmt.Printf("Error %v\n", err)
		panic(err)
	}
	n := *nRequest * *nThread
	durations := make([]int, 0, n)
	muDurations := sync.Mutex{}
	batch := smartbatching.NewSmartBatching()
	balance := &balanceHotspot{db}

	wg := sync.WaitGroup{}
	t0 := time.Now()
	for t := 0; t < *nThread; t++ {
		wg.Add(1)
		go func() {
			for r := 0; r < *nRequest; r++ {
				t1 := time.Now()
				batch.Add(balance, keyTest, int64(1))
				muDurations.Lock()
				durations = append(durations, int(time.Since(t1).Milliseconds()))
				muDurations.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	sort.Ints(durations)
	k := int(time.Since(t0).Seconds())
	if k < 1 {
		k = 1
	}
	fmt.Printf("TPS: %d\n", n/k)
	percentiles := []int{99, 95, 90, 75, 50}
	for _, percentile := range percentiles {
		fmt.Printf("P%d: %d(ms)\n", percentile, durations[n*percentile/100])
	}

}
