package cli

import (
	"bufio"
	"context"
	"fmt"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"
)

func handleCallStatus(ctx context.Context, reader *bufio.Reader, repo repository.CallStatusRepository) {
	for {
		fmt.Println("\n--- Статусы вызовов ---")
		fmt.Println("1. Создать  2. Все  3. По ID  4. Обновить  5. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			fmt.Print("Название статуса: ")
			name := readLine(reader)
			id, err := repo.Create(ctx, &models.CallStatus{Name: name})
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Создан статус с ID=%d\n", id)
			}
		case "2":
			list, _ := repo.GetAll(ctx)
			for _, cs := range list {
				fmt.Printf("ID=%d, Название=%s\n", cs.ID, cs.Name)
			}
		case "3":
			fmt.Print("ID: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			cs, err := repo.GetByID(ctx, id)
			if err != nil {
				fmt.Println("Не найдено")
			} else {
				fmt.Printf("ID=%d, Название=%s\n", cs.ID, cs.Name)
			}
		case "4":
			fmt.Print("ID для обновления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			fmt.Print("Новое название: ")
			name := readLine(reader)
			if err := repo.Update(ctx, &models.CallStatus{ID: id, Name: name}); err != nil {
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
