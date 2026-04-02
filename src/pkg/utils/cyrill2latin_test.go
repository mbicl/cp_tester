package utils

import "testing"

func TestCyrill2Latin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Pangram",
			input:    "\u042e\u043b\u0438\u044f, \u0441\u044a\u0435\u0448\u044c \u0435\u0449\u0451 \u044d\u0442\u0438\u0445 \u043c\u044f\u0433\u043a\u0438\u0445 \u0444\u0440\u0430\u043d\u0446\u0443\u0437\u0441\u043a\u0438\u0445 \u0431\u0443\u043b\u043e\u043a \u0438\u0437 \u0419\u043e\u0448\u043a\u0430\u0440-\u041e\u043b\u044b, \u0434\u0430 \u0432\u044b\u043f\u0435\u0439 \u0430\u043b\u0442\u0430\u0439\u0441\u043a\u043e\u0433\u043e \u0447\u0430\u044e",
			expected: "Yuliya, syesh yeshchyo etikh myagkikh frantsuzskikh bulok iz Yoshkar-Oly, da vypey altayskogo chayu",
		},
		{
			name:     "Address",
			input:    "\u0420\u043e\u0441\u0441\u0438\u044f, \u0433\u043e\u0440\u043e\u0434 \u0419\u043e\u0448\u043a\u0430\u0440-\u041e\u043b\u0430, \u0443\u043b\u0438\u0446\u0430 \u042f\u043d\u0430 \u041a\u0440\u0430\u0441\u0442\u044b\u043d\u044f",
			expected: "Rossiya, gorod Yoshkar-Ola, ulitsa Yana Krastynya",
		},
		{name: "Yeltsin", input: "\u0415\u043b\u044c\u0446\u0438\u043d", expected: "Yeltsin"},
		{name: "Razdolnoye", input: "\u0420\u0430\u0437\u0434\u043e\u043b\u044c\u043d\u043e\u0435", expected: "Razdolnoye"},
		{name: "Yuryev", input: "\u042e\u0440\u044c\u0435\u0432", expected: "Yuryev"},
		{name: "Belkin", input: "\u0411\u0435\u043b\u043a\u0438\u043d", expected: "Belkin"},
		{name: "Biysk", input: "\u0411\u0438\u0439\u0441\u043a", expected: "Biysk"},
		{name: "Podyarsky", input: "\u041f\u043e\u0434\u044a\u044f\u0440\u0441\u043a\u0438\u0439", expected: "Podyarsky"},
		{name: "Musiykyongiykote", input: "\u041c\u0443\u0441\u0438\u0439\u043a\u044a\u043e\u043d\u0433\u0438\u0439\u043a\u043e\u0442\u0435", expected: "Musiykyongiykote"},
		{name: "Davydov", input: "\u0414\u0430\u0432\u044b\u0434\u043e\u0432", expected: "Davydov"},
		{name: "Usolye", input: "\u0423\u0441\u043e\u043b\u044c\u0435", expected: "Usolye"},
		{name: "Vykhukhol", input: "\u0412\u044b\u0445\u0443\u0445\u043e\u043b\u044c", expected: "Vykhukhol"},
		{name: "Dalnegorsk", input: "\u0414\u0430\u043b\u044c\u043d\u0435\u0433\u043e\u0440\u0441\u043a", expected: "Dalnegorsk"},
		{name: "Ilyinsky", input: "\u0418\u043b\u044c\u0438\u043d\u0441\u043a\u0438\u0439", expected: "Ilyinsky"},
		{name: "Krasny", input: "\u041a\u0440\u0430\u0441\u043d\u044b\u0439", expected: "Krasny"},
		{name: "Veliky", input: "\u0412\u0435\u043b\u0438\u043a\u0438\u0439", expected: "Veliky"},
		{name: "Naberezhnye Chelny", input: "\u041d\u0430\u0431\u0435\u0440\u0435\u0436\u043d\u044b\u0435 \u0427\u0435\u043b\u043d\u044b", expected: "Naberezhnye Chelny"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cyrill2Latin(tt.input)
			if got != tt.expected {
				t.Errorf("Cyrill2Latin(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
