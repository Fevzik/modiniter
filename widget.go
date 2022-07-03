package modiniter

var Widgets WidgetsList

type WidgetItem struct {
	Label  string
	Grants []string
}

type WidgetsList []WidgetItem
