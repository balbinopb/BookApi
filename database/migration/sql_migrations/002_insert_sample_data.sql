-- +migrate Up
INSERT INTO categories (name, created_by)
VALUES
  ('Fiction', 'system'),
  ('Science', 'system'),
  ('Technology', 'system'),
  ('History', 'system');

INSERT INTO users (username, password, created_by)
VALUES
  -- Note: password here is plain text for testing, in production use hashing (bcrypt)
  ('admin', 'admin123', 'system'),
  ('john_doe', 'password123', 'system');

INSERT INTO books (
    title,
    category_id,
    description,
    image_url,
    release_year,
    price,
    total_page,
    thickness,
    created_by
)
VALUES
  (
    'The Great Novel',
    1,
    'An epic fictional story about courage and hope.',
    'https://example.com/images/great_novel.jpg',
    2020,
    19.99,
    320,
    2.5,
    'system'
  ),
  (
    'The Science of Everything',
    2,
    'Explores the mysteries of the universe in simple terms.',
    'https://example.com/images/science_everything.jpg',
    2019,
    25.50,
    410,
    3.1,
    'system'
  ),
  (
    'Learn Go Programming',
    3,
    'A practical guide to programming in Go for beginners and professionals.',
    'https://example.com/images/learn_go.jpg',
    2023,
    29.99,
    550,
    4.0,
    'system'
  );

-- +migrate Down
DELETE FROM books WHERE title IN ('The Great Novel', 'The Science of Everything', 'Learn Go Programming');
DELETE FROM users WHERE username IN ('admin', 'john_doe');
DELETE FROM categories WHERE name IN ('Fiction', 'Science', 'Technology', 'History');
