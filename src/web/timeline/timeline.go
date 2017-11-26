package timeline

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"tools"
	"years"
)

//RegisterPath registers web path
func RegisterPath(path string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		timeLineWriter(w, r.Form["places"])
	})
}

func timeLineWriter(w http.ResponseWriter, reqPlaces []string) {
	timelineLeft, timelineRight, allPlaces := readTimeLine(os.Getenv("APOSTAXI_TIMELINE_LOCATION"), reqPlaces)
	datas := map[string]template.HTML{
		"TimelineLeft":  timelineLeft,
		"TimelineRight": timelineRight,
		"Places":        allPlaces,
	}
	tmpl, err := template.ParseFiles("../res/template/tl.tm")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, datas)
	if err != nil {
		panic(err)
	}
}

func readTimeLine(tlPath string, reqPlaces []string) (template.HTML, template.HTML, template.HTML) {
	tmpl, err := template.ParseFiles("../res/template/point.tm")
	if err != nil {
		panic(err)
	}
	places := make(map[string]bool, 0)
	pointRight := bytes.Buffer{}
	pointLeft := bytes.Buffer{}
	for _, p := range years.ReadTimeline(tlPath) {
		places[p.Place] = true
		if tools.ContainsString(reqPlaces, p.Place) {
			datas := map[string]template.HTML{
				"Title": template.HTML(p.Context),
				"Desc":  "unknown",
				"Place": template.HTML(p.Place),
				"Year":  template.HTML(strconv.Itoa(p.Years[0])),
			}
			pos := pos(reqPlaces, p.Place)
			if pos == 1 {
				tmpl.Execute(&pointRight, datas)
			} else {
				tmpl.Execute(&pointLeft, datas)
			}
		}
	}

	placesStr := bytes.Buffer{}
	for p := range places {
		datas := map[string]string{
			"Selected": selectedOption(reqPlaces, p),
			"Value":    p,
		}
		tmpl, err := template.ParseFiles("../res/template/option.tm")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(&placesStr, datas)
	}
	return template.HTML(pointLeft.String()), template.HTML(pointRight.String()), template.HTML(placesStr.String())
}

func selectedOption(places []string, place string) string {
	if tools.ContainsString(places, place) {
		return "selected"
	}
	return ""
}

func link(u string) string {
	return "<a href='" + u + "' target='_blank'>" + u + "</a>"
}

func pos(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}
