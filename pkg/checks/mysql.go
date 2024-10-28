// Code generated by scorify, DO NOT EDIT.
// Import generated from "github.com/scorify/mysql@v0.0.2"

package checks

import (
	"github.com/scorify/schema"

	mysql "github.com/scorify/mysql"
)

func init() {
	schema, err := schema.Describe(mysql.Schema{})
	if err != nil {
		panic(err)
	}

	Checks["mysql"] = Check{
		Func:     mysql.Run,
		Validate: mysql.Validate,
		Schema:   schema,
	}
}