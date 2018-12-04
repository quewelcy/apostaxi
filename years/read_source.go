package years

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type byYearRow []YearRow

func (y byYearRow) Len() int {
	return len(y)
}
func (y byYearRow) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}
func (y byYearRow) Less(i, j int) bool {
	return len(y[i].Years) < len(y[j].Years) || y[i].Years[0] < y[j].Years[0]
}

//ReadSourceAndSaveTimeline reads source of years
func ReadSourceAndSaveTimeline(readPath, place, writePath string, overwrite bool) {
	timeline := []YearRow{}
	if file, err := os.Open(readPath); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			yr := LookupYearRow(scanner.Text(), place)
			log.Println("parsed", yr)
			if len(yr) > 0 {
				timeline = append(timeline, yr...)
			}
		}
	} else {
		log.Fatal(err)
		return
	}
	sort.Sort(byYearRow(timeline))

	d1 := []byte{}
	for _, y := range timeline {
		d1 = append(d1, []byte(y.String())...)
	}

	if !overwrite {
		by, err := ioutil.ReadFile(writePath)
		if err == nil {
			d1 = append(d1, by...)
		}

	}
	err := ioutil.WriteFile(writePath, d1, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//ReadTimeline reads years and context
func ReadTimeline(path string) []YearRow {
	timeline := []YearRow{}
	if file, err := os.Open(path); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			timeline = append(timeline, yearRowFromString(scanner.Text()))
		}
	} else {
		log.Fatal(err)
	}
	return timeline
}

func yearRowFromString(str string) YearRow {
	r := YearRow{}
	colonInd := strings.Index(str, "|")
	if colonInd > 0 {
		rangeInd := strings.Index(str[0:colonInd], ",")
		if rangeInd > 0 {
			num, _ := strconv.ParseInt(str[:rangeInd], 10, 64)
			r.Years = append(r.Years, int(num))
			num, _ = strconv.ParseInt(str[rangeInd+len(",")+1:colonInd], 10, 64)
			r.Years = append(r.Years, int(num))
		} else {
			num, _ := strconv.ParseInt(str[:colonInd], 10, 64)
			r.Years = append(r.Years, int(num))
		}
		colonInd2 := colonInd + 1 + strings.Index(str[colonInd+1:], "|")
		r.Place = str[colonInd+1 : colonInd2]
		r.Context = str[colonInd2+1:]
	}
	return r
}
