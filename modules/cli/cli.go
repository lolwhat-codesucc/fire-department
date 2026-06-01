package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"fire-department/modules/postgres"
)

func Run(pool *pgxpool.Pool) {
	rankRepo := postgres.NewRankRepo(pool)
	specRepo := postgres.NewSpecializationRepo(pool)
	districtRepo := postgres.NewDistrictRepo(pool)
	carModelRepo := postgres.NewCarModelRepo(pool)
	callStatusRepo := postgres.NewCallStatusRepo(pool)
	carRepo := postgres.NewCarRepo(pool)
	teamRepo := postgres.NewTeamRepo(pool)
	ffRepo := postgres.NewFirefighterRepo(pool)
	callRepo := postgres.NewCallRepo(pool)

	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Println("\n=== Главное меню ===")
		fmt.Println("1. Звания")
		fmt.Println("2. Специализации")
		fmt.Println("3. Районы")
		fmt.Println("4. Модели автомобилей")
		fmt.Println("5. Статусы вызовов")
		fmt.Println("6. Автомобили")
		fmt.Println("7. Команды")
		fmt.Println("8. Пожарные")
		fmt.Println("9. Вызовы")
		fmt.Println("0. Выход")
		fmt.Print("Выберите раздел: ")

		choice := readLine(reader)
		switch choice {
		case "1":
			handleRank(ctx, reader, rankRepo)
		case "2":
			handleSpecialization(ctx, reader, specRepo)
		case "3":
			handleDistrict(ctx, reader, districtRepo)
		case "4":
			handleCarModel(ctx, reader, carModelRepo)
		case "5":
			handleCallStatus(ctx, reader, callStatusRepo)
		case "6":
			handleCar(ctx, reader, carRepo)
		case "7":
			handleTeam(ctx, reader, teamRepo)
		case "8":
			handleFirefighter(ctx, reader, ffRepo)
		case "9":
			handleCall(ctx, reader, callRepo)
		case "0":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
