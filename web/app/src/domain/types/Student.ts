import { School } from '@domain/types/School';

export type Student = {
  id: string;
  tag_number: string;
  first_name: string;
  last_name: string;
  school: School;
  created_at: string;
  modified_at?: string;
  deleted_at?: string;
};
