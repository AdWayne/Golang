package main

import "fmt"

type Library struct {
	Books   []Book
	Readers []Reader
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
	fmt.Println("Книга добавлена:", book.Title)
}

func (l *Library) RemoveBook(isbn string) {
	for i, b := range l.Books {
		if b.ISBN == isbn {
			fmt.Println("Книга удалена:", b.Title)
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return
		}
	}
	fmt.Println("Книга не найдена")
}

func (l *Library) RegisterReader(reader Reader) {
	l.Readers = append(l.Readers, reader)
	fmt.Println("Читатель зарегистрирован:", reader.Name)
}

func (l *Library) RemoveReader(id string) {
	for i, r := range l.Readers {
		if r.ID == id {
			fmt.Println("Читатель удалён:", r.Name)
			l.Readers = append(l.Readers[:i], l.Readers[i+1:]...)
			return
		}
	}
	fmt.Println("Читатель не найден")
}

func (l *Library) BorrowBook(isbn string) {
	for i := range l.Books {
		if l.Books[i].ISBN == isbn {
			if l.Books[i].Copies > 0 {
				l.Books[i].Copies--
				fmt.Println("Книга выдана:", l.Books[i].Title)
			} else {
				fmt.Println("Нет доступных экземпляров:", l.Books[i].Title)
			}
			return
		}
	}
	fmt.Println("Книга не найдена")
}

func (l *Library) ReturnBook(isbn string) {
	for i := range l.Books {
		if l.Books[i].ISBN == isbn {
			l.Books[i].Copies++
			fmt.Println("Книга возвращена:", l.Books[i].Title)
			return
		}
	}
	fmt.Println("Книга не найдена")
}

func (l *Library) ListBooks() {
	fmt.Println("\nСписок книг:")
	for _, b := range l.Books {
		fmt.Printf(" - %s | %s | %s | Экз: %d\n", b.Title, b.Author, b.ISBN, b.Copies)
	}
}

func (l *Library) ListReaders() {
	fmt.Println("\nСписок читателей:")
	for _, r := range l.Readers {
		fmt.Printf(" - %s (ID: %s)\n", r.Name, r.ID)
	}
}
