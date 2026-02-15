package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	library := Library{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== МЕНЮ БИБЛИОТЕКИ =====")
		fmt.Println("1. Добавить книгу")
		fmt.Println("2. Удалить книгу")
		fmt.Println("3. Зарегистрировать читателя")
		fmt.Println("4. Удалить читателя")
		fmt.Println("5. Выдать книгу")
		fmt.Println("6. Вернуть книгу")
		fmt.Println("7. Показать книги")
		fmt.Println("8. Показать читателей")
		fmt.Println("0. Выход")
		fmt.Print("Выберите пункт: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			fmt.Print("Название: ")
			title, _ := reader.ReadString('\n')

			fmt.Print("Автор: ")
			author, _ := reader.ReadString('\n')

			fmt.Print("ISBN: ")
			isbn, _ := reader.ReadString('\n')

			fmt.Print("Количество экземпляров: ")
			copiesStr, _ := reader.ReadString('\n')
			copies, _ := strconv.Atoi(strings.TrimSpace(copiesStr))

			book := Book{
				Title:  strings.TrimSpace(title),
				Author: strings.TrimSpace(author),
				ISBN:   strings.TrimSpace(isbn),
				Copies: copies,
			}
			library.AddBook(book)

		case 2:
			fmt.Print("Введите ISBN книги: ")
			isbn, _ := reader.ReadString('\n')
			library.RemoveBook(strings.TrimSpace(isbn))

		case 3:
			fmt.Print("Имя читателя: ")
			name, _ := reader.ReadString('\n')

			fmt.Print("ID читателя: ")
			id, _ := reader.ReadString('\n')

			reader := Reader{
				Name: strings.TrimSpace(name),
				ID:   strings.TrimSpace(id),
			}
			library.RegisterReader(reader)

		case 4:
			fmt.Print("Введите ID читателя: ")
			id, _ := reader.ReadString('\n')
			library.RemoveReader(strings.TrimSpace(id))

		case 5:
			fmt.Print("Введите ISBN книги: ")
			isbn, _ := reader.ReadString('\n')
			library.BorrowBook(strings.TrimSpace(isbn))

		case 6:
			fmt.Print("Введите ISBN книги: ")
			isbn, _ := reader.ReadString('\n')
			library.ReturnBook(strings.TrimSpace(isbn))

		case 7:
			library.ListBooks()

		case 8:
			library.ListReaders()

		case 0:
			fmt.Println("Выход из программы...")
			return

		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
