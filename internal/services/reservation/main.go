package reservation

import (
	"fmt"
)

func main() {
	hdb, err := NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	var rdb *storage
	rdb, err = NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	srv := NewServer(rdb, hdb)
	err = srv.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}
