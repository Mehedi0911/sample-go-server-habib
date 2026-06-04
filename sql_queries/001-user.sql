CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    age INT NOT NULL,
    money DOUBLE PRECISION NOT NULL,
    description TEXT,
    is_shop_owner BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- data types 

-- SERIAL - 32 bits
-- BIGSERIAL - 64 bits

-- SMALLINT - 16 bits
-- INT - 32 bits
-- BIGINT - 64 bits

-- REAL - 32 bits
-- DOUBLE PRECISION - 64 bits - more precision than REAL after decimal point

-- CHAR(n) - fixed length string with a maximum of n characters, waste of memory if the string is shorter than n
-- VARCHAR(255) - variable length string with a maximum of 255 characters
-- TEXT - variable length string with no specific maximum length

-- BOOLEAN - true or false

TIME - timestamp without time zone, stores the time of day (hour, minute, second) without any date or time zone information. example: '14:30:00'
TIMESTAMPTZ - timestamp with time zone, stores both the date and time along with time zone information. example: '2024-06-01 14:30:00+00'
DATE - stores only the date (year, month, day) without any time information. example: '2024-06-01'
TIMESTAMP WITH TIME ZONE - stores both the date and time along with time zone information, similar to TIMESTAMPTZ. example: '2024-06-01 14:30:00+00'