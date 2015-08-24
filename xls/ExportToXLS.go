package main

import (
	"encoding/csv"
	"os"
)

func main() {

	f, err := os.Create("1.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(f)
	w.Write([]string{"Number", "Name", "Age"})
	w.Write([]string{"1", "Lucy", "23"})
	w.Write([]string{"2", "Hanmeimei", "23"})
	w.Write([]string{"3", "Lilei", "23"})
	w.Write([]string{"4", "Lili", "23"})
	w.Flush()

}
