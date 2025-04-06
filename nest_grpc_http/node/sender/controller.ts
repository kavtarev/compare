import { Controller, Get } from "@nestjs/common";

@Controller()
export class SenderController {
  @Get("json")
  async executeJson() {
    const res = await fetch("http://localhost:3001/json", {
      method: "POST",
      headers: { "content-type": "application/json" },
      body: JSON.stringify({ hello: "world" }),
    });

    const data = await res.json();
    return data;
  }

  @Get("grpc")
  async executeGrpc() {
    return "grpc";
  }
}
