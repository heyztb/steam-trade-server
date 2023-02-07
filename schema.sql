CREATE TABLE Bots (
  id INTEGER PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  passwd TEXT NOT NULL,
  shared_secret TEXT NOT NULL UNIQUE,
  identity_secret TEXT NOT NULL UNIQUE,
  api_key TEXT NOT NULL UNIQUE
);

CREATE TABLE Items (
  id INTEGER PRIMARY KEY,
  bot_id INTEGER NOT NULL,
  app_id INTEGER NOT NULL,
  asset_id INTEGER NOT NULL,
  class_id INTEGER NOT NULL,
  instance_id INTEGER NOT NULL
);