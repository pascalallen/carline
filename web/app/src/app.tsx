import React, { ReactElement } from 'react';
import 'bootstrap/dist/js/bootstrap.min.js';
import { createRoot } from 'react-dom/client';
import { Helmet, HelmetProvider } from 'react-helmet-async';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import routes from '@routes/routes';
import { storesInstance, StoresProvider } from '@stores/Stores';
import '@assets/scss/app.scss';

const container: HTMLElement | null = document.getElementById('root');
if (container === null) {
  throw new Error('No matching element found with ID: root');
}

const App = (): ReactElement => {
  return (
    <React.StrictMode>
      <StoresProvider value={storesInstance}>
        <HelmetProvider>
          <Helmet>
            <title>CarLine</title>
            <meta name="description" content="CarLine app" />
            <meta name="robots" content="index, follow" />
            <link rel="canonical" href="https://carline.com/" />
          </Helmet>
          <RouterProvider router={createBrowserRouter(routes)} />
        </HelmetProvider>
      </StoresProvider>
    </React.StrictMode>
  );
};

const root = createRoot(container);
root.render(<App />);
