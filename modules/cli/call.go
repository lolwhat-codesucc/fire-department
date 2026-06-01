package cli

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"time"

	"fire-department/modules/models"
	"fire-department/modules/repository"
)

func handleCall(ctx context.Context, reader *bufio.Reader, repo repository.CallRepository) {
	for {
		fmt.Println("\n--- Вызовы ---")
		fmt.Println("1. Создать  2. Все  3. По ID  4. Обновить  5. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			var call models.Call
			fmt.Print("ID команды: ")
			call.TeamID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID автомобиля: ")
			call.CarID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID района (Enter для пропуска): ")
			distStr := readLine(reader)
			if distStr != "" {
				distID, _ := strconv.Atoi(distStr)
				call.DistrictID = &distID
			}
			fmt.Print("ID статуса: ")
			call.StatusID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("Время (YYYY-MM-DD HH:MM): ")
			timeStr := readLine(reader)
			t, err := time.Parse("2006-01-02 15:04", timeStr)
			if err != nil {
				fmt.Println("Неверный формат времени")
				continue
			}
			call.Time = t
			fmt.Print("Комментарий (Enter для пропуска): ")
			comment := readLine(reader)
			if comment != "" {
				call.Comment = &comment
			}
			id, err := repo.Create(ctx, &call)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Создан вызов с ID=%d\n", id)
			}
		case "2":
			list, _ := repo.GetAll(ctx)
			for _, c := range list {
				fmt.Printf("ID=%d, Команда=%d, Авто=%d, Время=%s, Статус=%d\n",
					c.ID, c.TeamID, c.CarID, c.Time.Format("2006-01-02 15:04"), c.StatusID)
			}
		case "3":
			fmt.Print("ID: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			call, err := repo.GetByID(ctx, id)
			if err != nil {
				fmt.Println("Не найден")
			} else {
				fmt.Printf("ID=%d, Команда=%d, Авто=%d, Район=%v, Статус=%d, Время=%s\n",
					call.ID, call.TeamID, call.CarID, printPtr(call.DistrictID),
					call.StatusID, call.Time.Format("2006-01-02 15:04"))
				if call.Comment != nil {
					fmt.Println("Комментарий:", *call.Comment)
				}
			}
		case "4":
			fmt.Print("ID для обновления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			var call models.Call
			call.ID = id
			fmt.Print("ID команды: ")
			call.TeamID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID автомобиля: ")
			call.CarID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID района (Enter для пропуска): ")
			distStr := readLine(reader)
			if distStr != "" {
				distID, _ := strconv.Atoi(distStr)
				call.DistrictID = &distID
			}
			fmt.Print("ID статуса: ")
			call.StatusID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("Время (YYYY-MM-DD HH:MM): ")
			timeStr := readLine(reader)
			t, _ := time.Parse("2006-01-02 15:04", timeStr)
			call.Time = t
			fmt.Print("Комментарий (Enter для пропуска): ")
			comment := readLine(reader)
			if comment != "" {
				call.Comment = &comment
			}
			if err := repo.Update(ctx, &call); err != nil {
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
