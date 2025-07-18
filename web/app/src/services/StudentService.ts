import { queryStringify, removeEmptyKeys } from '@utilities/collections';
import request from '@utilities/request';
import { DomainEvents } from '@domain/constants/DomainEvents';
import HttpMethod from '@domain/constants/HttpMethod';
import { Student } from '@domain/types/Student';
import AuthStore from '@stores/AuthStore';
import eventDispatcher from '@services/eventDispatcher';

export type GetAllStudentsResponse = {
  students: Student[];
};

export type StudentRemovedResponse = {
  id: string;
};

export type DismissStudentRequest = {
  tag_number: string;
};

export type StudentDismissalResponse = {
  tag_number: string;
};

class StudentService {
  private readonly authStore: AuthStore;

  constructor(authStore: AuthStore) {
    this.authStore = authStore;
  }

  public async dismiss(schoolId: string, params: DismissStudentRequest) {
    const response = await request.send<StudentDismissalResponse>({
      method: HttpMethod.POST,
      uri: `/api/v1/schools/${schoolId}/students/dismissals`,
      body: params,
      options: { auth: true, authStore: this.authStore }
    });

    eventDispatcher.dispatch({
      name: DomainEvents.STUDENT_DISMISSED,
      data: {
        tag_number: response.body.data?.tag_number
      }
    });
  }

  public async getAll(schoolId: string, params?: { dismissed?: boolean }): Promise<Student[]> {
    if (!params) {
      params = {};
    }
    const response = await request.send<GetAllStudentsResponse>({
      method: HttpMethod.GET,
      uri: `/api/v1/schools/${schoolId}/students${queryStringify(removeEmptyKeys(params))}`,
      options: { auth: true, authStore: this.authStore }
    });

    return response.body.data?.students || [];
  }

  public async import(schoolId: string, formData: FormData): Promise<void> {
    await request.send({
      method: HttpMethod.POST,
      uri: `/api/v1/schools/${schoolId}/students/import`,
      body: formData,
      options: {
        auth: true,
        authStore: this.authStore,
        contentType: 'multipart/form-data'
      }
    });

    eventDispatcher.dispatch({
      name: DomainEvents.STUDENTS_IMPORTED,
      data: {}
    });
  }

  public async remove(schoolId: string, id: string) {
    const response = await request.send<StudentRemovedResponse>({
      method: HttpMethod.DELETE,
      uri: `/api/v1/schools/${schoolId}/students/${id}`,
      options: { auth: true, authStore: this.authStore }
    });

    eventDispatcher.dispatch({
      name: DomainEvents.STUDENT_REMOVED,
      data: {
        id: response.body.data?.id
      }
    });
  }
}

export default StudentService;
