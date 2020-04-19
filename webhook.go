package thailandpost

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TrackingWebhook interface {
	HookTrack(items ...string) (HookStatusMap, error)
	HookTrackWithStatus(status ItemStatus, items ...string) (HookStatusMap, error)

	ParseHookData(req *http.Request) (HookData, error)

	GetLastTrackCount() (TrackCount, error)
}

type trackingWebhookImpl struct {
	*clientMiddleware
	trackCount    *TrackCount
	lang          Lang
	myBearerToken string
}

func TrackingWebhookInit(lang Lang, token string, myBearerToken string) TrackingWebhook {
	return TrackingWebhookInitWithClient(lang, token, myBearerToken, http.DefaultClient)
}

func TrackingWebhookInitWithClient(lang Lang, token string, myBearerToken string, client *http.Client) TrackingWebhook {
	return &trackingWebhookImpl{
		clientMiddleware: &clientMiddleware{
			RefreshToken: token,
			AuthURL:      WebhookGetTokenURL,
			Client:       client,
		},
		lang:          lang,
		myBearerToken: myBearerToken,
	}
}

func (api *trackingWebhookImpl) HookTrack(items ...string) (HookStatusMap, error) {
	return api.HookTrackWithStatus(StatusAll, items...)
}

func (api *trackingWebhookImpl) HookTrackWithStatus(status ItemStatus, items ...string) (HookStatusMap, error) {
	if err := ValidateTrackingNumbers(items); err != nil {
		return nil, err
	}

	req := trackingRequest{
		Status:   string(status),
		Language: string(api.lang),
		Barcode:  items,
	}
	beYear := isLanguageNeedBEConversion(api.lang)

	var resp hookTrackResponse
	if err := api.doJSONPostRequest(WebhookHookTrackURL, req, &resp); err != nil {
		return nil, err
	}

	statuses, tc, err := convertHookTrackResponse(resp, beYear)
	if err != nil {
		return nil, err
	}

	api.trackCount = &tc

	return statuses, nil
}

func (api *trackingWebhookImpl) ParseHookData(req *http.Request) (HookData, error) {
	if req.Header.Get("Authorization") != "Bearer "+api.myBearerToken {
		return HookData{}, ErrInvalidBearerToken
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return HookData{}, err
	}
	defer func(body []byte) {
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
	}(body)

	var resp hookDataResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return HookData{}, err
	}

	data, err := convertHookDataResponse(resp, isLanguageNeedBEConversion(api.lang))
	if err != nil {
		return HookData{}, err
	}

	return data, nil
}

func (api *trackingWebhookImpl) GetLastTrackCount() (TrackCount, error) {
	if api.trackCount == nil {
		return TrackCount{}, ErrNoLastRequest
	}
	return *api.trackCount, nil
}
