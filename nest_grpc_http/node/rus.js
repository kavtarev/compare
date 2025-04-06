const http = require('http');

const postData = JSON.stringify({ id: '90' });

const { 2: type, 3: count } = process.argv

const options = {
  hostname: '127.0.0.1',
  port: 3000,
  path: `/${type}`,
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
    'Content-Length': Buffer.byteLength(postData),
  },
};

async function sendRequest() {
  return new Promise((resolve, reject) => {
    const req = http.request(options, (res) => {
      let data = '';

      res.on('data', (chunk) => {
        data += chunk;
      });

      res.on('end', () => {
        resolve(data);
      });
    });

    req.write(postData);
    req.end();
  });
}

async function main() {
  const startTime = Date.now();

  for (let i = 0; i < count; i++) {
    try {
      await sendRequest();
    } catch (error) {
      console.log(error);

    }
  }

  console.log(`Общее время выполнения: ${Date.now() - startTime} миллисекунд`);
}

main();
