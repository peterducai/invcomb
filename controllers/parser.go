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

//WriteDirs create directories group_vars and host_vars
func WriteDirs() {
	_ = os.Mkdir("group_vars", 0700)
	_ = os.Mkdir("host_vars", 0700)
}
