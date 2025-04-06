import { Body, Controller, Get, Post } from "@nestjs/common";
import { GrpcMethod } from "@nestjs/microservices";

@Controller()
export class ReceiverController {
  @Post("json")
  async executeJson(@Body() dto: any) {
    return { data: `json Data for ID: ${dto.id}` };
  }

  @GrpcMethod("ExampleService", "GetData")
  getData(data: { id: string }): { data: string } {
    return { data: `grpc Data for ID: ${data.id}` };
  }
}
