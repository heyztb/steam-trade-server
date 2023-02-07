-- +goose Up
-- +goose StatementBegin
create table bots (
  id integer primary key,
  username text not null,
  passwd text not null,
  shared_secret text not null,
  identity_secret text not null,
  api_key text not null
);

create table items (
  id integer primary key,
  bot_id integer not null,
  app_id integer not null,
  asset_id integer not null,
  class_id integer not null,
  instance_id integer not null,
  foreign key(bot_id) references bots(id)
);

create unique index idx_bots_username on bots(username);
create unique index idx_bots_shared_secret on bots(shared_secret);
create unique index idx_bots_identity_secret on bots(identity_secret);
create unique index idx_bots_api_key on bots(api_key);

create unique index idx_unique_item on items(app_id, asset_id);
create index idx_items_app_id on items(app_id);
create index idx_items_asset_id on items(asset_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index idx_bots_username;
drop index idx_bots_shared_secret;
drop index idx_bots_identity_secret;
drop index idx_bots_api_key;
drop index idx_unique_item;
drop index idx_items_app_id;
drop index idx_items_asset_id;

drop table items;
drop table bots;
-- +goose StatementEnd