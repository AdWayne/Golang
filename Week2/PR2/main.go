package main

import "fmt"

func main() {
	manager := UserManager{}

	u1 := User{"Azamat", "azamat@mail.com", "Admin"}
	u2 := User{"Ali", "ali@mail.com", "User"}

	manager.AddUser(u1)
	manager.AddUser(u2)

	fmt.Println("После добавления:", manager.users)

	manager.UpdateUser("ali@mail.com", User{"Ali", "ali@mail.com", "Admin"})
	fmt.Println("После обновления:", manager.users)

	manager.RemoveUser("azamat@mail.com")
	fmt.Println("После удаления:", manager.users)
}
