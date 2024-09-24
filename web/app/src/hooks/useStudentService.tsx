import StudentService from '@services/StudentService';
import useStore from './useStore';

const useStudentService = (): StudentService => {
  const authStore = useStore('authStore');

  return new StudentService(authStore);
};

export default useStudentService;
