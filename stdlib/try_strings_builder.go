package stdlib

import (
	"fmt"
	"strings"
)

func StrBuilder() {
	var builder strings.Builder
	builder.Write([]byte("Test Strings.Builder"))
	builder.Write([]byte("\\div\\"))
	builder.Write([]byte("OK"))
	fmt.Println(builder.String())
	builder.Reset()
	builder.WriteString("this is a string")
	builder.WriteString("\\did\\")
	fmt.Println(builder.String())
}
