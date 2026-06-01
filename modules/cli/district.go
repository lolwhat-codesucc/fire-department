package cli

import (
	"bufio"
	"context"
	"fmt"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"
)

func handleDistrict(ctx context.Context, reader *bufio.Reader,
	districtRepo repository.DistrictRepository,
	specRepo repository.SpecializationRepository) {
	for {
		fmt.Println("\n--- Районы ---")
		fmt.Println("1. Создать  2. Все  3. По ID  4. Обновить  5. Удалить")
		fmt.Println("6. Управление специализациями  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			fmt.Print("Название: ")
			name := readLine(reader)
			id, err := districtRepo.Create(ctx, &models.District{Name: name})
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Создан район ID:", id)
			}
		case "2":
			list, _ := districtRepo.GetAll(ctx)
			for _, d := range list {
				fmt.Printf("ID=%d, Название=%s\n", d.ID, d.Name)
			}
		case "3":
			fmt.Print("ID: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			d, err := districtRepo.GetByID(ctx, id)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Printf("ID=%d, %s\n", d.ID, d.Name)
			}
		case "4":
			fmt.Print("ID для обновления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			fmt.Print("Новое название: ")
			name := readLine(reader)
			if err := districtRepo.Update(ctx, &models.District{ID: id, Name: name}); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Обновлено")
			}
		case "5":
			fmt.Print("ID для удаления: ")
			idStr := readLine(reader)
			id, _ := strconv.Atoi(idStr)
			if err := districtRepo.Delete(ctx, id); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Удалено")
			}
		case "6":
			fmt.Print("ID района: ")
			distIDStr := readLine(reader)
			distID, _ := strconv.Atoi(distIDStr)
			handleDistrictSpecializations(ctx, reader, districtRepo, specRepo, distID)
		case "0":
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}

func handleDistrictSpecializations(ctx context.Context, reader *bufio.Reader,
	districtRepo repository.DistrictRepository,
	specRepo repository.SpecializationRepository,
	districtID int) {
	for {
		fmt.Println("\nСпециализации района")
		fmt.Println("1. Показать  2. Добавить  3. Удалить  0. Назад")
		fmt.Print("> ")
		choice := readLine(reader)
		switch choice {
		case "1":
			specs, err := districtRepo.GetSpecializations(ctx, districtID)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				for _, s := range specs {
					fmt.Printf("ID=%d, %s\n", s.ID, s.Name)
				}
			}
		case "2":
			fmt.Print("ID специализации для добавления: ")
			specIDStr := readLine(reader)
			specID, _ := strconv.Atoi(specIDStr)
			if err := districtRepo.AddSpecialization(ctx, districtID, specID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Добавлено")
			}
		case "3":
			fmt.Print("ID специализации для удаления: ")
			specIDStr := readLine(reader)
			specID, _ := strconv.Atoi(specIDStr)
			if err := districtRepo.RemoveSpecialization(ctx, districtID, specID); err != nil {
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
