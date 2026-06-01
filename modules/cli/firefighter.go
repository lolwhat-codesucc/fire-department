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

func handleFirefighter(ctx context.Context, reader *bufio.Reader, repo repository.FirefighterRepository) {
	for {
		fmt.Println("\n--- Пожарные ---")
		fmt.Println("1. Создать  2. Все  3. По ID  4. Обновить  5. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			var ff models.Firefighter
			fmt.Print("Имя: ")
			ff.Name = readLine(reader)
			fmt.Print("Дата рождения (YYYY-MM-DD): ")
			birthStr := readLine(reader)
			t, _ := time.Parse("2006-01-02", birthStr)
			ff.YearOfBirth = models.Date{Time: t}
			fmt.Print("ID звания: ")
			ff.RankID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("Квалификация (число): ")
			ff.Qualification, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID команды (Enter для пропуска): ")
			teamStr := readLine(reader)
			if teamStr != "" {
				teamID, _ := strconv.Atoi(teamStr)
				ff.TeamID = &teamID
			}
			id, err := repo.Create(ctx, &ff)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Создан пожарный с ID=%d\n", id)
			}
		case "2":
			list, _ := repo.GetAll(ctx)
			for _, f := range list {
				teamStr := "не назначен"
				if f.TeamID != nil {
					teamStr = strconv.Itoa(*f.TeamID)
				}
				fmt.Printf("ID=%d, Имя=%s, Звание=%d, Команда=%s\n", f.ID, f.Name, f.RankID, teamStr)
			}
		case "3":
			fmt.Print("ID: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			ff, err := repo.GetByID(ctx, id)
			if err != nil {
				fmt.Println("Не найден")
			} else {
				teamStr := "не назначен"
				if ff.TeamID != nil {
					teamStr = strconv.Itoa(*ff.TeamID)
				}
				fmt.Printf("ID=%d, Имя=%s, Родился=%s, Звание=%d, Квалификация=%d, Команда=%s\n",
					ff.ID, ff.Name, ff.YearOfBirth.Time.Format("2006-01-02"),
					ff.RankID, ff.Qualification, teamStr)
			}
		case "4":
			fmt.Print("ID для обновления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			var ff models.Firefighter
			ff.ID = id
			fmt.Print("Имя: ")
			ff.Name = readLine(reader)
			fmt.Print("Дата рождения (YYYY-MM-DD): ")
			birthStr := readLine(reader)
			t, _ := time.Parse("2006-01-02", birthStr)
			ff.YearOfBirth = models.Date{Time: t}
			fmt.Print("ID звания: ")
			ff.RankID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("Квалификация: ")
			ff.Qualification, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID команды (Enter для пропуска): ")
			teamStr := readLine(reader)
			if teamStr != "" {
				teamID, _ := strconv.Atoi(teamStr)
				ff.TeamID = &teamID
			}
			if err := repo.Update(ctx, &ff); err != nil {
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
