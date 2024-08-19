CREATE TABLE items (
    id uuid PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    description text
);

CREATE TABLE users (
    id uuid PRIMARY KEY,
    name varchar NOT NULL,
    is_admin integer DEFAULT 0,
    active_cart_id integer,
    FOREIGN KEY (active_cart_id) REFERENCES carts (id)
);

CREATE TABLE cart_items (
    item_id uuid,
    cart_id uuid,
    quantity integer DEFAULT 1,
    FOREIGN KEY (item_id) REFERENCES items (id),
    FOREIGN KEY (cart_id) REFERENCES carts (id)
);

CREATE TABLE carts (
    id uuid PRIMARY KEY,
    user_id uuid,
    name varchar,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

