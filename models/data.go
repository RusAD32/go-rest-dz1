package models

var DB = make(map[int]Item)

type Item struct {
	Title  string  `json:"title"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

func init() {
	initItem := Item{
		Title: "Lightbulb",
		Amount: 100,
		Price: 1.0,
	}
	AddItemToDB(1, initItem)
}

func FindItemById(id int) (Item, bool) {
	book, ok := DB[id]
	return book, ok
}

func AddItemToDB(id int, item Item) bool {
	if _, ok := DB[id]; ok {
		return false
	}
	DB[id] = item
	return true
}

func DeleteItemFromDB(id int) bool {
	if _, ok := DB[id]; !ok {
		return false
	}
	delete(DB, id)
	return true
}

func UpdateItemInDB(id int, item Item) bool {
	if _, ok := DB[id]; !ok {
		return false
	}
	DB[id] = item
	return true
}

func IsDBEmpty() bool {
	return len(DB) == 0
}
