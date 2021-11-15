package titlecase

import (
	"strings"
)

// TitleCase(str, minor) возвращает строку со всеми словами в верхнем регистре
// Первое слово всегда должно быть с заглавной
// При этом также есть строка со словами, которые нужно оставить в нижнем регистре (передается вторым аргументом)
//
// примеры:
// TitleCase("the quick fox in the bag", "") = "The Quick Fox In The Bag"
// TitleCase("the quick fox in the bag", "in the") = "The Quick Fox in the Bag"
//
// Задание: написать тесты:)

func TitleCase(str, minor string) string {
	minorWords := make(map[string]struct{})
	for _, s := range strings.Split(minor, " ") {
		minorWords[s] = struct{}{}
	}

	words := strings.Split(strings.ToLower(str), " ")
	var ret strings.Builder
	for i, w := range words {
		_, isMinor := minorWords[w]
		if i == 0 || !isMinor {
			w = strings.Title(w)
		}

		ret.WriteString(w)

		if i != len(words)-1 {
			ret.WriteString(" ")
		}
	}

	return ret.String()
}
