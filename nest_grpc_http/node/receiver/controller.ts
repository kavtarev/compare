import { Body, Controller, Get, Post } from "@nestjs/common";

@Controller()
export class ReceiverController {
  @Post("json")
  async executeJson(@Body() dto: any) {
    console.log(dto);

    return { some: "body" };
  }

  @Get("grpc")
  async executeGrpc() {
    return "grpc";
  }
}
