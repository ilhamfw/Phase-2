package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

var db *sql.DB

func main() {
	// Memuat variabel lingkungan dari file .env
	loadEnv()

	// Baca konfigurasi dari variabel lingkungan
	dbUser := os.Getenv("DBUser") // Menggunakan "DBUser" sesuai dengan nama variabel lingkungan
	
	dbHost := os.Getenv("DBHost") // Menggunakan "DBHost" sesuai dengan nama variabel lingkungan
	dbPort := os.Getenv("DBPort") // Menggunakan "DBPort" sesuai dengan nama variabel lingkungan
	dbName := os.Getenv("DBName") // Menggunakan "DBName" sesuai dengan nama variabel lingkungan

	// Kini Anda dapat menggunakan variabel-variabel ini dalam koneksi database Anda
	dbURI := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	router := httprouter.New()

	// Middleware untuk logging request
	router.GET("/*path", LogRequest)

	// Register endpoint untuk Register and Login
	router.POST("/register", Register)
	router.POST("/login", Login)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func LogRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Logging request dengan format yang diinginkan
	log.Printf("%s - HTTP request sent to %s %s\n", time.Now().Format("2006/01/02 15:04:05"), r.Method, r.URL.Path)
}

type User struct {
	Email      string
	Password   string
	FullName   string
	Age        int
	Occupation string
	Role       string
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ageStr := r.FormValue("age")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		http.Error(w, "Invalid age format", http.StatusBadRequest)
		return
	}

	// Menerima data pengguna dari permintaan HTTP
	user := User{
		Email:      r.FormValue("email"),
		Password:   r.FormValue("password"),
		FullName:   r.FormValue("full_name"),
		Age:        age, // Menggunakan nilai age yang sudah dikonversi
		Occupation: r.FormValue("occupation"),
		Role:       r.FormValue("role"),
	}

	//Validasi data pengguna
	if err := ValidateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password sebelum menyimpannya
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
		return
	}

	// Simpan data pengguna ke database
	_, err = db.Exec("INSERT INTO users (email, password, full_name, age, occupation, role) VALUES (?,?,?,?,?,?)",
		user.Email, hashedPassword, user.FullName, user.Age, user.Occupation, user.Role)
	if err != nil {
		http.Error(w, "Failed to register the user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User registered successfully")
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Implementasi login akan ditambahkan di sini
}

func ValidateUser(user User) error {
	// Lakukan validasi sesuai dengan ketentuan yang telah diberikan
	// Anda dapat menggunakan pustaka pihak ketiga seperti "validator" atau
	// mengimplementasikannya secara manual.

	// Contoh validasi email
	if user.Email == "" {
		return fmt.Errorf("email is required")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return fmt.Errorf("invalid email format")
	}

	// Implementasikan validasi lainnya sesuai dengan kebutuhan

	return nil
}
