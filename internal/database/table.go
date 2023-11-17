package database

import (
	"database/sql"
	"fmt"
)

func MakeTable(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err.Error())
	}

	CreateUserTable(db)
	CreateCategoryTable(db)
	CreateProductTable(db)
	CreateUserProductTable(db)
	CreateAddressTable(db)
	CreateShippingTable(db)
	CreatePaymentTable(db)
	CreateOrderTable(db)
	CreateOrderItemTable(db)
	CreateCartItemTable(db)
	CreateWishlistItemTable(db)
	CreateReviewTable(db)

	fmt.Println("Tables created")

}

func CreateUserTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS users (
		user_id serial PRIMARY KEY,
		name VARCHAR ( 255 ) NOT NULL,
		phone_number VARCHAR ( 255 ) UNIQUE NOT NULL,
		password VARCHAR ( 255 ) NOT NULL,
		email VARCHAR ( 255 ) UNIQUE NOT NULL,
		role VARCHAR ( 50 ) NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateCategoryTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS categories (
		category_id serial PRIMARY KEY,
		name VARCHAR ( 50 ) NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateProductTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS products (
		product_id serial PRIMARY KEY,
		name VARCHAR ( 50 ) NOT NULL,
		description VARCHAR ( 255 ) NOT NULL,
		price INT NOT NULL,
		category_id INT NOT NULL references categories(category_id) ON DELETE CASCADE,
		image VARCHAR ( 255 ) NOT NULL,
		quantity INT NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateUserProductTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS userproducts (
		userproduct_id serial PRIMARY KEY,
		product_id INT NOT NULL references products(product_id),
		quantity INT NOT NULL,
		order_id INT NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateAddressTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS addresses (
		address_id serial PRIMARY KEY,
		user_id INT NOT NULL references users(user_id),
		street VARCHAR ( 255 ) NOT NULL,
		city VARCHAR ( 50 ) NOT NULL,
		state VARCHAR ( 50 ) NOT NULL,
		country VARCHAR ( 50 ) NOT NULL,
		zip_code VARCHAR ( 50 ) NOT NULL,
		is_default BOOLEAN NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateShippingTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS shippings (
		shipping_id serial PRIMARY KEY,
		user_id INT NOT NULL references users(user_id),
		method VARCHAR ( 50 ) NOT NULL,
		address_id INT NOT NULL references addresses(address_id),
		estimated_delivery_days INT NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreatePaymentTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS payments (
		payment_id serial PRIMARY KEY,
		user_id INT NOT NULL references users(user_id),
		method VARCHAR ( 50 ) NOT NULL,
		status VARCHAR ( 50 ) NOT NULL,
		amount INT NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateOrderTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS orders (
		order_id serial PRIMARY KEY,
		user_id INT NOT NULL references users(user_id),
		total_amount INT NOT NULL,
		payment_id INT NOT NULL references payments(payment_id),
		shipping_id INT NOT NULL references shippings(shipping_id),
		status VARCHAR ( 50 ) NOT NULL,
		order_date TIMESTAMP NOT NULL,
		delivery_date TIMESTAMP NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateOrderItemTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS orderitems (
		orderitem_id serial PRIMARY KEY,
		order_id INT NOT NULL references orders(order_id),
		product_id INT NOT NULL references products(product_id),
		quantity INT NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateCartItemTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS cartitems (
		cartitem_id serial PRIMARY KEY,
		user_id INT NOT NULL references users(user_id),
		product_id INT NOT NULL references products(product_id),
		quantity INT NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateWishlistItemTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS wishlistitems (
		wishlistitem_id serial PRIMARY KEY,
		user_id INT NOT NULL references users(user_id),
		product_id INT NOT NULL references products(product_id)
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func CreateReviewTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS reviews (
		review_id serial PRIMARY KEY,
		user_id INT NOT NULL references users(user_id),
		product_id INT NOT NULL references products(product_id),
		rating INT NOT NULL,
		comment VARCHAR ( 255 ) NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}
}

func DropTable(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err.Error())
	}

	q := `DROP TABLE IF EXISTS users, categories, products, userproducts, addresses, shippings, payments, orders, orderitems, cartitems, wishlistitems, reviews;`

	_, err = db.Exec(q)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Tables dropped")
}
