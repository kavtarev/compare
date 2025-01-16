const protobuf = require('protobufjs');
const fs = require('node:fs/promises');
const path = require('node:path');

async function benchmarkSerialization() {
  const { 2: level } = process.argv
  if (!level) throw "level needed";

  const root = await protobuf.load('example.proto');
  const buffer = root.lookupType(level);

  const jsonDataSting = await fs.readFile(path.join('..', 'common', 'json', `${level}.json`))
  const jsonData = JSON.parse(jsonDataSting);
  const protoData = buffer.encode(jsonData).finish();

  console.time('JSON Serialize');
  for (let i = 0; i < 10000; i++) {
    JSON.stringify(jsonData);
  }
  console.timeEnd('JSON Serialize');

  console.time('Protobuf Serialize');
  for (let i = 0; i < 10000; i++) {
    buffer.encode(jsonData).finish();
  }
  console.timeEnd('Protobuf Serialize');

  console.time('JSON Parse');
  for (let i = 0; i < 10000; i++) {
    JSON.parse(jsonDataSting);
  }
  console.timeEnd('JSON Parse');

  console.time('Protobuf Parse');
  for (let i = 0; i < 10000; i++) {
    buffer.decode(protoData);
  }
  console.timeEnd('Protobuf Parse');
}

benchmarkSerialization();
