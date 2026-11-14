package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — общий интерфейс для тренировок и прогулок.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info перебирает dataset, парсит каждую строку и выводит информацию.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга строки %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации: %v", err)
			continue
		}

		fmt.Println(info)
	}
}
