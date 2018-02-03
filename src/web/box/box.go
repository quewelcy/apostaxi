package box

import (
	"bufio"
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/russross/blackfriday"
)

//DataPath location of knowledge
var DataPath = os.Getenv("APOSTAXI_KRIFES_LOCATION")
var htmlPath = "/box"

//RegisterPath registers web path
func RegisterPath() {
	http.HandleFunc(htmlPath, boxHandler)
}

func boxHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]template.HTML{
		"dir":     template.HTML(readDir(r.FormValue("p"))),
		"content": template.HTML(getContent(r.FormValue("c"))),
	}
	tmpl, _ := template.ParseFiles("../res/template/title.tm")
	tmpl.Execute(w, data)
}

func readDir(path string) string {
	if path == "" {
		path = DataPath
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var b bytes.Buffer
	for _, file := range files {

		var href string
		if file.IsDir() {
			href = htmlPath + "?p=" + getProperPath(path, file.Name())
		} else {
			href = htmlPath + "?p=" + path + "&c=" + getProperPath(path, file.Name())
		}
		data := map[string]string{
			"href":  href,
			"fname": file.Name(),
		}
		tmpl, _ := template.ParseFiles("../res/template/file.tm")
		tmpl.Execute(&b, data)
	}
	return b.String()
}

func getContent(p string) string {
	if p == "" {
		return ""
	}
	b := readMdFile(p)
	return b.String()
}

func getProperPath(dir, file string) string {
	if strings.HasSuffix(dir, string(os.PathSeparator)) {
		return dir + file
	}
	return dir + string(os.PathSeparator) + file
}

func readMdFile(path string) bytes.Buffer {
	lind := strings.LastIndex(path, string(os.PathSeparator))
	npath := "/pic/?p=" + strings.Replace(path[0:lind+1], `\`, `\\`, -1)
	f, _ := ioutil.ReadFile(path)
	sf := strings.Replace(string(f), "](", "]("+npath, -1)
	md := blackfriday.MarkdownBasic([]byte(sf))

	var b bytes.Buffer
	bw := bufio.NewWriter(&b)
	bw.Write(md)
	bw.Flush()
	return b
}
