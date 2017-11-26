package years

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"tools"
)

type token struct {
	str        string
	sign       float64
	correction float64
}

type multiplier struct {
	str        string
	multiplier float64
}

//YearRow years period, place and mentioning context
type YearRow struct {
	Years   []int
	Context string
	Place   string
}

func (y YearRow) String() string {
	str := []string{}
	for _, year := range y.Years {
		if len(str) > 0 {
			str = append(str, ",")
		}
		str = append(str, strconv.Itoa(year))
	}
	return strings.Join(append(str, "|", y.Place, "|", y.Context, "\n"), "")
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

var tokens = []token{
	token{"лет назад", -1, 2000},
	token{"до н. э.", -1, 0},
	token{"до н.э.", -1, 0},
	token{"до новой эры", -1, 0},
	token{"до нашей эры", -1, 0},
	token{"до Р. Х.", -1, 0},
	token{"до Р.Х.", -1, 0},
	token{"год", 1, 0},
	token{"гг.", 1, 0},
	token{"г.", 1, 0},
}

var multipliers = []multiplier{
	multiplier{"в.", 100},
	multiplier{"вв.", 100},
	multiplier{"век", 100},
	multiplier{"тыс", 1000},
	multiplier{"млн", 1000000},
	multiplier{"миллион", 1000000},
}

var rangeTokens = []rune{
	'-',
	'—',
}

var romanNumbers = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func firstLeftRange(s string, sign float64, correction float64) ([]int, int) {
	years := make([]int, 0)
	preceding := []rune(s)

	year, j, ri := leftNumber(preceding, sign)
	if j >= 0 {

		multiplier := 1.0
		strmul := s[j:]

		for _, m := range multipliers {
			if strings.Contains(strmul, m.str) {
				multiplier = m.multiplier
				break
			}
		}

		years = append(years, int(year*multiplier+correction))
		ln := len(preceding[0:j])
		isNoticedRange := false
		for i := ln - 1; i >= 0; i-- {
			r := preceding[i]
			if tools.ContainsRune(rangeTokens, r) {
				isNoticedRange = true
			}
			if isDigit(r) && isNoticedRange {
				year, j, _ = leftNumber(preceding[0:i+1], sign)
				if j >= 0 {
					i = j
					years = append(years, int(year*multiplier+correction))
					if len(years) == 2 {
						break
					}
				}
			}
		}
	}
	if len(years) == 2 && years[0] < years[1] && years[0] > 0 {
		years[0] = -years[0]
		years[1] = -years[1]
	}
	sort.Ints(years)
	return years, ri
}

func leftNumber(runes []rune, sign float64) (float64, int, int) {
	ln := len(runes)
	leftDigitind := -1
	rightDigitInd := -1

	for i := ln - 1; i >= 0; i-- {
		r := runes[i]
		if isDigit(r) && rightDigitInd < 0 {
			rightDigitInd = i + 1
		} else if rightDigitInd >= 0 {
			if !isDigit(r) && ',' != r && '.' != r {
				leftDigitind = i + 1
				break
			}
			if i == 0 {
				leftDigitind = i
			}
		}
	}

	if leftDigitind >= 0 && rightDigitInd >= 0 {
		s := string(runes[leftDigitind:rightDigitInd])
		if strings.Contains(s, ",") && !strings.Contains(s, ".") {
			s = strings.Replace(s, ",", ".", -1)
		}
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			num, err = parseRoman(s)
		}
		if err == nil {
			return sign * num, leftDigitind, rightDigitInd
		}
	}
	return -1, -1, -1
}

func isDigit(r rune) bool {
	if unicode.IsDigit(r) {
		return true
	}
	if _, ok := romanNumbers[r]; ok {
		return true
	}
	return false
}

func parseRoman(roman string) (float64, error) {
	number := 0
	lastDigit := 1000
	for _, romRune := range []rune(roman) {
		digit := romanNumbers[romRune]
		if lastDigit < digit {
			number -= 2 * lastDigit
		}
		lastDigit = digit
		number += lastDigit
	}
	if number != 0 {
		return float64(number), nil
	}
	return -1, errors.New("Cant parse roman digits")
}

//WalkRow walks row from right to left
func WalkRow(str, place string) []YearRow {
	yrs := make([]YearRow, 0)
	lastRightInd := -1
	for _, t := range tokens {
		i := len(str)
		for i > 0 {
			i = strings.LastIndex(str[0:i], t.str)
			if i > 0 {
				leftRange, rightInd := firstLeftRange(str[0:i], t.sign, t.correction)
				if rightInd != lastRightInd {
					lastRightInd = rightInd
					yr := YearRow{leftRange, str, place}
					yrs = append(yrs, yr)
				}
			}
		}
	}
	return yrs
}
