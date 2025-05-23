const { request, createServer } = require('http');
const protobuf = require('protobufjs');
const fs = require('fs/promises');
const path = require('path');

async function run() {
  const { 2: type, 3: level, 4: amount } = process.argv;

  const root = await protobuf.load('example.proto')
  const buffer = root.lookupType(level)

  const jsonData = await fs.readFile(path.join('..', 'common', 'json', `${level}.json`))
  const jsonObject = JSON.parse(jsonData)

  createServer((req, res) => {
    const data = [];
    req.on('data', chunk => { data.push(chunk); });

    if (req.url === '/json') {
      req.on('end', () => { JSON.parse(Buffer.concat(data)); });
    }
    if (req.url === '/proto') {
      req.on('end', () => { buffer.decode(Buffer.concat(data)); });
    }
    res.end('end')

  }).listen(3000, () => { console.log('Server listening on port 3000') });


  function makeRequest(type) {
    if (type === 'json') {
      return new Promise((resolve, reject) => {
        const backReq = request(
          { host: '127.0.0.1', port: 3000, method: 'POST', path: '/json' },
          res => {
            const data = [];
            res.on('data', chunk => { data.push(chunk); });
            res.on('end', resolve);
            res.on('error', reject);
          }
        );

        backReq.write(JSON.stringify(jsonObject));
        backReq.end();
      });
    }
    if (type === 'proto') {
      return new Promise((resolve, reject) => {
        const json = buffer.encode(jsonObject).finish()

        const backReq = request(
          { host: '127.0.0.1', port: 3000, method: 'POST', path: '/proto', headers: { 'content-type': 'application/octet-stream', "content-length": json.length } },
          res => {
            const data = [];
            res.on('data', chunk => { data.push(chunk); });
            res.on('end', resolve);
            res.on('error', reject);
          }
        );

        backReq.write(json);
        backReq.end();
      });
    }
  }

  async function runSequentialRequests() {
    console.time(`${amount} of ${type} took`)
    for (let i = 0; i < amount; i++) {
      await makeRequest(type);
    }
    console.timeEnd(`${amount} of ${type} took`)
  };

  await runSequentialRequests()
}

run()