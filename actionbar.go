package modiniter

var ActionBar ActionBarItemsList

type ActionBarItem struct {
	Label  string
	Grants []string
}

type ActionBarItemsList []ActionBarItem
