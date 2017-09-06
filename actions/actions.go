package actions

import (
	"reflect"

	"github.com/Djoulzy/ImgStock/clog"
)

func BuildArgsList(iface interface{}, args map[string]interface{}) map[string]interface{} {
	values := reflect.ValueOf(iface)
	types := reflect.TypeOf(iface)

	for i := 0; i < types.NumField(); i++ {
		v := values.Field(i)
		switch v.Kind() {
		case reflect.Struct:
			args = BuildArgsList(v.Interface(), args)
		default:
			args[types.Field(i).Name] = v
		}
	}
	return args
}

func Process(data interface{}) {
	var args = make(map[string]interface{})

	args = BuildArgsList(data, args)
	clog.Trace("", "", "%+v", args)

	if _, valid := args["ID"]; valid {
		clog.Trace("", "", "Object is BasedItem like")
	}

	// switch unwrapped := data.(type) {
	// case datamodel.AnotherItem:
	// 	clog.Info("actions", "Process", "Unwrapped AnotherItem: %s", unwrapped.Surname)
	// case datamodel.RichItem:
	// 	clog.Info("actions", "Process", "Unwrapped RichItem: %d", unwrapped.Width)
	// }

	// for i := 0; i < schemaValues.NumField(); i++ {
	// 	fmt.Printf("%+v\n", schemaValues.Field(i))
	// }
	//
	// name := schemaValues.FieldByName("Name")
	//
	// clog.Info("actions", "Process", "Args - Name: %s", name)
	// clog.Info("actions", "Process", "Args - Width: %d", args["Width"])
}
