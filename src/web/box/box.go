package box

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/russross/blackfriday"
)

//DataPath location of knowledge
var DataPath = getDataPath()

func getDataPath() string {
	return os.Getenv("APOSTAXI_KRIFES_LOCATION")
}

//RegisterPath registers web path
func RegisterPath(path string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		aWriter(w)
	})
	http.HandleFunc("/dir", func(w http.ResponseWriter, r *http.Request) {
		dirWriter(w, r.FormValue("path"))
	})
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		fileWriter(w, r.FormValue("path"))
	})
}

func fileWriter(w io.Writer, path string) {
	b := readFile(path)
	w.Write(b.Bytes())
}

func dirWriter(w io.Writer, path string) {
	b := readPath(path)
	w.Write(b.Bytes())
}

func aWriter(w io.Writer) {
	datas := map[string]template.HTML{
		"Pillar": readRoot(DataPath),
	}
	tmpl, err := template.ParseFiles("../res/template/title.tm")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, datas)
	if err != nil {
		panic(err)
	}
}

func readPath(path string) bytes.Buffer {
	fileDir, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	bw := bufio.NewWriter(&b)
	for _, fi := range fileDir {
		includePath := !fi.IsDir()
		bw.WriteString("<li><a ic-push-url='true' ic-post-to='/")
		if includePath {
			bw.WriteString("file")
		} else {
			bw.WriteString("dir")
		}
		bw.WriteString("?path=")
		bw.WriteString(path)
		if !strings.HasSuffix(path, "/") {
			bw.WriteString("/")
		}
		bw.WriteString(fi.Name())
		bw.WriteString("' ic-target='#")
		if includePath {
			bw.WriteString("contentid")
		} else {
			bw.WriteString("pillarRight")
		}
		bw.WriteString("'>")
		bw.WriteString(fi.Name())
		bw.WriteString("</a></li>")

		//todo move into template
		//todo read
	}
	bw.Flush()
	return b
}

func readRoot(path string) template.HTML {
	b := readPath(path)
	return template.HTML(b.String())
}

func readFile(path string) bytes.Buffer {
	fmt.Println("received path", path)
	npath := "/pic" + strings.Replace(path, DataPath, "", -1)
	lind := strings.LastIndex(npath, "/") + 1
	npath = npath[0:lind]
	fmt.Println("cut path", npath, lind)

	f, _ := ioutil.ReadFile(path)
	sf := strings.Replace(string(f), "](", "]("+npath, -1)
	md := blackfriday.MarkdownBasic([]byte(sf))

	var b bytes.Buffer
	bw := bufio.NewWriter(&b)
	bw.Write(md)
	bw.Flush() //todo get rid of unnecessay code
	return b
}
