CREATE TABLE Items (
  id BIGSERIAL PRIMARY KEY,
  bot_id BIGINT NOT NULL,
  app_id BIGINT NOT NULL,
  asset_id BIGINT NOT NULL,
  class_id BIGINT NOT NULL,
  instance_id BIGINT NOT NULL
);

CREATE TABLE Bots (
  id BIGSERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  passwd TEXT NOT NULL,
  shared_secret TEXT NOT NULL,
  identity_secret TEXT NOT NULL
);

CREATE INDEX app_id_idx ON Inventory (
  app_id 
);

CREATE INDEX asset_id_idx ON Inventory (
  asset_id 
);