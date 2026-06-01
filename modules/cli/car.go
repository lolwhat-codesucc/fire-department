package cli

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"fire-department/modules/models"
	"fire-department/modules/repository"
)

func handleCar(ctx context.Context, reader *bufio.Reader, repo repository.CarRepository) {
	for {
		fmt.Println("\n--- Автомобили ---")
		fmt.Println("1. Создать  2. Все  3. По ID  4. Обновить  5. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			var car models.Car
			fmt.Print("ID модели: ")
			car.ModelID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("Дата приобретения (YYYY-MM-DD): ")
			dateStr := readLine(reader)
			t, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				fmt.Println("Неверный формат даты")
				continue
			}
			car.AcquisitionDate = models.Date{Time: t}
			fmt.Print("Готовность (true/false): ")
			readyStr := readLine(reader)
			car.Ready = strings.ToLower(readyStr) == "true"
			fmt.Print("Дата последнего ТО (YYYY-MM-DD, Enter для пропуска): ")
			maintStr := readLine(reader)
			if maintStr != "" {
				maintTime, err := time.Parse("2006-01-02", maintStr)
				if err == nil {
					car.LastMaintenance = &models.Date{Time: maintTime}
				}
			}
			id, err := repo.Create(ctx, &car)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Создан автомобиль с ID=%d\n", id)
			}
		case "2":
			list, _ := repo.GetAll(ctx)
			for _, c := range list {
				fmt.Printf("ID=%d, ModelID=%d, Ready=%t\n", c.ID, c.ModelID, c.Ready)
			}
		case "3":
			fmt.Print("ID: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			car, err := repo.GetByID(ctx, id)
			if err != nil {
				fmt.Println("Не найден")
			} else {
				fmt.Printf("ID=%d, ModelID=%d, Ready=%t, Acquired=%s\n",
					car.ID, car.ModelID, car.Ready, car.AcquisitionDate.Time.Format("2006-01-02"))
			}
		case "4":
			fmt.Print("ID для обновления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			var car models.Car
			car.ID = id
			fmt.Print("ID модели: ")
			car.ModelID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("Дата приобретения (YYYY-MM-DD): ")
			dateStr := readLine(reader)
			t, _ := time.Parse("2006-01-02", dateStr)
			car.AcquisitionDate = models.Date{Time: t}
			fmt.Print("Готовность (true/false): ")
			readyStr := readLine(reader)
			car.Ready = strings.ToLower(readyStr) == "true"
			fmt.Print("Дата последнего ТО (YYYY-MM-DD, Enter для пропуска): ")
			maintStr := readLine(reader)
			if maintStr != "" {
				maintTime, _ := time.Parse("2006-01-02", maintStr)
				car.LastMaintenance = &models.Date{Time: maintTime}
			}
			if err := repo.Update(ctx, &car); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Обновлено")
			}
		case "5":
			fmt.Print("ID для удаления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			if err := repo.Delete(ctx, id); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Удалено")
			}
		case "0":
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}
