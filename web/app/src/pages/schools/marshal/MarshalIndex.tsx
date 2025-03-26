import React, { ReactElement, useEffect } from 'react';
import { Helmet } from 'react-helmet-async';
import { useParams } from 'react-router-dom';
import env, { EnvKey } from '@utilities/env';
import Navbar from '@components/Navbar';

const MarshalIndex = (): ReactElement => {
  const { schoolId } = useParams();

  useEffect(() => {
    const socket = new WebSocket(`${env(EnvKey.APP_BASE_URL)}/api/v1/schools/${schoolId}/students/dismissals/ws`);

    socket.onopen = () => console.log('WebSocket connected');
    socket.onmessage = msg => {
      try {
        const data = JSON.parse(msg.data);
        console.log('Received:', data);
      } catch (err) {
        console.error('Error parsing message:', err);
      }
    };
    socket.onclose = () => console.log('WebSocket disconnected');
    socket.onerror = error => console.error('WebSocket error:', error);

    return () => {
      socket.close();
      console.log('WebSocket cleaned up');
    };
  }, [schoolId]);

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
