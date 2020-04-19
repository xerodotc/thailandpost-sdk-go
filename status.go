package thailandpost

// Item status
type ItemStatus string

const (
	StatusPreload         ItemStatus = "101"
	StatusAcceptedByAgent ItemStatus = "102"
	StatusCollected       ItemStatus = "103"

	StatusInTransit               ItemStatus = "201"
	StatusCustomClearance         ItemStatus = "202"
	StatusReturnToSender          ItemStatus = "203"
	StatusArriveAtOutwardExchange ItemStatus = "204"
	StatusArriveAtInwardExchange  ItemStatus = "205"
	StatusArriveAtPostOffice      ItemStatus = "206"
	StatusPrepareTransit          ItemStatus = "207"

	StatusOutForDelivery  ItemStatus = "301"
	StatusArriveForPickUp ItemStatus = "302"

	StatusDeliveryFailed ItemStatus = "401"

	StatusDeliverySuccess ItemStatus = "501"

	StatusAll ItemStatus = "all"
)

type ItemDeliveryStatus string

// Delivery status
// this one is undocumented, need to be reversed lookup
const (
	DeliveryStatusSuccess ItemDeliveryStatus = "S"

	DeliveryStatusWaitAtPostOffice ItemDeliveryStatus = "F"

	DeliveryStatusLocationClosed  ItemDeliveryStatus = "11C"
	DeliveryStatusPaymentRequired ItemDeliveryStatus = "21C"
)
