package main

import (
	"time"

	"github.com/Djoulzy/DataTest/actions"
	"github.com/Djoulzy/DataTest/datamodel"
	"github.com/Djoulzy/Tools/clog"
)

func main() {
	clog.LogLevel = 5
	clog.StartLogging = true

	data := datamodel.RichItem{
		BaseItem: datamodel.BaseItem{
			ID:        2017,
			Crea_date: time.Now(),
			Views:     2,
		},
		Height: 12,
		Width:  24,
		Tags:   "ok",
	}

	actions.Process(data)

	otherdata := datamodel.AnotherItem{
		Name:    "Julien",
		Surname: "Marusi",
		Age:     45,
	}
	actions.Process(otherdata)
}
