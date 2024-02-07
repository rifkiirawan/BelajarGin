package main

import (
	"belajar-gin/database"
	"belajar-gin/routers"
	"fmt"
)

// func main() {
// 	database.StartDB()

// 	// createUser("jondoe@mail.com")
// 	// updateUserById(1, "jongjong@mail.com")
// 	// createProduct(1, "YLO", "YYYY")
// 	// getUsersWithProducts()
// 	// deleteProductById(1)
// }

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("error creating user data:", err)
		return
	}

	fmt.Println("new user data:", User)
}

// func updateUserById(id uint, email string) {
// 	db := database.GetDB()

// 	user := models.User{}

// 	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

// 	if err != nil {
// 		fmt.Println("error updating user data:", err)
// 		return
// 	}

// 	fmt.Printf("update user's email: %+v \n", user.Email)
// }

// func createProduct(userID uint, brand string, name string) {
// 	db := database.GetDB()

// 	Product := models.Product{
// 		UserID: userID,
// 		Brand:  brand,
// 		Name:   name,
// 	}

// 	err := db.Create(&Product).Error

// 	if err != nil {
// 		fmt.Println("error creating product data:", err.Error())
// 		return
// 	}

// 	fmt.Println("New Product Data:", Product)
// }

// func deleteProductById(id uint) {
// 	db := database.GetDB()

// 	product := models.Product{}
// 	err := db.Where("id = ?", id).Delete(&product).Error

// 	if err != nil {
// 		fmt.Println("error deleting product:", err.Error())
// 		return
// 	}

// 	fmt.Printf("product with id %d has been successfully deleted", id)
// }

// func getUsersWithProducts() {
// 	db := database.GetDB()

// 	users := models.User{}
// 	err := db.Preload("Products").Find(&users).Error

// 	if err != nil {
// 		fmt.Println("Error getting user datas with products:", err.Error())
// 		return
// 	}

// 	fmt.Println("User datas with producs")
// 	fmt.Printf("%+v", users)
// }

//UNTUK GIN

func main() {
	var PORT = ":8080"
	database.StartDB()

	routers.StartServer().Run(PORT)

}
