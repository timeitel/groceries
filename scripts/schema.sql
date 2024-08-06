CREATE TABLE IF NOT EXISTS items (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR NOT NULL UNIQUE,
  description TEXT
);

CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS shopping_lists (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  item_id INTEGER,
  user_id INTEGER,
  FOREIGN KEY (item_id) REFERENCES items(item_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

