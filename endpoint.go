package thailandpost

const (
	APIBase     = "https://trackapi.thailandpost.co.th"
	WebhookBase = "https://trackwebhook.thailandpost.co.th"

	APIGetTokenURL     = APIBase + "/post/api/v1/authenticate/token"
	APIGetItemsURL     = APIBase + "/post/api/v1/track"
	APIRequestItemsURL = APIBase + "/post/api/v1/track/batch"

	WebhookGetTokenURL  = WebhookBase + "/post/api/v1/authenticate/token"
	WebhookHookTrackURL = WebhookBase + "/post/api/v1/hook"
)
