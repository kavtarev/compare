import { Body, Controller, Inject, OnModuleInit, Post } from "@nestjs/common";
import { ClientGrpc } from "@nestjs/microservices";
import { ExampleService } from "./interface";

@Controller()
export class SenderController implements OnModuleInit {
  private exampleService: ExampleService;

  constructor(@Inject("GRPC_CLIENT") private client: ClientGrpc) {}

  onModuleInit() {
    this.exampleService =
      this.client.getService<ExampleService>("ExampleService");
  }

  @Post("json-tiny")
  async executeJsonTiny(@Body() dto: any) {
    const res = await fetch("http://localhost:3001/json-tiny", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify(dto),
    });

    return res.json();
  }

  @Post("json-medium")
  async executeJsonMedium(@Body() dto: any) {
    const res = await fetch("http://localhost:3001/json-medium", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify(dto),
    });

    return res.json();
  }

  @Post("grpc-tiny")
  async executeGrpcTiny(@Body() dto: any) {
    return this.exampleService.DataTiny(dto);
  }

  @Post("grpc-medium")
  async executeGrpcMedium(@Body() dto: any) {
    return this.exampleService.DataMedium(dto);
  }
}
