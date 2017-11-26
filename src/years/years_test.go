package years_test

import (
	"testing"

	"years"
)

func TestYears(t *testing.T) {
	testYear(t, "и палеогенового периода, около 66 миллионов лет назад.", []int{-65998000})
	testYear(t, "Древний Египет во II тыс. до н.э.  (карта).", []int{-2000})
	testYear(t, "522-486 годы. Дарий I", []int{-522, -486})
	testYear(t, "486-464 годы. Ксеркс I", []int{-486, -464})
	testYear(t, "Поздний (Ливийско-Саисский и Персидский) период(XI - VI века до Р. Х.)", []int{-1100, -600})
	testYear(t, "Хор Аха 3007—2970 гг.", []int{-3007, -2970})
	testYear(t, "Хотепсехемуи около 2.84 тыс лет до н. э.", []int{-2840})
	testYear(t, "380 до н. э. После Неферита II правит Тридцатая династия (380—343 до н. э.), основанная Нектанебом Первым", []int{-380, -343})
	testYear(t, "именно 380 — 363 до н. э. При Нектанебе I наблюдается некоторое усиление Египта", []int{-380, -363})
	testYear(t, "363—360 до н. э. Правление фараона Тахоса.", []int{-363, -360})
	testYear(t, "В VI веке до н. э. - Колхидское царство.", []int{-600})
	testYear(t, "В IV-III веках до н. э. - Иберия.", []int{-400, -300})
	testYear(t, "Анаксимандр (ок. 610 — после 547 до н. э.), древнегреческий философ", []int{-610, -547})
}

func testYear(t *testing.T, r string, expYears ...[]int) {
	years := years.WalkRow(r, "")
	for i, expYearSet := range expYears {
		for j, expYear := range expYearSet {
			if i >= len(years) {
				t.Error("Not enough size", i, len(years))
				return
			}
			yr := years[i]
			if expYear != yr.Years[j] {
				t.Error("Expected", expYear, "got", yr.Years[j])
				return
			}
		}
	}
}
