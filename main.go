package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"regexp"
	_"encoding/hex"
	"path/filepath"
	"io/ioutil"
)

//var hexfl = flag.Bool("hex", false, "Decode args as base16")
//var refl = flag.Bool("re", false, "Treat args as regular expression")
//var posixfl //undone MustCompilePOSIX?

func main() {
	flag.Parse()
	if flag.NArg()==0 {os.Exit(1)}
	//sbuf:=flag.Args()
	//if *hexfl{//undone hexfl
	//}
	//undone refl
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
