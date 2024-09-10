import React, { ReactElement, ReactNode } from 'react';
import { useLocation } from 'react-router';
import { Navigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';

export type RequiresAuthenticationProps = {
  children: ReactNode;
};

const RequiresAuthentication = (props: RequiresAuthenticationProps): ReactElement => {
  const { children } = props;

  const authService = useAuth();
  const location = useLocation();

  if (!authService.isLoggedIn()) {
    return <Navigate to={Path.LOGIN} state={{ from: location }} replace />;
  }

  return <>{children}</>;
};

export default RequiresAuthentication;
