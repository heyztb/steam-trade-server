CREATE TABLE Inventory (
  id BIGSERIAL PRIMARY KEY,
  bot_id BIGINT NOT NULL,
  app_id BIGINT NOT NULL,
  asset_id BIGINT NOT NULL,
  class_id BIGINT NOT NULL,
  instance_id BIGINT NOT NULL
);

CREATE INDEX app_id_idx ON Inventory (
  app_id 
);

CREATE INDEX asset_id_idx ON Inventory (
  asset_id 
);