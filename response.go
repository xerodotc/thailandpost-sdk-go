package thailandpost

type tokenResponse struct {
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

type getItemsResponse struct {
	Response struct {
		Items      map[string][]trackingStatusResponse `json:"items"`
		TrackCount trackCountResponse                  `json:"track_count"`
	} `json:"response"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type hookTrackResponse struct {
	Response struct {
		Items []struct {
			Barcode string `json:"barcode"`
			Status  bool   `json:"status"`
		} `json:"items"`
		TrackCount trackCountResponse `json:"track_count"`
	} `json:"response"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type hookDataResponse struct {
	Items         []trackingStatusResponse `json:"items"`
	TrackDatetime string                   `json:"track_datetime"`
}

type trackCountResponse struct {
	TrackDate       string `json:"track_date"`
	CountNumber     int    `json:"count_number"`
	TrackCountLimit int    `json:"track_count_limit"`
}

type trackingStatusResponse struct {
	Barcode             string `json:"barcode"`
	Status              string `json:"status"`
	StatusDescription   string `json:"status_description"`
	StatusDate          string `json:"status_date"`
	Location            string `json:"location"`
	Postcode            string `json:"postcode"`
	DeliveryStatus      string `json:"delivery_status"`
	DeliveryDescription string `json:"delivery_description"`
	DeliveryDateTime    string `json:"delivery_datetime"`
	ReceiverName        string `json:"receiver_name"`
	Signature           string `json:"signature"`
}
