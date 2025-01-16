import pg from 'pg';

export const Pool = new pg.Pool({
  host: 'localhost',
  port: 5433,
  password: 'postgres',
  database: 'compare_node_go',
  user: 'postgres'
});

