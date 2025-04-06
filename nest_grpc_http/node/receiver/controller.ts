import { Body, Controller, Get, Post } from "@nestjs/common";
import { GrpcMethod } from "@nestjs/microservices";
import { IListTrackingExclusionResponse } from "./interface";

const response: IListTrackingExclusionResponse = {
  page: 0,
  hasNextPage: false,
  data: [
    {
      exclusionId: "exclusionId",
      name: "name",
      vin: "vin",
      leasingSubjectId: "leasingSubjectId",
      authorId: "authorId",
      authorFio: "authorFio",
      fromDate: "fromDate",
      toDate: "toDate",
      isActive: true,
    },
    {
      exclusionId: "exclusionId",
      name: "name",
      vin: "vin",
      leasingSubjectId: "leasingSubjectId",
      authorId: "authorId",
      authorFio: "authorFio",
      fromDate: "fromDate",
      toDate: "toDate",
      isActive: true,
    },
  ],
};
@Controller()
export class ReceiverController {
  @Post("json-tiny")
  async executeJsonTiny(@Body() dto: any) {
    return { data: `json Data for ID: ${dto.id}` };
  }

  @Post("json-medium")
  async executeJsonMedium(@Body() dto: any) {
    return response;
  }

  @GrpcMethod("ExampleService", "DataTiny")
  DataTiny(data: { id: string }): { data: string } {
    return { data: `json Data for ID: ${data.id}` };
  }

  @GrpcMethod("ExampleService", "DataMedium")
  DataMedium(data: { id: string }): IListTrackingExclusionResponse {
    return response;
  }
}
