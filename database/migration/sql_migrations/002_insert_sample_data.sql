-- +migrate Up
-- Insert sample categories
INSERT INTO categories (name, created_by)
VALUES 
  ('Fiction', 'seed'),
  ('Science', 'seed'),
  ('Programming', 'seed');

-- Insert sample users
INSERT INTO users (username, password, created_by)
VALUES 
  ('admin', 'password123', 'seed'),
  ('user1', 'password123', 'seed');

-- Insert sample books
INSERT INTO books (title, category_id, description, image_url, release_year, price, total_page, thickness, created_by)
VALUES
  ('The Go Programming Language', 3, 'Comprehensive guide to Go', 'https://example.com/go.jpg', 2015, 500.00, 300, 'tebal', 'seed'),
  ('A Brief History of Time', 2, 'Stephen Hawking classic', 'https://example.com/time.jpg', 1988, 250.00, 200, 'tebal', 'seed'),
  ('Short Stories for Kids', 1, 'Collection of bedtime stories', 'https://example.com/stories.jpg', 2020, 50.00, 80, 'tipis', 'seed');

-- +migrate Down
DELETE FROM books WHERE created_by = 'seed';
DELETE FROM categories WHERE created_by = 'seed';
DELETE FROM users WHERE created_by = 'seed';
