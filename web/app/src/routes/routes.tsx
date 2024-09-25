import React from 'react';
import { RouteObject } from 'react-router-dom';
import Path from '@domain/constants/Path';
import RequiresAuthentication from '@routes/middleware/RequiresAuthentication';
import RouteElementWrapper from '@routes/middleware/RouteElementWrapper';
import IndexPage from '@pages/IndexPage';
import LoginPage from '@pages/LoginPage';
import MarshalPage from '@pages/MarshalPage';
import RegisterPage from '@pages/RegisterPage';
import SchoolsDetail from '@pages/schools/SchoolsDetail';
import SchoolsIndex from '@pages/schools/SchoolsIndex';
import StudentsIndex from '@pages/schools/students/StudentsIndex';
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
    children: [
      {
        index: true,
        element: (
          <RouteElementWrapper>
            <RequiresAuthentication>
              <SchoolsIndex />
            </RequiresAuthentication>
          </RouteElementWrapper>
        )
      },
      {
        path: Path.SCHOOL,
        children: [
          {
            index: true,
            element: (
              <RouteElementWrapper>
                <RequiresAuthentication>
                  <SchoolsDetail />
                </RequiresAuthentication>
              </RouteElementWrapper>
            )
          },
          {
            path: Path.STUDENTS,
            element: (
              <RouteElementWrapper>
                <RequiresAuthentication>
                  <StudentsIndex />
                </RequiresAuthentication>
              </RouteElementWrapper>
            )
          }
        ]
      }
    ]
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
