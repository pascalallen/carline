import request from '@utilities/request';
import { DomainEvents } from '@domain/constants/DomainEvents';
import HttpMethod from '@domain/constants/HttpMethod';
import { User } from '@domain/types/User';
import AuthStore from '@stores/AuthStore';
import eventDispatcher from '@services/eventDispatcher';

export type GetAllUsersResponse = {
  users: User[];
};

export type UserRemovedResponse = {
  id: string;
};

class UserService {
  private readonly authStore: AuthStore;

  constructor(authStore: AuthStore) {
    this.authStore = authStore;
  }

  public async getAll(schoolId: string): Promise<User[]> {
    const response = await request.send<GetAllUsersResponse>({
      method: HttpMethod.GET,
      uri: `/api/v1/schools/${schoolId}/users`,
      options: { auth: true, authStore: this.authStore }
    });

    return response.body.data?.users || [];
  }

  // TODO
  public async add(schoolId: string, formData: FormData): Promise<void> {
    await request.send({
      method: HttpMethod.POST,
      uri: `/api/v1/schools/${schoolId}/users`,
      body: formData,
      options: {
        auth: true,
        authStore: this.authStore,
        contentType: 'multipart/form-data'
      }
    });

    eventDispatcher.dispatch({
      name: DomainEvents.USER_ADDED,
      data: {}
    });
  }

  public async remove(schoolId: string, id: string) {
    const response = await request.send<UserRemovedResponse>({
      method: HttpMethod.DELETE,
      uri: `/api/v1/schools/${schoolId}/users/${id}`,
      options: { auth: true, authStore: this.authStore }
    });

    eventDispatcher.dispatch({
      name: DomainEvents.USER_REMOVED,
      data: {
        id: response.body.data?.id
      }
    });
  }
}

export default UserService;
