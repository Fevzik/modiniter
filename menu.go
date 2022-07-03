package modiniter

var Menu MenuItemList

type MenuItem struct {
	Label  string
	Grants []string
}

type MenuItemList []MenuItem
