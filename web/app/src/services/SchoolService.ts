import { queryStringify, removeEmptyKeys } from '@utilities/collections';
import request from '@utilities/request';
import { DomainEvents } from '@domain/constants/DomainEvents';
import HttpMethod from '@domain/constants/HttpMethod';
import { School } from '@domain/types/School';
import AuthStore from '@stores/AuthStore';
import eventDispatcher from '@services/eventDispatcher';

export type GetSchoolByIdResponse = {
  school: School;
};

export type GetAllSchoolsRequest = {
  include_deleted?: boolean;
};

export type GetAllSchoolsResponse = {
  schools: School[];
};

export type AddSchoolRequest = {
  name: string;
};

export type SchoolAddedResponse = {
  id: string;
};

export type SchoolRemovedResponse = {
  id: string;
};

class SchoolService {
  private readonly authStore: AuthStore;

  constructor(authStore: AuthStore) {
    this.authStore = authStore;
  }

  public async getById(id: string): Promise<School | undefined> {
    const response = await request.send<GetSchoolByIdResponse>({
      method: HttpMethod.GET,
      uri: `/api/v1/schools/${id}`,
      options: { auth: true, authStore: this.authStore }
    });

    return response.body.data?.school;
  }

  public async getAll(params: GetAllSchoolsRequest = { include_deleted: false }): Promise<School[]> {
    const queryParams = queryStringify(removeEmptyKeys(params || {}));
    const response = await request.send<GetAllSchoolsResponse>({
      method: HttpMethod.GET,
      uri: `/api/v1/schools${queryParams}`,
      options: { auth: true, authStore: this.authStore }
    });

    return response.body.data?.schools || [];
  }

  public async add(params: AddSchoolRequest) {
    const response = await request.send<SchoolAddedResponse>({
      method: HttpMethod.POST,
      uri: '/api/v1/schools',
      body: params,
      options: { auth: true, authStore: this.authStore }
    });

    eventDispatcher.dispatch({
      name: DomainEvents.SCHOOL_ADDED,
      data: {
        id: response.body.data?.id
      }
    });
  }

  public async remove(id: string) {
    const response = await request.send<SchoolRemovedResponse>({
      method: HttpMethod.DELETE,
      uri: `/api/v1/schools/${id}`,
      options: { auth: true, authStore: this.authStore }
    });

    eventDispatcher.dispatch({
      name: DomainEvents.SCHOOL_REMOVED,
      data: {
        id: response.body.data?.id
      }
    });
  }
}

export default SchoolService;
