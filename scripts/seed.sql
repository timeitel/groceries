
  CREATE TABLE IF NOT EXISTS items (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR NOT NULL UNIQUE,
  description TEXT
  );

  INSERT OR IGNORE INTO items (name, description)
  VALUES ('Apple', 'A delicious fruit.');

  INSERT OR IGNORE INTO items (name, description)
  VALUES ('Banana', 'Another delicious fruit.');
