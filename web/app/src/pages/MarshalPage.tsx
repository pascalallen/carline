import React, { ReactElement } from 'react';
import { Helmet } from 'react-helmet-async';
import Navbar from '@components/Navbar';

const MarshalPage = (): ReactElement => {
  return (
    <div id="marshal-page" className="marshal-page d-flex flex-column vh-100">
      <Helmet>
        <title>CarLine - Marshal</title>
      </Helmet>
      <header>
        <Navbar />
      </header>
      <main className="container flex-fill mt-5">
        <h1>Ready For Pickup</h1>
        <section>
          <div className="row">
            <div className="col"></div>
          </div>
        </section>
      </main>
    </div>
  );
};

export default MarshalPage;
