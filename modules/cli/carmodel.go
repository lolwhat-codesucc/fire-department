package cli

import (
	"bufio"
	"context"
	"fmt"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"
)

func handleCarModel(ctx context.Context, reader *bufio.Reader, repo repository.CarModelRepository) {
	for {
		fmt.Println("\n--- Модели автомобилей ---")
		fmt.Println("1. Создать  2. Все  3. По ID  4. Обновить  5. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			fmt.Print("Название: ")
			name := readLine(reader)
			fmt.Print("Период ТО (дней): ")
			periodStr := readLine(reader)
			period, _ := strconv.Atoi(periodStr)
			id, err := repo.Create(ctx, &models.CarModel{
				Name:                  name,
				MaintenancePeriodDays: period,
			})
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Создана модель с ID=%d\n", id)
			}
		case "2":
			list, _ := repo.GetAll(ctx)
			for _, cm := range list {
				fmt.Printf("ID=%d, Название=%s, ТО=%d дн.\n", cm.ID, cm.Name, cm.MaintenancePeriodDays)
			}
		case "3":
			fmt.Print("ID: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			cm, err := repo.GetByID(ctx, id)
			if err != nil {
				fmt.Println("Не найдено")
			} else {
				fmt.Printf("ID=%d, Название=%s, ТО=%d дн.\n", cm.ID, cm.Name, cm.MaintenancePeriodDays)
			}
		case "4":
			fmt.Print("ID для обновления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			fmt.Print("Новое название: ")
			name := readLine(reader)
			fmt.Print("Новый период ТО (дней): ")
			periodStr := readLine(reader)
			period, _ := strconv.Atoi(periodStr)
			if err := repo.Update(ctx, &models.CarModel{
				ID:                    id,
				Name:                  name,
				MaintenancePeriodDays: period,
			}); err != nil {
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
