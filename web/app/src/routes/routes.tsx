import React from 'react';
import { RouteObject } from 'react-router-dom';
import Path from '@domain/constants/Path';
import IndexPage from '@pages/IndexPage';
import LoginPage from '@pages/LoginPage';
import TempPage from '@pages/TempPage';
import WalkerPage from '@pages/WalkerPage';
import RequiresAuthentication from './middleware/RequiresAuthentication';
import RouteElementWrapper from './middleware/RouteElementWrapper';

const routes: RouteObject[] = [
  {
    path: Path.INDEX,
    element: <IndexPage />
  },
  {
    path: Path.LOGIN,
    element: <LoginPage />
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
      // <RouteElementWrapper>
      //   <RequiresAuthentication>
      <WalkerPage />
      // </RequiresAuthentication>
      // </RouteElementWrapper>
    )
  }
];

export default routes;
