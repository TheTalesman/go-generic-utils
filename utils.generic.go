package controller

import (
	"fmt"
	"time"
)

//NeedleHaystackValue algoritmo para encontrar uma string em uma array de string, retorna um boolean.
func NeedleHaystackValue(haystack []string, needle string) (encontrado bool) {
	for _, val := range haystack {
		if val == needle {
			encontrado = true
			return
		}
	}
	encontrado = false

	return
}

//NeedleHaystackKey algoritmo para encontrar o índice de uma string em uma array de string, retorna o índice e um boolean true se for encontrado ou false se a string não for encontrada.
func NeedleHaystackKey(haystack []string, needle string) (indice int, encontrado bool) {
	for key, val := range haystack {
		if val == needle {
			indice = key
			encontrado = true
			return
		}
	}
	encontrado = false
	return
}

//ValidaDia verifica se uma string de data é uma data válida
func ValidaDia(dia string) (valido bool) {
	layout := "02/01/2006"
	t, err := time.Parse(layout, dia)
	if err != nil {
		fmt.Println("Dia inválido: " + t.String())
		fmt.Println("Erro: " + err.Error())
		valido = false
		return
	}
	fmt.Println("Dia válido: " + t.String())
	valido = true
	return
}

//RetornaDiaDaSemana Recebe data e passa dia da semana
func RetornaDiaDaSemana(dia string) (diaTratado string) {
	layout := "02/01/2006"

	t, _ := time.Parse(layout, dia)
	fmt.Println(t.Format(layout))
	weekday := t.Weekday()
	diaTratado, _ = parseWeekday(weekday)

	return
}

var daysOfWeek = map[time.Weekday]string{
	time.Sunday:    "Domingo",
	time.Monday:    "Segunda",
	time.Tuesday:   "Terça-feira",
	time.Wednesday: "Quarta-feira",
	time.Thursday:  "Quinta-feira",
	time.Friday:    "Sexta-feira",
	time.Saturday:  "Sábado",
}

func parseWeekday(v time.Weekday) (string, error) {
	if d, ok := daysOfWeek[v]; ok {
		return d, nil
	}

	return "", fmt.Errorf("invalid weekday '%s'", v)
}

//GenToString Gera string a partir de qq coisa
func GenToString(qualquerCoisa interface{}) (stringTratada string) {
	stringTratada = fmt.Sprintf("%v", qualquerCoisa)
	return
}
