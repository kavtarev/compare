import { Body, Controller, Inject, OnModuleInit, Post } from "@nestjs/common";
import { ClientGrpc } from "@nestjs/microservices";

interface ExampleService {
  getData(data: { id: string }): Promise<{ data: string }>;
}

@Controller()
export class SenderController implements OnModuleInit {
  private exampleService: ExampleService;

  constructor(@Inject("GRPC_CLIENT") private client: ClientGrpc) {}

  onModuleInit() {
    this.exampleService =
      this.client.getService<ExampleService>("ExampleService");
  }

  @Post("json")
  async executeJson(@Body() dto: any) {
    const res = await fetch("http://localhost:3001/json", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify(dto),
    });

    return res.json();
  }

  @Post("grpc")
  async executeGrpc(@Body() dto: any) {
    return this.exampleService.getData({ id: dto.id });
  }
}
