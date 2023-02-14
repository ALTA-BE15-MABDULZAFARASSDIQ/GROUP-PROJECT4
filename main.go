package main

/*
jalankan:
	go mod init namaproject
download mysql db driver:
	go get -u github.com/go-sql-driver/mysql
*/

import (
	"be15/gp4/controllers"
	"be15/gp4/entities"
	"fmt"
	"be15/gp4/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := config.ConnectToDB()
	defer db.Close()

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Add Account (register)")
		fmt.Println("2. Login")
		fmt.Println("3. Read Account")
		fmt.Println("4. Update Account")
		fmt.Println("5. Delete Account")
		fmt.Println("6. Top-up")
		fmt.Println("7. Transfer")
		fmt.Println("8. History Top-up")
		fmt.Println("9. History Transfer")
		fmt.Println("10. View Other User Profile")
		fmt.Println("0. Exit")
		fmt.Print("Input: ")
		
		var input int
		fmt.Scanln(&input)
		
		fmt.Scan(&input)

		switch input {
		case 1:
			// Add Account (register)
			func register(username string, password string, email string) error {
				// logic untuk melakukan registrasi user di sini
				return nil // nilai kembalian error akan berisi pesan kesalahan jika ada kesalahan saat registrasi
			}
		case 2:
			// Login
			
		case 3:
			// Read Account
			case 4: 
            // Update Account 
			case 5: 
            // Delete Account 
			case 6: 
            // Top-up 
			entities.Topup
	case 7: 
            // Transfer 
			entities.Transfer
    case 8: 
            // History Top-up 
    case 9: 
            // History Transfer 
    case 10: 
            // View Other User Profile 
    case 0:
      fmt.Println("Terima kasih telah bertransaksi")
      return
    default:
      fmt.Println("Input salah")
	  fmt.Println("update")
	  fmt.Println("delete")
    }
  }
}