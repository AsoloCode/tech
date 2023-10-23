package model

import (
	"tech/mypackage/pb"
	"time"
)

type Order struct {
	OrderUID        string    `json:"order_uid"`
	TrackNumber     string    `json:"track_number"`
	Entry           string    `json:"entry"`
	Locale          string    `json:"locale"`
	InternalSig     string    `json:"internal_signature"`
	CustomerID      string    `json:"customer_id"`
	DeliveryService string    `json:"delivery_service"`
	ShardKey        string    `json:"shardkey"`
	SMID            int       `json:"sm_id"`
	OofShard        string    `json:"oof_shard"`
	DateCreated     time.Time `json:"date_created"`
	Delivery        Delivery  `json:"delivery"`
	Payment         Payment   `json:"payment"`
	Items           []Item    `json:"items"`
}

func (p *Order) OrderFromProto(pbOrder *pb.Order) {
	p.TrackNumber = pbOrder.TrackNumber
	p.Entry = pbOrder.Entry
	p.Locale = pbOrder.Locale
	p.InternalSig = pbOrder.InternalSig
	p.CustomerID = pbOrder.CustomerID
	p.DeliveryService = pbOrder.DeliveryService
	p.ShardKey = pbOrder.ShardKey
	p.SMID = int(pbOrder.SMID)
	p.OofShard = pbOrder.OofShard
	p.DateCreated = time.Unix(pbOrder.DateCreated.GetSeconds(), int64(pbOrder.DateCreated.GetNanos()))

	p.Delivery = Delivery{
		Name:    pbOrder.Delivery.Name,
		Phone:   pbOrder.Delivery.Phone,
		Zip:     pbOrder.Delivery.Zip,
		City:    pbOrder.Delivery.City,
		Address: pbOrder.Delivery.Address,
		Region:  pbOrder.Delivery.Region,
		Email:   pbOrder.Delivery.Email,
	}

	p.Payment = Payment{
		Transaction:  pbOrder.Payment.Transaction,
		RequestID:    pbOrder.Payment.RequestId,
		Currency:     pbOrder.Payment.Currency,
		Provider:     pbOrder.Payment.Provider,
		Amount:       int(pbOrder.Payment.Amount),
		PaymentDT:    time.Unix(pbOrder.Payment.PaymentDt.Seconds, int64(pbOrder.Payment.PaymentDt.Nanos)),
		Bank:         pbOrder.Payment.Bank,
		DeliveryCost: int(pbOrder.Payment.DeliveryCost),
		GoodsTotal:   int(pbOrder.Payment.GoodsTotal),
		CustomFee:    int(pbOrder.Payment.CustomFee),
	}

	p.Items = make([]Item, len(pbOrder.Items))
	for i, pbItem := range pbOrder.Items {
		p.Items[i] = Item{
			ChrtID:      int(pbItem.ChrtId),
			TrackNumber: pbItem.TrackNumber,
			Price:       float64(pbItem.Price),
			RID:         pbItem.Rid,
			Name:        pbItem.Name,
			Sale:        float64(pbItem.Sale),
			Size:        pbItem.Size,
			TotalPrice:  float64(pbItem.TotalPrice),
			NMID:        int(pbItem.NmId),
			Brand:       pbItem.Brand,
			Status:      int(pbItem.Status),
		}
	}
}
