package payment

import (
	"fmt"
)

func main() {
	db, err := NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	srv := NewServer(db)
	err = srv.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}
