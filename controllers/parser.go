package controllers

import (
	"fmt"
	"os"

	"github.com/peterducai/invcomb/models"
)

//WriteSingleInventoryFile writes file
func WriteSingleInventoryFile(path string, single bool) {

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	if _, err = fmt.Fprintf(f, "%+v", models.Inv); err != nil {
		panic(err)
	}

	f.Sync()
	return
}
