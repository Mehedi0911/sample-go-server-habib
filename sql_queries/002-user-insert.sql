	INSERT INTO users (
first_name,
last_name,
email,
password,
	) VALUES (
"Mehedi",
"Hasan",
"mehedi@example.com",
"password123"
    ) RETURNING id, first_name, last_name, email, password, age, money, description, is_shop_owner, created_at, updated_at;