package logger

import (
	"log"

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

func Warn(textTemplate string, params ...interface{}) {
	colorPrint.Yellow("\n"+INDENT+"WARNING: "+textTemplate, params...)
}

func Debug(textTemplate string, params ...interface{}) {
	colorPrint.Magenta("\n"+INDENT+"DEBUG: "+textTemplate, params...)
}

func Error(textTemplate string, params ...interface{}) {
	colorPrint.Red("\n"+INDENT+"ERROR: "+textTemplate, params...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v)
}
