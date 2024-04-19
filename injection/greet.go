package injection

import (
	"bytes"
	"fmt"
)

func Greeting(writer *bytes.Buffer, name string){
	fmt.Fprintf(writer,"Hello, %s", name)
}

