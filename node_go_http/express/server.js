import express from 'express';
import { readFile } from 'fs/promises';
import { join } from 'path';

const PORT = 3000;

async function run() {
  const path = join('..', 'common', 'json', `${process.argv[2]}.json`);
  const data = await readFile(path);
  const json = JSON.parse(data);

  const app = express()

  app.get('/json-stringify', function (req, res) {
    res.send(json)
  })

  app.listen(PORT, () => { console.log(`up on ${PORT}`); })
}

run()