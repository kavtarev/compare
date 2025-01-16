import { createServer } from 'http';

const PORT = 3000;

async function run() {
  const server = createServer((req, res) => { })
  server.listen(PORT, () => { console.log(`up on ${PORT}`); })
}

run()

/**
 * Represents a book.
 * @param {(string | number)[]} title - The title of the book.
 * @param {string} author - The author of the book.
 */
function Book(title, author) {

}