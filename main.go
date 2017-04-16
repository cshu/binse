package main
//note there is an issue if you use CXXFLAGS (https://github.com/golang/go/issues/12516)
//compile a shared library to circumvent it
//#cgo CXXFLAGS: -x c++ -std=c++14 -pedantic-errors -Wall -Wextra -Wshadow

/*
#cgo LDFLAGS: -L. -lbinse
#include <stdlib.h>
unsigned char cxxinit(char *);
unsigned char cxxse(long long,char *);
*/
import "C"
import (
	"fmt"
	"log"
	"os"
	"flag"
	"regexp"
	_"encoding/hex"
	"path/filepath"
	"io/ioutil"
	"unsafe"
)

//var hexfl = flag.Bool("hex", false, "Decode args as base16")
//var refl = flag.Bool("re", false, "Treat args as regular expression")
//var posixfl //undone MustCompilePOSIX?
var cxxfl = flag.Bool("cxx", false, "Use C++ standard library instead")

func main() {
	flag.Parse()
	if flag.NArg()==0 {os.Exit(1)}
	//sbuf:=flag.Args()
	//if *hexfl{//undone hexfl
	//}
	//undone refl
	if *cxxfl{
		cstr := C.CString(flag.Arg(0))
		defer C.free(unsafe.Pointer(cstr))
		if 0!=C.cxxinit(cstr){
			os.Exit(1)
		}
		err:=filepath.Walk(".",func(path string, info os.FileInfo, err error)error{
			if info.Mode().IsRegular(){
				bytes,err:=ioutil.ReadFile(path)//undone big file causes problem
				if err!=nil{
					log.Println(err)
					return nil//note
				}
				if len(bytes)==0{return nil}
				switch C.cxxse(C.longlong(len(bytes)),(*C.char)(unsafe.Pointer(&bytes[0]))){
				case 0:
				case 1:
					fmt.Println(path)
				default:
					os.Exit(1)//? is it okay to exit here?
				}
			}
			return nil
		})
		if err!=nil{
			os.Exit(1)
		}
	}else{
		re:=regexp.MustCompile(flag.Arg(0))
		err:=filepath.Walk(".",func(path string, info os.FileInfo, err error)error{
			//undone checking err is necessary?
			//if !info.IsDir()
			if info.Mode().IsRegular(){
				bytes,err:=ioutil.ReadFile(path)//undone big file causes problem
				if err!=nil{
					log.Println(err)
					return nil//note
				}
				if re.Match(bytes){
					fmt.Println(path)
				}
			}
			return nil
		})
		if err!=nil{
			os.Exit(1)
		}
	}
}
