CREATE TABLE items (
    id integer PRIMARY KEY AUTOINCREMENT,
    name varchar NOT NULL UNIQUE,
    description text
);

CREATE TABLE users (
    id uuid PRIMARY KEY,
    name varchar NOT NULL,
    is_admin boolean DEFAULT 0
);

CREATE TABLE carts (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    name varchar,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE cart_items (
    id uuid NOT NULL,
    item_id integer NOT NULL,
    cart_id uuid NOT NULL,
    quantity integer DEFAULT 1,
    FOREIGN KEY (item_id) REFERENCES items (id),
    FOREIGN KEY (cart_id) REFERENCES carts (id)
);

