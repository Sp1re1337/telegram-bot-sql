package main

import (
	"log"
    "telegram-bot-sql-example/bot"
    "telegram-bot-sql-example/database"
)

func main() {
    // Ініціалізуємо базу даних
    err := database.InitDB()
    if err != nil {
        log.Fatalf("Помилка ініціалізації бази даних: %v", err)
    }

<<<<<<< HEAD
    // Запускаємо Telegram бота
    err = bot.StartBot()
    if err != nil {
        log.Fatalf("Помилка запуску бота: %v", err)
    }
}
=======
	database.CreateTable(db)

	database.AddUser(db, "Вася", 28)
  database.AddUser(db, "Коля", 24)

	users, err := database.GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Список користувачів:")
	for _, user := range users {
		fmt.Printf("ID: %d, Ім'я: %s, Вік: %d\n", user.ID, user.Name, user.Age)
	}

	database.DeleteUser(db, 1)

	users, err = database.GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Оновлений список користувачів:")
	for _, user := range users {
		fmt.Printf("ID: %d, Ім'я: %s, Вік: %d\n", user.ID, user.Name, user.Age)
	}
}
>>>>>>> 14870955344d5ae401485de3c994b08cc46353fe
