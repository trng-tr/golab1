package maps

import (
	"fmt"
	"golab1/utils"
)

func initialiserMap() (map[string]int, map[int]string, map[string]User) {
	var map1 map[string]int = make(map[string]int)
	var map2 map[int]string = make(map[int]string)
	var map3 map[string]User = make(map[string]User)
	return map1, map2, map3
}
func RemplirMaps() {
	map1, _, map3 := initialiserMap()
	v1, _ := utils.SaisirNumber()
	v2, _ := utils.SaisirNumber()
	map1["key1"], map1["key2"] = int(v1), int(v2) //affectation sur une meme ligne
	//création de l' objet address
	var address Address = Address{
		streetNumber: 184,
		streetName:   "Avenue de Liège",
		postalCode:   59300,
		city:         "Valanciennes",
		country:      "France",
	}
	var user1 User = User{
		firstname: "Placide", lastname: "Nduwayo", email: "placide.nd@domain.com", address: address,
	}
	var user2 User = User{
		firstname: "Belyse", lastname: "Ndiwayo", email: "arakizabelyse1@gmail.com", address: address,
	}
	map3["person1"], map3["person2"] = user1, user2 //affectation sur une meme ligne
	for key1, value1 := range map1 {
		fmt.Printf("(key: %v,value:%v)\n", key1, value1)
	}

	for key2, value2 := range map3 {
		fmt.Printf("(key: %v,value:%v)\n", key2, value2)
	}
}
