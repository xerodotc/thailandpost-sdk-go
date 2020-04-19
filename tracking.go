package thailandpost

import "net/http"

type TrackingAPI interface {
	// Get items' tracking status (less than 100 items)
	GetItems(items ...string) (TrackingStatusMap, error)
	// Get items' tracking status (less than 100 items) with specific status
	GetItemsWithStatus(status ItemStatus, items ...string) (TrackingStatusMap, error)
	// Request items' tracking status via e-mail
	RequestBatchItems(items []string) error
	// Request items' tracking status via e-mail with specific status
	RequestBatchItemsWithStatus(status ItemStatus, items []string) error

	// Get last API call track count
	GetLastTrackCount() (TrackCount, error)
}

type trackingAPIImpl struct {
	*clientMiddleware
	trackCount *TrackCount
	lang       Lang
}

// Initialize tracking API
func TrackingAPIInit(lang Lang, token string) TrackingAPI {
	return TrackingAPIInitWithClient(lang, token, http.DefaultClient)
}

// Initialize tracking API with custom http.Client
func TrackingAPIInitWithClient(lang Lang, token string, client *http.Client) TrackingAPI {
	return &trackingAPIImpl{
		clientMiddleware: &clientMiddleware{
			RefreshToken: token,
			AuthURL:      APIGetTokenURL,
			Client:       client,
		},
		lang: lang,
	}
}

func (api *trackingAPIImpl) GetItems(items ...string) (TrackingStatusMap, error) {
	return api.GetItemsWithStatus(StatusAll, items...)
}

func (api *trackingAPIImpl) GetItemsWithStatus(status ItemStatus, items ...string) (TrackingStatusMap, error) {
	if len(items) > GetItemsLimit {
		return nil, ErrTooMuchTrackings
	}

	if err := ValidateTrackingNumbers(items); err != nil {
		return nil, err
	}

	req := trackingRequest{
		Status:   string(status),
		Language: string(api.lang),
		Barcode:  items,
	}
	beYear := isLanguageNeedBEConversion(api.lang)

	var resp getItemsResponse
	if err := api.doJSONPostRequest(APIGetItemsURL, req, &resp); err != nil {
		return nil, err
	}

	statuses, tc, err := convertGetItemsResponse(resp, beYear)
	if err != nil {
		return nil, err
	}

	api.trackCount = &tc

	return statuses, nil
}

func (api *trackingAPIImpl) RequestBatchItems(items []string) error {
	return api.RequestBatchItemsWithStatus(StatusAll, items)
}

func (api *trackingAPIImpl) RequestBatchItemsWithStatus(status ItemStatus, items []string) error {
	if err := ValidateTrackingNumbers(items); err != nil {
		return err
	}

	req := trackingRequest{
		Status:   string(status),
		Language: string(api.lang),
		Barcode:  items,
	}
	beYear := isLanguageNeedBEConversion(api.lang)

	var resp getItemsResponse
	if err := api.doJSONPostRequest(APIRequestItemsURL, req, &resp); err != nil {
		return err
	}

	_, tc, err := convertGetItemsResponse(resp, beYear)
	if err != nil {
		return err
	}

	api.trackCount = &tc

	return nil
}

func (api *trackingAPIImpl) GetLastTrackCount() (TrackCount, error) {
	if api.trackCount == nil {
		return TrackCount{}, ErrNoLastRequest
	}
	return *api.trackCount, nil
}
