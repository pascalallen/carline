import { Student } from '@domain/types/Student';

export type School = {
  id: string;
  name: string;
  students: Student[];
  created_at: string;
  modified_at?: string;
  deleted_at?: string;
};
