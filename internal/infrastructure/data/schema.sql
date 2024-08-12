CREATE TABLE products (
    id integer PRIMARY KEY AUTOINCREMENT,
    name varchar NOT NULL UNIQUE,
    description text
);

CREATE TABLE users (
    id integer PRIMARY KEY AUTOINCREMENT,
    name varchar NOT NULL,
    is_admin integer DEFAULT 0
);

CREATE TABLE cart_items (
    product_id integer,
    cart_id integer,
    quantity integer DEFAULT 1,
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (cart_id) REFERENCES carts (id)
);

CREATE TABLE carts (
    id integer PRIMARY KEY AUTOINCREMENT,
    user_id integer,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

