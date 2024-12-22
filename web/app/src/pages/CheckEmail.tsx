import React, { ReactElement } from 'react';
import { Helmet } from 'react-helmet-async';
import { useLocation } from 'react-router';
import Path from '@domain/constants/Path';
import Footer from '@components/Footer';

const CheckEmail = (): ReactElement => {
  const location = useLocation();
  const emailAddress = location.state?.emailAddress || 'your email address';

  return (
    <div id="check-email-page" className="check-email-page d-flex flex-column vh-100">
      <Helmet>
        <title>CarLine - Check Email</title>
      </Helmet>
      <main className="container flex-fill mt-3">
        <section>
          <div className="row row-cols-auto justify-content-center">
            <div className="col">
              <h1 className="text-success">Thank You for Registering!</h1>
              <h2>Just one more step to complete your registration.</h2>
              <p>
                We&#39;ve sent an activation email to <strong>{emailAddress}</strong>. Please check your inbox and click
                on the activation link to verify your email address and activate your account.
              </p>
              <p className="text-muted">
                If you don&#39;t see the email within a few minutes, check your spam or junk folders. Once your email is
                verified, you&#39;ll be able to log in and access all the amazing features of CarLine.
              </p>
              <div className="d-flex justify-content-center mt-4">
                <a href={Path.LOGIN} className="btn btn-primary me-2">
                  Back to Login
                </a>
              </div>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default CheckEmail;
