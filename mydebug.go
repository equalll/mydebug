

package mydebug

import (
  "fmt"
  "log"
  "strings"
  "runtime"
 _"reflect"
  "path/filepath"
)

type MyStruct struct {
}

func (m *MyStruct) foo(p string) {
  ENTRY("")
  ENTRY("Param p=%s", p)
  DEBUG("Test %s %s", "Hello", "World")
}

func DEBUG(formating string, args... interface{}) {
  LOG("DEBUG", formating, args...)
}

func ENTRY(formating string, args... interface{}) {
  LOG("ENTRY", formating, args...)
}

func LOG(level string, formating string, args... interface{}) {
  filename, line, funcname := "???", 0, "???"
  pc, filename, line, ok := runtime.Caller(2)
  // fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
  if ok {
      funcname = runtime.FuncForPC(pc).Name()       // main.(*MyStruct).foo
      funcname = filepath.Ext(funcname)             // .foo
      funcname = strings.TrimPrefix(funcname, ".")  // foo

      filename = filepath.Base(filename)  // /full/path/basename.go => basename.go
  }

  log.Printf("%s:%d:%s: %s: %s\n", filename, line, funcname, level, fmt.Sprintf(formating, args...))
}

func INFO() {
  filename, line, funcname := "???", 0, "???"
  pc, filename, line, ok := runtime.Caller(2)
  // fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))
  if ok {
      funcname = runtime.FuncForPC(pc).Name()       // main.(*MyStruct).foo
      funcname = filepath.Ext(funcname)             // .foo
      funcname = strings.TrimPrefix(funcname, ".")  // foo

      filename = filepath.Base(filename)  // /full/path/basename.go => basename.go
  }

  log.Printf("%s:%d:%s\n", filename, line, funcname)
}

func main() {
  ss := MyStruct{}
  ss.foo("helloworld")
}