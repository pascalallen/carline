import SchoolService from '@services/SchoolService';
import useStore from './useStore';

const useSchoolService = (): SchoolService => {
  const authStore = useStore('authStore');

  return new SchoolService(authStore);
};

export default useSchoolService;
