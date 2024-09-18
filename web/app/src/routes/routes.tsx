import React from 'react';
import { RouteObject } from 'react-router-dom';
import Path from '@domain/constants/Path';
import RequiresAuthentication from '@routes/middleware/RequiresAuthentication';
import RouteElementWrapper from '@routes/middleware/RouteElementWrapper';
import IndexPage from '@pages/IndexPage';
import LoginPage from '@pages/LoginPage';
import MarshalPage from '@pages/MarshalPage';
import RegisterPage from '@pages/RegisterPage';
import SchoolsPage from '@pages/SchoolsPage';
import StudentsPage from '@pages/StudentsPage';
import TempPage from '@pages/TempPage';
import WalkerPage from '@pages/WalkerPage';

const routes: RouteObject[] = [
  {
    path: Path.INDEX,
    element: <IndexPage />
  },
  {
    path: Path.REGISTER,
    element: <RegisterPage />
  },
  {
    path: Path.LOGIN,
    element: <LoginPage />
  },
  {
    path: Path.MARSHALL,
    element: (
      <RouteElementWrapper>
        <RequiresAuthentication>
          <MarshalPage />
        </RequiresAuthentication>
      </RouteElementWrapper>
    )
  },
  {
    path: Path.SCHOOLS,
    element: (
      <RouteElementWrapper>
        <RequiresAuthentication>
          <SchoolsPage />
        </RequiresAuthentication>
      </RouteElementWrapper>
    )
  },
  {
    path: Path.STUDENTS,
    element: (
      <RouteElementWrapper>
        <RequiresAuthentication>
          <StudentsPage />
        </RequiresAuthentication>
      </RouteElementWrapper>
    )
  },
  {
    path: Path.TEMP,
    element: (
      <RouteElementWrapper>
        <RequiresAuthentication>
          <TempPage />
        </RequiresAuthentication>
      </RouteElementWrapper>
    )
  },
  {
    path: Path.WALKER,
    element: (
      <RouteElementWrapper>
        <RequiresAuthentication>
          <WalkerPage />
        </RequiresAuthentication>
      </RouteElementWrapper>
    )
  }
];

export default routes;
