import Fastify from 'fastify'
import { readFile } from 'fs/promises';
import { join } from 'path';

async function run() {
  const path = join('..', 'common', 'json', `${process.argv[2]}.json`);
  const data = await readFile(path);
  const json = JSON.parse(data);

  const fastify = Fastify({})

  // Declare a route
  fastify.get('/json-stringify', function (request, reply) {
    reply.send(json)
  })

  // Run the server!
  fastify.listen({ port: 3000 }, function (err, address) {
    if (err) {
      fastify.log.error(err)
      process.exit(1)
    }
    // Server is now listening on ${address}
  })
}

run()