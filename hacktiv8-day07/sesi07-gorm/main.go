package main

import (
	"fmt"
	"hacktiv8-day07/sesi07-gorm/database"
	"hacktiv8-day07/sesi07-gorm/repository"
	"strings"

	"gorm.io/gorm"
)

func main() {
	db := database.StartDB()
	// user(db)
	product(db)

}

func user(db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)

	// //Create users

	// user := models.User{
	// 	Email: "supports",
	// }

	// err := userRepo.CreateUser(&user)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// fmt.Println("Created user success")

	// //Get All users

	// employees, err := userRepo.GetAllUsers()
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// for k, employee := range *employees {
	// 	fmt.Println("User ke :", k+1)
	// 	employee.Print()
	// 	fmt.Println(strings.Repeat("=", 20))
	// }

	// //Get user By Id

	// employee, err := userRepo.GetUserByID(2)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// employee.Print()

	// //Update user

	// emp, err := userRepo.UpdateUserByID(&models.User{
	// 	Email:     "indra@koinworks.com",
	// 	UpdatedAt: time.Now(),
	// }, 1)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// emp.Print()

	//Delete user

	// err := userRepo.DeleteUserByID(3)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// } else {
	// 	fmt.Println("user success deleted")
	// }

	// // Get Users With Products
	employees, err := userRepo.GetUsersWithProducts()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	// employees.Print()

	for k, employee := range *employees {
		fmt.Println("User ke :", k+1)
		employee.Print()
		fmt.Println(strings.Repeat("=", 20))
	}
}

func product(db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)

	// //Create product

	// product := models.Product{
	// 	Name:   "Celana",
	// 	Brand:  "Matahari",
	// 	UserID: 2,
	// }
	// err := productRepository.CreateProduct(&product)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Create Product success")

	// // //Get All Products

	// products, err := productRepository.GetAllProducts()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// for k, p := range *products {
	// 	fmt.Println("Product ke :", k+1)
	// 	p.Print()
	// 	fmt.Println(strings.Repeat("=", 20))
	// }

	//Get product By Id

	// p, err := productRepository.GetProductByID(5)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// p.Print()

	//Update product

	// pr, err := productRepository.UpdateProductByID(&models.Product{
	// 	Name:      "Jeans",
	// 	Brand:     "Levis",
	// 	UpdatedAt: time.Now(),
	// }, 5)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("Product success updated")
	// pr.Print()

	//Delete user

	err := productRepository.DeleteProductByID(5)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	} else {
		fmt.Println("Product success deleted")
	}
}
