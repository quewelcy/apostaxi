package years_test

import (
	"testing"

	"github.com/quewelcy/apostaxi/years"
)

func TestYears(t *testing.T) {
	// testYear(t, "и палеогенового периода, около 66 миллионов лет назад.", []int{-65998000})
	// testYear(t, "Древний Египет во II тыс. до н.э.  (карта).", []int{-2000})
	// testYear(t, "522-486 годы. Дарий I", []int{-522, -486})
	// testYear(t, "486-464 годы. Ксеркс I", []int{-486, -464})
	// testYear(t, "Поздний (Ливийско-Саисский и Персидский) период(XI - VI века до Р. Х.)", []int{-1100, -600})
	// testYear(t, "Хор Аха 3007—2970 гг.", []int{-3007, -2970})
	// testYear(t, "Хотепсехемуи около 2.84 тыс лет до н. э.", []int{-2840})
	// testYear(t, "380 до н. э. После Неферита II правит Тридцатая династия (380—343 до н. э.), основанная Нектанебом Первым", []int{-380, -343})
	// testYear(t, "именно 380 — 363 до н. э. При Нектанебе I наблюдается некоторое усиление Египта", []int{-380, -363})
	// testYear(t, "363—360 до н. э. Правление фараона Тахоса.", []int{-363, -360})
	// testYear(t, "В VI веке до н. э. - Колхидское царство.", []int{-600})
	// testYear(t, "В IV-III веках до н. э. - Иберия.", []int{-400, -300})
	// testYear(t, "Анаксимандр (ок. 610 — после 547 до н. э.), древнегреческий философ", []int{-610, -547})

	testYear(t, "In 305 BC, Ptolemy took the title of Pharaoh. That was to rule Egypt for nearly 300 years.", []int{-305})
	testYear(t, "Ptolemy successfully defended Egypt against an invasion by Perdiccas in 321 B.C.", []int{-321})
	testYear(t, "during the Wars of the Diadochi (322–301 BC).", []int{-322, -301})
	testYear(t, "At the end of the II century BC, the Pharnavazid king Pharnajom was dethroned", []int{-200})
	testYear(t, "Arshak ascended the Iberian throne in 93 BC.", []int{-93})

	// testYear(t, "by Claudius II Gothicus at the Battle of Naissus in 268 or 269.", []int{268, 269})
	// testYear(t, "By late 274, the Roman Empire had been reunited into a single entity.", []int{-274})
}

func testYear(t *testing.T, r string, expYears ...[]int) {
	years := years.LookupYearRow(r, "")
	for i, expYearSet := range expYears {
		for j, expYear := range expYearSet {
			if i >= len(years) {
				t.Error("Not all expected years parsed. Expected", expYears, "received", years)
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
