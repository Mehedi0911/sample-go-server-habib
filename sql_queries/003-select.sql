SELECT * FROM users;

SELECT id, first_name, last_name FROM users;

SELECT * FROM users WHERE email = '$1';
SELECT * FROM users WHERE email = '$1' AND password = '$2';