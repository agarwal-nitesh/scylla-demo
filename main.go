package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"time"
)

type Booking struct {
	Id int64 `db:"id"`
	Amount int64 `db:"amount"`
	BookingTime int64 `db:"bookingtime"`
	Status int64 `db:"status"`
	UserId int64 `db:"userid"`
}


func main() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.Compressor = &gocql.SnappyCompressor{}
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{}
	cluster.Consistency = gocql.LocalQuorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("scylla init done")

	// -------BASIC OPERATIONS--------
	// SELECT
	var booking Booking
	stmt, names := qb.Select("test_dev.booking").Where(qb.Eq("id")).ToCql()
	q := gocqlx.Query(session.Query(stmt).Consistency(gocql.One), names).BindMap(qb.M{
		"id": 1,
	})
	if err := q.GetRelease(&booking); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(booking)
	}

	// INSERT
	iq := "INSERT INTO test_dev.booking(id, amount, bookingtime, status, userid) VALUES(?, ?, ?, ?, ?)"
	fmt.Println(iq)
	if err := session.Query(iq, 11, 2000, time.Now(), 10, 100).Consistency(gocql.One).Exec(); err != nil {
		fmt.Println(err)
	}

	// UPDATE
	uq := "INSERT INTO test_dev.booking(id, amount, bookingtime, status, userid) VALUES(?, ?, ?, ?, ?)"
	fmt.Println(uq)
	if err := session.Query(uq, 11, 4000, time.Now(), 10, 100).Consistency(gocql.One).Exec(); err != nil {
		fmt.Println(err)
	}


	//// DELETE
	//dq := "DELETE FROM test_dev.booking WHERE id = ?"
	//fmt.Println(dq)
	//if err := session.Query(dq, 11).Consistency(gocql.One).Exec(); err != nil {
	//	fmt.Println(err)
	//}
	// -------BASIC OPERATIONS END-----


}


