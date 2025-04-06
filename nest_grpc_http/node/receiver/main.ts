import { NestFactory } from "@nestjs/core";
import { MicroserviceOptions, Transport } from "@nestjs/microservices";
import { join } from "path";

import { AppModule } from "./app.module";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.enableShutdownHooks();

  const microservice = app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.GRPC,
    options: {
      package: "example",
      protoPath: join(__dirname, "..", "..", "common.proto"),
      url: "localhost:3003", // Укажите адрес и порт для gRPC-сервера
    },
  });
  await app.startAllMicroservices();
  await app.listen(3001);

  console.log(`Up on port: ${3001}`);
}

if (require.main === module) {
  bootstrap();
}
