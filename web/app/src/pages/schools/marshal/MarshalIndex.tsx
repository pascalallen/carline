import React, { ReactElement } from 'react';
import { Helmet } from 'react-helmet-async';
import Navbar from '@components/Navbar';

const MarshalIndex = (): ReactElement => {
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
            <div className="col">
              <h1>Marshal View</h1>
              <p>
                Hey there! Today, you&apos;re the <strong>marshal</strong>. Your job is to monitor the incoming tag
                numbers from the parking lot walker and then dismiss students accordingly.
              </p>
            </div>
          </div>
        </section>
      </main>
    </div>
  );
};

export default MarshalIndex;
