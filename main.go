package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

// Структура одной строки таблицы
type LotRow struct {
	Columns []string `json:"columns"`
}

func main() {
	// Настройка Chrome с GUI
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Показываем окно
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("no-sandbox", true),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Таймаут
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var rawRows []string

	err := chromedp.Run(ctx,
		// Открываем сайт
		chromedp.Navigate("https://goszakup.gov.kz/ru/search/lots"),

		// Ждём таблицу
		chromedp.Sleep(4*time.Second),

		// Забираем все строки таблицы
		chromedp.Evaluate(`
			Array.from(document.querySelectorAll("tbody tr")).map(tr => 
				Array.from(tr.querySelectorAll("td")).map(td => td.textContent.trim()).join("|")
			)
		`, &rawRows),
	)

	if err != nil {
		log.Fatal("❌ Ошибка загрузки:", err)
	}

	// Обработка и преобразование
	var lots []LotRow
	for _, row := range rawRows {
		cols := strings.Split(row, "|")
		lots = append(lots, LotRow{Columns: cols})
	}

	// Сохраняем в JSON
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal("❌ Не удалось создать файл:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // форматированный JSON
	if err := encoder.Encode(lots); err != nil {
		log.Fatal("❌ Ошибка записи JSON:", err)
	}

	fmt.Printf("✅ Сохранено %d строк в файл data.json\n", len(lots))
}
