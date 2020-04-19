package thailandpost

func convertGetItemsResponse(resp getItemsResponse, beYear bool) (TrackingStatusMap, TrackCount, error) {
	if !resp.Status {
		return nil, TrackCount{}, UnsuccessfulError{
			Message: resp.Message,
		}
	}

	trackCount, err := convertTrackCountResponse(resp.Response.TrackCount, beYear)
	if err != nil {
		return nil, TrackCount{}, err
	}

	statusMap := make(TrackingStatusMap)

	for barcode, statuses := range resp.Response.Items {
		for _, status := range statuses {
			s, err := convertTrackingStatusResponse(status, beYear)
			if err != nil {
				return nil, TrackCount{}, err
			}
			statusMap[barcode] = append(statusMap[barcode], s)
		}
	}

	return statusMap, trackCount, nil
}

func convertHookTrackResponse(resp hookTrackResponse, beYear bool) (HookStatusMap, TrackCount, error) {
	if !resp.Status {
		return nil, TrackCount{}, UnsuccessfulError{
			Message: resp.Message,
		}
	}

	trackCount, err := convertTrackCountResponse(resp.Response.TrackCount, beYear)
	if err != nil {
		return nil, TrackCount{}, err
	}

	statusMap := make(HookStatusMap)

	for _, status := range resp.Response.Items {
		statusMap[status.Barcode] = status.Status
	}

	return statusMap, trackCount, nil
}

func convertHookDataResponse(resp hookDataResponse, beYear bool) (HookData, error) {
	trackDt, err := parseWebhookTrackingTime(resp.TrackDatetime, beYear)
	if err != nil {
		return HookData{}, err
	}

	statusMap := make(TrackingStatusMap)

	for _, status := range resp.Items {
		s, err := convertTrackingStatusResponse(status, beYear)
		if err != nil {
			return HookData{}, err
		}
		barcode := s.Barcode
		statusMap[barcode] = append(statusMap[barcode], s)
	}

	return HookData{
		Items:         statusMap,
		TrackDatetime: trackDt,
	}, nil
}

func convertTrackCountResponse(resp trackCountResponse, beYear bool) (TrackCount, error) {
	d, err := parseTrackingDate(resp.TrackDate, beYear)
	if err != nil {
		return TrackCount{}, err
	}

	return TrackCount{
		TrackDate:       d,
		CountNumber:     resp.CountNumber,
		TrackCountLimit: resp.TrackCountLimit,
	}, nil
}

func convertTrackingStatusResponse(resp trackingStatusResponse, beYear bool) (TrackingStatus, error) {
	t, err := parseTrackingTime(resp.StatusDate, beYear)
	if err != nil {
		return TrackingStatus{}, err
	}
	dt, err := parseTrackingTime(resp.DeliveryDateTime, beYear)
	if err != nil {
		return TrackingStatus{}, err
	}

	return TrackingStatus{
		Barcode:             resp.Barcode,
		Status:              ItemStatus(resp.Status),
		StatusDescription:   resp.StatusDescription,
		StatusDate:          t,
		Location:            resp.Location,
		Postcode:            resp.Postcode,
		DeliveryStatus:      ItemDeliveryStatus(resp.DeliveryStatus),
		DeliveryDescription: resp.DeliveryDescription,
		DeliveryDateTime:    dt,
		ReceiverName:        resp.ReceiverName,
		Signature:           resp.Signature,
	}, nil
}
