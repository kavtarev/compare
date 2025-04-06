import { Module } from "@nestjs/common";
import { ReceiverController } from "./controller";

@Module({
  controllers: [ReceiverController],
})
export class AppModule {}
