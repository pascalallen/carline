import UserService from '@services/UserService';
import useStore from './useStore';

const useUserService = (): UserService => {
  const authStore = useStore('authStore');

  return new UserService(authStore);
};

export default useUserService;
