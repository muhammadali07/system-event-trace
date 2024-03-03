-- Insert a new user into the database
INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email;

-- Retrieve a user by ID
SELECT id, name, email FROM users WHERE id = $1;

-- Update a user's information
UPDATE users SET name = $2, email = $3 WHERE id = $1 RETURNING id, name, email;

-- Delete a user by ID
DELETE FROM users WHERE id = $1 RETURNING id, name, email;
