package streaming

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"go_nats-streaming_pg/internal/config"
	"go_nats-streaming_pg/internal/db"
	"log"
)

func Publish(nats config.Stan) {
	log.Println("Start publishing")
	sc, _ := stan.Connect(nats.ClusterID, nats.ClientID, stan.NatsURL(nats.SvrURL))
	defer sc.Close()
	item1 := db.Items{
		ChrtId:      1,
		TrackNumber: "abcd",
		Price:       20,
		RId:         "dasdasd",
		Name:        "Name",
		Sale:        10,
		Size:        "XL",
		TotalPrice:  20,
		NmId:        302010,
		Brand:       "gusi",
		Status:      1,
	}
	item2 := db.Items{
		ChrtId:      2,
		TrackNumber: "abcddasdasd",
		Price:       31210,
		RId:         "dasdasd",
		Name:        "Name",
		Sale:        112310,
		Size:        "L",
		TotalPrice:  20,
		NmId:        302010,
		Brand:       "gusi",
		Status:      1,
	}
	delivery := db.Delivery{
		Name:    "Name",
		Phone:   "mobile",
		ZIP:     "3464",
		City:    "NVCH",
		Address: "KK",
		Region:  "RR",
		Email:   "@",
	}
	payment := db.Payment{
		Transaction:  "dd",
		RequestId:    "das",
		Currency:     "dsa",
		Provider:     "dd",
		Amount:       10,
		PaymentDt:    11,
		Bank:         "alf",
		DeliveryCost: 12,
		GoodsTotal:   23,
		CustomFee:    32,
	}
	order := db.Order{
		OrderUID:          "Order 2",
		TrackNumber:       "track 2",
		Entry:             "adsdas",
		Delivery:          delivery,
		Payment:           payment,
		Items:             []db.Items{item1, item2},
		Locale:            "das",
		InternalSignature: "dddddd",
		CustomerId:        "customer 2",
		DeliveryService:   "DS2",
		Shardkey:          "adsasd",
		SmId:              123123,
		DateCreated:       "12.12.2013",
		CofShard:          "asdasd",
	}

	subject := nats.Subject
	for n := 0; n < 1; n++ {
		msg, _ := json.Marshal(order)
		sc.Publish(subject, msg)
		log.Println("Publish successfully")
	}
}
