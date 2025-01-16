import { Pool } from './pool';

/**
 * @param {{users: number, tasks: number}} props task < users !!!
 */
export async function InitializeTables(props) {
  await Pool.query(`
    insert into users (name, status, level)
    select
      md5(random()::text)
      , case
          when random() > 0.2 then 'active'
          else 'non_active'
        end 
      , case
          when random() > 0.5 then 1
          else 2
        end
      from generate_series(0, ${props.users}); 
  `)

  await Pool.query(`
    insert into tasks (name, status, amount, implementer_id)
    select
      md5(random()::text)
      , case
          when random() > 0.3 then 'in_progress'
          else 'done'
        end 
      , (random() * 1000)::int
      , null 
    from generate_series(0, 5); 
  `)

  await Pool.query(`
    update tasks
    set implementer_id =
      case
        when random() > 0.8 then null
        else id
      end;
  `)
}