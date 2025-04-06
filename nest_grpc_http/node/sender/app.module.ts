import { Module } from "@nestjs/common";
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
  ],
})
export class AppModule {}
