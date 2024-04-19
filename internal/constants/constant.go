package constants

const DatetimeFormat = "2006-01-02 15:04:05"

type TrxStatus string

const (
	Unpaid TrxStatus = "Unpaid"
	Paid   TrxStatus = "Paid"
	Cancel TrxStatus = "Cancel"
	Failed TrxStatus = "Failed"
)
