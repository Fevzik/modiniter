package modiniter

var Reports ReportsList

type ReportItem struct {
	Label  string
	Grants []string
}

type ReportsList []ReportItem
