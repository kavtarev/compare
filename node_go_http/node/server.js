import { createServer } from 'http';
import { readFile } from 'fs/promises';
import { join } from 'path';
import { createReadStream } from 'fs';
import { XMLParser } from 'fast-xml-parser';

const PORT = 3000;

async function run() {
  const jsonPath = join('..', 'common', 'json', `${process.argv[2]}.json`);
  const jsonData = await readFile(jsonPath);
  const json = JSON.parse(jsonData);

  const filePath = join('..', 'common', 'files', `${process.argv[2]}.txt`);

  const parser = new XMLParser();

  const xmlPath = join('..', 'common', 'xml', `${process.argv[2]}.xml`);
  const xmlData = await readFile(xmlPath);

  createServer((req, res) => {
    if (req.url === '/json-stringify') {
      res.end(JSON.stringify(json));
    }
    if (req.url === '/download-file') {
      createReadStream(filePath, { highWaterMark: 50000 }).pipe(res);
    }
    if (req.url === '/parse-xml') {
      const xml = parser.parse(xmlData);
      res.end(JSON.stringify(xml))
    }

  }).listen(PORT, () => { console.log(`up on ${PORT}`) })

}

run()