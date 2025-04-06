import { NestFactory } from "@nestjs/core";

import { AppModule } from "./app.module";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.enableShutdownHooks();

  await app.listen(3000);

  console.log(`Up on port: ${3000}`);
}

if (require.main === module) {
  bootstrap();
}
