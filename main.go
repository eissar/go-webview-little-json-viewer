package main

import (
	"encoding/binary"
	"fmt"
	webview "github.com/webview/webview_go"
	_ "io"
	"os"
	_ "path/filepath"
	"strings"
)

func writeHost(msgs ...string) {
	joinedString := strings.Join(msgs, ", ")
	fmt.Fprintln(os.Stderr, joinedString)
}

func getContents(fp string) string {
	file := func() *os.File {
		fmt.Fprintln(os.Stderr, "[INFO]: Opening url")
		file, err := os.Open(fp)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file url:", file, err)
			return nil
		}
		return file // Return the file without deferring the close
	}()
	defer file.Close()

	/* get file length in bytes */
	fileInfo, err := file.Stat()

	if err != nil {
		writeHost("err while getting file details", err.Error())
		panic("")
	}
	fileSize := fileInfo.Size()
	byteArr := make([]byte, fileSize)
	binary.LittleEndian.PutUint64(byteArr, uint64(fileSize))
	_, err = file.Read(byteArr) // bytes, err
	if err != nil {
		writeHost("err while reading bytes of file", fileInfo.Name(), err.Error())
		panic("")
	}
	contents := string(byteArr)
	return contents
}

func main() {
	view := webview.New(true)
	defer view.Destroy()
	view.SetTitle("Basic Example")
	view.SetSize(480, 320, webview.HintNone)

	/* init by loading in the js files */

	// view.Init(getContents("C:/Users/eshaa/Dropbox/Code/go/webview/ace-builds/src/ace.js"))
	// view.Init(getContents("C:/Users/eshaa/Dropbox/Code/go/webview/ace-builds/src/ext-themelist.js"))
	// view.Init(getContents("C:/Users/eshaa/Dropbox/Code/go/webview/ace-builds/src/ext-language_tools.js"))

	contents := getContents("C:/Users/eshaa/Dropbox/Code/go/webview/index.html")
	view.SetHtml(contents)
	view.Bind("hello", func() {
		fmt.Println("hello from js!")
	})
	view.Run()
}
