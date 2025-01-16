import { Pool } from './pool';

export async function InitializeTables() {
  await Pool.query(`
    create table if not exists users (
      id int generated always as identity
      , name text
      , status text
      , level int
      , primary key (id)
    );
  `)

  await Pool.query(`
    create table if not exists tasks (
      id int generated always as identity
      , name text
      , status text
      , amount int
      , implementer_id int
      , primary key (id)
      , constraint fk_user foreign key(implementer_id) references users(id)
    );
  `)
}