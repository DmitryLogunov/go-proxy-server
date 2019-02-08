package logger

import (
	"log"
	"os"

	colorPrint "github.com/fatih/color"
)

const INDENT = "   "

func Header(text string) {
	colorPrint.Green("\n" + INDENT + " ------------ " + text + " -----------\n\n")
}

func Info(textTemplate string, params ...interface{}) {
	colorPrint.White(INDENT+textTemplate, params...)
}

func Text(text string) {
	colorPrint.White(INDENT + text + "\n")
}

func Debug(textTemplate string, params ...interface{}) {
	if os.Getenv("LOG_LEVEL") == "debug" {
		colorPrint.Magenta("\n"+INDENT+textTemplate, params...)
	}	
}

func Warn(textTemplate string, params ...interface{}) {
	colorPrint.Yellow("\n"+INDENT+"WARNING: "+textTemplate, params...)
}


func Error(textTemplate string, params ...interface{}) {
	colorPrint.Red("\n"+INDENT+"ERROR: "+textTemplate, params...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v)
}
