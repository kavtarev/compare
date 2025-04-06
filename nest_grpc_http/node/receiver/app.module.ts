import { Module, NotAcceptableException, ValidationPipe } from "@nestjs/common";
import { APP_PIPE } from "@nestjs/core";
import { ReceiverController } from "./controller";

@Module({
  controllers: [ReceiverController],
  providers: [
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
