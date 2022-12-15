package seed

import (
"fmt"
"net/url"
handler "JSE_API/pkg/configs/config_database"
"log"
)


func seed() {
	fmt.Println("hello")
	u, err := url.Parse("https://www.jamstockex.com/trading/trade-summary/")
	if err != nil {
		log.Println(err)
	}
	err = handler.Storage.Put(u)
	if err != nil {
		log.Println(err)
	}

}
