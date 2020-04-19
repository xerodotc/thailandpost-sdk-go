package thailandpost

import "time"

var timeFunc = time.Now

const tokenExpireTimeFormat = "2006-01-02 15:04:05-07:00"
const trackingTimeFormat = "02/01/2006 15:04:05-07:00"
const trackingDateFormat = "02/01/2006"
const webhookTimeFormat = "02/01/2006 15:04-07:00"

func parseTokenExpireTime(expire string) (time.Time, error) {
	return time.Parse(tokenExpireTimeFormat, expire)
}

func parseTrackingTime(timeString string, beYear bool) (time.Time, error) {
	if timeString == "" {
		return time.Time{}, nil
	}

	t, err := time.Parse(trackingTimeFormat, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return beYearToCEYearWrap(t, beYear), nil
}

func parseTrackingDate(dateString string, beYear bool) (time.Time, error) {
	t, err := time.Parse(trackingDateFormat, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return beYearToCEYearWrap(t, beYear), nil
}

func parseWebhookTrackingTime(timeString string, beYear bool) (time.Time, error) {
	t, err := time.Parse(webhookTimeFormat, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return beYearToCEYearWrap(t, beYear), nil
}

func beYearToCEYear(t time.Time) time.Time {
	return t.AddDate(-543, 0, 0)
}

func beYearToCEYearWrap(t time.Time, beYear bool) time.Time {
	if beYear {
		return beYearToCEYear(t)
	}
	return t
}
