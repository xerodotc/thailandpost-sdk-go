package thailandpost

import "time"

// Tracking status for item
type TrackingStatus struct {
	Barcode             string             `json:"barcode"`
	Status              ItemStatus         `json:"status"`
	StatusDescription   string             `json:"status_description"`
	StatusDate          time.Time          `json:"status_date"`
	Location            string             `json:"location"`
	Postcode            string             `json:"postcode"`
	DeliveryStatus      ItemDeliveryStatus `json:"delivery_status"`
	DeliveryDescription string             `json:"delivery_description"`
	DeliveryDateTime    time.Time          `json:"delivery_datetime"`
	ReceiverName        string             `json:"receiver_name"`
	Signature           string             `json:"signature"`
}

// Map from barcode to historical tracking statuses
type TrackingStatusMap map[string][]TrackingStatus

// Hook track result
type HookStatusMap map[string]bool

// Hook data
type HookData struct {
	Items         TrackingStatusMap `json:"items"`
	TrackDatetime time.Time         `json:"track_datetime"`
}

// API track count
type TrackCount struct {
	TrackDate       time.Time `json:"track_date"`
	CountNumber     int       `json:"count_number"`
	TrackCountLimit int       `json:"track_count_limit"`
}
