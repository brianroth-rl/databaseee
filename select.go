package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"fmt"
	"errors"
)

type Event struct {
	ID                 int64          `gorm:"primarykey" json:"id" form:"id"`
	CreatedAt          time.Time      `gorm:"column:created_at;"json:"created_at" form:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;"json:"updated_at" form:"updated_at"`
	ArchivedAt         gorm.DeletedAt `gorm:"column:archived_at;"gorm:"index" json:"archived_at" form:"archived_at"`
	EventableType      string
	EventableId        int64
	OccurrenceTime     time.Time
	ExternalId         string
	CaptureVisitorId   string
	MasterCampaignId   int64
	PlatformId         int64
	MasterAdvertiserId int64
	WpcId              string
	ReferrerType       string
	ReferrerSource     string
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "lips_user:reabIefT@tcp(127.0.0.1:3306)/lips_dev"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// // Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var event Event
	var events []Event

	db.First(&event, 18543719) // find product with integer primary key

	// Get first matched record
	db.Where("master_campaign_id = ?", 1679630).First(&event)

	// Get all matched records
	result := db.Where("last_name = ?", "Davis").Find(&events)
	fmt.Printf("rows = %d\n", result.RowsAffected)
	fmt.Printf("events = %d\n", len(events))
	fmt.Printf("error = %t\n", errors.Is(result.Error, gorm.ErrRecordNotFound))

	// IN
	db.Where("last_name IN ?", []string{"Davis", "Shah"}).Find(&event)
	fmt.Printf("rows = %d\n", result.RowsAffected)
	fmt.Printf("events = %d\n", len(events))
	fmt.Printf("error = %t\n", errors.Is(result.Error, gorm.ErrRecordNotFound))

	// Delete - delete product
	db.Delete(&event, 18543719)
}
