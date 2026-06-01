package cli

import (
	"bufio"
	"context"
	"fmt"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"
)

func handleTeam(ctx context.Context, reader *bufio.Reader, repo repository.TeamRepository) {
	for {
		fmt.Println("\n--- Команды ---")
		fmt.Println("1. Создать  2. Все  3. По номеру  4. Обновить  5. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			var team models.Team
			fmt.Print("ID специализации: ")
			team.SpecializationID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID района: ")
			team.DistrictID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID автомобиля (Enter для пропуска): ")
			carStr := readLine(reader)
			if carStr != "" {
				carID, _ := strconv.Atoi(carStr)
				team.CarID = &carID
			}
			number, err := repo.Create(ctx, &team)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("Создана команда №%d\n", number)
			}
		case "2":
			list, _ := repo.GetAll(ctx)
			for _, t := range list {
				fmt.Printf("Номер=%d, Специализация=%d, Район=%d, Автомобиль=%v\n",
					t.Number, t.SpecializationID, t.DistrictID, printPtr(t.CarID))
			}
		case "3":
			fmt.Print("Номер команды: ")
			numStr := readLine(reader)
			num, _ := strconv.Atoi(numStr)
			team, err := repo.GetByNumber(ctx, num)
			if err != nil {
				fmt.Println("Не найдена")
			} else {
				fmt.Printf("Номер=%d, Специализация=%d, Район=%d, Автомобиль=%v\n",
					team.Number, team.SpecializationID, team.DistrictID, printPtr(team.CarID))
			}
		case "4":
			fmt.Print("Номер команды для обновления: ")
			numStr := readLine(reader)
			num, _ := strconv.Atoi(numStr)
			var team models.Team
			team.Number = num
			fmt.Print("ID специализации: ")
			team.SpecializationID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID района: ")
			team.DistrictID, _ = strconv.Atoi(readLine(reader))
			fmt.Print("ID автомобиля (Enter для пропуска): ")
			carStr := readLine(reader)
			if carStr != "" {
				carID, _ := strconv.Atoi(carStr)
				team.CarID = &carID
			}
			if err := repo.Update(ctx, &team); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Обновлено")
			}
		case "5":
			fmt.Print("Номер команды для удаления: ")
			numStr := readLine(reader)
			num, _ := strconv.Atoi(numStr)
			if err := repo.Delete(ctx, num); err != nil {
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

func printPtr(p *int) string {
	if p == nil {
		return "не назначен"
	}
	return strconv.Itoa(*p)
}
