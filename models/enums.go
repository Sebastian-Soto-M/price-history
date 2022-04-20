package models

type StoreEnum string

const (
	Amazon StoreEnum = "amazon"
	Wish   StoreEnum = "wish"
	Ebay   StoreEnum = "ebay"
)

func StoreEnumList() []string {
	return []string{string(Amazon), string(Wish), string(Ebay)}
}

func (se StoreEnum) Store() (store *Store) {
	switch se {
	case Amazon:
		store = &Store{Name: "Amazon", Domain: "www.amazon.com"}
	case Ebay:
		store = &Store{Name: "Ebay", Domain: "www.amazon.com"}
	case Wish:
		store = &Store{Name: "Wish", Domain: "www.wish.com"}
	}
	return store
}
