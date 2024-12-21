import React from 'react';
import { RouteObject } from 'react-router-dom';
import Path from '@domain/constants/Path';
import RequiresAuthentication from '@routes/middleware/RequiresAuthentication';
import RouteElementWrapper from '@routes/middleware/RouteElementWrapper';
import Activate from '@pages/Activate';
import Index from '@pages/Index';
import Login from '@pages/Login';
import Register from '@pages/Register';
import MarshalIndex from '@pages/schools/marshal/MarshalIndex';
import SchoolsDetail from '@pages/schools/SchoolsDetail';
import SchoolsIndex from '@pages/schools/SchoolsIndex';
import StudentsIndex from '@pages/schools/students/StudentsIndex';
import UsersIndex from '@pages/schools/users/UsersIndex';
import WalkerIndex from '@pages/schools/walker/WalkerIndex';
import Temp from '@pages/Temp';

const routes: RouteObject[] = [
  {
    path: Path.INDEX,
    element: <Index />
  },
  {
    path: Path.REGISTER,
    element: <Register />
  },
  {
    path: Path.ACTIVATE,
    element: <Activate />
  },
  {
    path: Path.LOGIN,
    element: <Login />
  },
  {
    path: Path.MARSHAL,
    element: (
      <RouteElementWrapper>
        <RequiresAuthentication>
          <MarshalIndex />
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
          },
          {
            path: Path.USERS,
            element: (
              <RouteElementWrapper>
                <RequiresAuthentication>
                  <UsersIndex />
                </RequiresAuthentication>
              </RouteElementWrapper>
            )
          },
          {
            path: Path.MARSHAL,
            element: (
              <RouteElementWrapper>
                <RequiresAuthentication>
                  <MarshalIndex />
                </RequiresAuthentication>
              </RouteElementWrapper>
            )
          },
          {
            path: Path.WALKER,
            element: (
              <RouteElementWrapper>
                <RequiresAuthentication>
                  <WalkerIndex />
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
          <Temp />
        </RequiresAuthentication>
      </RouteElementWrapper>
    )
  },
  {
    path: Path.WALKER,
    element: (
      <RouteElementWrapper>
        <RequiresAuthentication>
          <WalkerIndex />
        </RequiresAuthentication>
      </RouteElementWrapper>
    )
  }
];

export default routes;
