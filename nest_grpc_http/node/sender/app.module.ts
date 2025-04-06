import { Module } from "@nestjs/common";
import { SenderController } from "./controller";

@Module({
  controllers: [SenderController],
})
export class AppModule {}
