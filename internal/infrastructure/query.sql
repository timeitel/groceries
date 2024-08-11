INSERT OR IGNORE INTO products (name, description)
VALUES ('Apple', 'A delicious fruit.');

INSERT OR IGNORE INTO products (name, description)
VALUES ('Banana', 'Another delicious fruit.');

INSERT OR IGNORE INTO users (name, is_admin)
VALUES ('Cool guy', 1);

INSERT OR IGNORE INTO carts (product_id, user_id)
VALUES (1, 1);
