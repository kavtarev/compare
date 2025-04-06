import { Module, NotAcceptableException, ValidationPipe } from "@nestjs/common";
import { APP_PIPE } from "@nestjs/core";
import { ClientProxyFactory, Transport } from "@nestjs/microservices";
import { join } from "path";
import { SenderController } from "./controller";

@Module({
  controllers: [SenderController],
  providers: [
    {
      provide: "GRPC_CLIENT",
      useFactory: () => {
        return ClientProxyFactory.create({
          transport: Transport.GRPC,
          options: {
            package: "example",
            protoPath: join(__dirname, "..", "..", "common.proto"),
            url: "localhost:3003", // Укажите адрес и порт gRPC-сервера
          },
        });
      },
    },
    {
      provide: APP_PIPE,
      useFactory: () => {
        return new ValidationPipe({
          whitelist: true,
          forbidNonWhitelisted: true,
          exceptionFactory: (errors: unknown) => {
            return new NotAcceptableException(errors);
          },
        });
      },
    },
  ],
})
export class AppModule {}
