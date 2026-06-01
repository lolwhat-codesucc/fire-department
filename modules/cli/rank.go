package cli

import (
	"bufio"
	"context"
	"fmt"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"
)

func handleRank(ctx context.Context, reader *bufio.Reader, repo repository.RankRepository) {
	for {
		fmt.Println("\n--- Звания ---")
		fmt.Println("1. Создать  2. Все  3. По ID  4. Обновить  5. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			fmt.Print("Введите название: ")
			name := readLine(reader)
			id, err := repo.Create(ctx, &models.Rank{Name: name})
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Создано звание с ID=%d\n", id)
			}
		case "2":
			list, err := repo.GetAll(ctx)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}
			for _, r := range list {
				fmt.Printf("ID=%d, Название=%s\n", r.ID, r.Name)
			}
		case "3":
			fmt.Print("Введите ID: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			rank, err := repo.GetByID(ctx, id)
			if err != nil {
				fmt.Println("Не найдено")
			} else {
				fmt.Printf("ID=%d, Название=%s\n", rank.ID, rank.Name)
			}
		case "4":
			fmt.Print("ID для обновления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			fmt.Print("Новое название: ")
			name := readLine(reader)
			if err := repo.Update(ctx, &models.Rank{ID: id, Name: name}); err != nil {
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
