package logwrapper

import (
	"fmt"

)

func New(servicename string) *Logger{
   var log Logger
    if (servicename != ""){
    fmt.Printf(servicename)
     log:= new (Logger)
     log.serviceName = servicename

  }
  return &log
}


type Logger struct {
	serviceName string
}
