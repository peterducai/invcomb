package controllers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/peterducai/invcomb/models"
)

//PrettyPrint prints nicely
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

//WriteSingleInventoryFile writes file
func WriteSingleInventoryFile(path string, single bool) {

	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	if _, err = fmt.Fprintf(f, "%s\n", PrettyPrint(models.Inv)); err != nil {
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
