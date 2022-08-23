
-- @block
CREATE TABLE Users(
    id INT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    bio TEXT,
    country VARCHAR(2)
)

-- 