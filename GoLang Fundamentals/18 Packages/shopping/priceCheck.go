package shopping

import (
	"shopping/db"
)

type Item struct{
	Price int
}


func Pricecheck(itemId int) (int, bool){
	item := db.LoadItem(itemId);
	if item == nil{
		return 0, false
	}

	return item.Price, true
}