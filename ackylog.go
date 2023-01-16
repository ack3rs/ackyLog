package ackyLog

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type stack struct {
	Line    int
	Package string
	File    string
	Fname   string
}

var SHOWCOLOURS = true
var SHOWDEBUG = true

//TODO: Put all the possible colours in like this.  I'd like more options but not to the point of having 
// to remember a ton of colours.
// I kinda like this... https://gist.github.com/RabaDabaDoba/145049536f815903c79944599c6f952a

var ColourMap = map[string]string{
	"F-BLACK":   "\033[1;30m",
	"F-RED":     "\033[1;31m",
	"F-GREEN":   "\033[1;32m",
	"F-YELLOW":  "\033[1;33m",
	"F-BLUE":    "\033[1;34m",
	"F-MAGENTA": "\033[1;35m",
	"F-CYAN":    "\033[1;36m",
	"F-WHITE":   "\033[1;37m",
	"F-NORMAL":  "\033[0;37m", // Normal White (Not Bold )
	"F-RESET":   "\033[0m",
	"B-BLACK":   "\033[1;40m",
	"B-RED":     "\033[1;41m",
	"B-GREEN":   "\033[1;42m",
	"B-YELLOW":  "\033[1;43m",
	"B-BLUE":    "\033[1;44m",
	"B-MAGENTA": "\033[1;45m",
	"B-CYAN":    "\033[1;46m",
	"B-WHITE":   "\033[1;47m",
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

// WEBFORM - Dump all Web Values to the Log. (Carefull with Passwords)
func WEBFORM(Message string, r *http.Request, variables ...interface{}) {
	r.ParseForm()
	for key, values := range r.Form {
		for _, value := range values {
			sendOut(fmt.Sprintf(Message, variables...) + fmt.Sprintf(" KEY:[F-MAGENTA]%s[F-NORMAL] is VALUE:[F-MAGENTA]%v[F-NORMAL]", key, value))
		}
	}
}

// TIMED - Start a timer and return a function to call once complete.
func TIMED(Message string, variables ...interface{}) func() {
	t := time.Now()
	return func() {
		d := time.Since(t)
		out := "[[F-CYAN]TIMED[F-NORMAL]] " + fmt.Sprintf(Message, variables...) + " took " + d.String()
		sendOut(out)
	}
}

// WEB - Write a WEB Message to the Logs
func WEB(Message string, r *http.Request, variables ...interface{}) {
	webInfo := "[" + r.RemoteAddr + " - " + r.Method + " - " + r.URL.String() + "] "
	out := "[[F-GREEN]<-WEB[F-NORMAL]] " + webInfo + fmt.Sprintf(Message, variables...)
	sendOut(out)
}

// WARNING - Write a Warning message to the Log
func WARNING(Message string, variables ...interface{}) {
	out := "[[F-YELLOW]WARNING[F-NORMAL]] " + fmt.Sprintf(Message, variables...)
	sendOut(out)
}

// ERROR - Write an ERROR message to the LOG
func ERROR(Message string, variables ...interface{}) {
	out := "[[F-RED]ERROR[F-NORMAL]] " + fmt.Sprintf(Message, variables...)
	sendOut(out)
}

// DEBUG - Write a DEBUG message to the LOG
func DEBUG(Message string, variables ...interface{}) {
	out := "[[F-CYAN]DEBUG[F-NORMAL]] " + fmt.Sprintf(Message, variables...)
	if SHOWDEBUG {
		sendOut(out)
	}
}

// INFO - Write an INFO Message and write to the logs.
func INFO(Message string, variables ...interface{}) {
	out := "[[F-GREEN]INFO[F-NORMAL]] " + fmt.Sprintf(Message, variables...)
	sendOut(out)
}

// SPEW - Dump a variable to the Logs.
func SPEW(variables ...interface{}) {
	raw := spew.Sdump(variables...)
	out := "[SPEW] " + fmt.Sprintf(raw)
	sendOut(out)
}

// CUSTOM - Just in Case you want a Custom Header
func CUSTOM(Header string, Message string, variables ...interface{}) {
	out := Header + fmt.Sprintf(Message, variables...)
	sendOut(out)
}

// RAW  - Without the Sprintf
func RAW(Message string) {
	out := "[[F-RED]RAW[F-NORMAL]] " + Message
	sendOut(out)
}

// Apparently I am going straight to Hell for using this ...
// https://blog.sgmansfield.com/2015/12/goroutine-ids/
// But in a busy application with lots of go rountines,  I have found it useful to see the thread id on the logs.
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func stackTrace() stack {

	CallStackTrack := stack{}

	// What called this function
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return CallStackTrack
	}

	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	partslength := len(parts)
	fname := parts[partslength-1]
	pname := strings.Join(parts[0:partslength-1], ".")

	pathSlice := strings.Split(pname, "/")
	pathLength := len(pathSlice)
	packageName := pathSlice[pathLength-1]

	CallStackTrack.Package = packageName
	CallStackTrack.File = fileName
	CallStackTrack.Fname = fname
	CallStackTrack.Line = line

	return CallStackTrack
}


func colourReplacement(LogEntryMessage string) string {

	// Does the LogEntryMessage contain any Strings
	// Loop through the ColourMap and Replace all the Strings

	for k, v := range ColourMap {
		if SHOWCOLOURS {
			LogEntryMessage = strings.ReplaceAll(LogEntryMessage, "["+k+"]", v)
		} else {
			LogEntryMessage = strings.ReplaceAll(LogEntryMessage, "["+k+"]", "")
		}
	}
	return LogEntryMessage
}

func sendOut(logEntry string) {

	st := stackTrace()
	logEntry = fmt.Sprintf("[F-NORMAL][[F-GREEN]%d[F-NORMAL]] [[F-WHITE]%v/%v[F-NORMAL] %v:%v] ", getGID(), st.Package, st.File, st.Fname, st.Line) + logEntry
	logEntry = colourReplacement(logEntry)
	_ = log.Output(1, string(logEntry))
}
