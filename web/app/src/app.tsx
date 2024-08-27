import React, { ReactElement } from 'react';
import { createRoot } from 'react-dom/client';
import '@assets/scss/app.scss';

const container: HTMLElement | null = document.getElementById('root');
if (container === null) {
  throw new Error('No matching element found with ID: root');
}

const App = (): ReactElement => {
  return <React.StrictMode>Hello, World!</React.StrictMode>;
};

const root = createRoot(container);
root.render(<App />);
