package internal

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Status string

const (
	Live          Status = "Live"
	Finished      Status = "Finished"
	ReserveNotMet Status = "ReserveNotMet"
)

type Item struct {
	Base
	Make      string
	Model     string
	Year      int
	Color     string
	Mileage   int
	ImageURL  string
	AuctionID uuid.UUID
	Auction   Auction
}

type Auction struct {
	Base
	ReservePrice   int
	Seller         string
	Winner         string
	SoldAmount     *int
	CurrentHighBid *int
	AuctionEnd     time.Time
	Status         Status `gorm:"type:varchar(20)"`
}

func (a *Auction) HasReservePrice() bool {
	return a.ReservePrice > 0
}
