export interface ExampleService {
  DataTiny(data: { id: string }): Promise<{ data: string }>;

  DataMedium(
    data: IListTrackingExclusionDto
  ): Promise<IListTrackingExclusionResponse>;
}

export interface IListTrackingExclusionDto {
  page: number;
  limit: number;
  orderBy: ListTrackingExclusionOrder | undefined;
  q: string | undefined;
  asc: boolean | undefined;
}

export enum ListTrackingExclusionOrder {
  vin = "vin",
  toDate = "toDate",
  fromDate = "fromDate",
  name = "name",
  authorFio = "authorFio",
}

export interface IListTrackingExclusionResponse {
  page: number;
  hasNextPage: boolean;
  data: {
    exclusionId: string;
    name: string;
    vin: string;
    leasingSubjectId: string;
    authorId: string;
    authorFio: string;
    fromDate: string;
    toDate: string;
    isActive: boolean;
  }[];
}
