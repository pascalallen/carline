import React, { ReactElement } from 'react';
import { Helmet } from 'react-helmet-async';
import { useNavigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import Footer from '@components/Footer';

const Index = (): ReactElement => {
  const navigate = useNavigate();

  return (
    <div id="index-page" className="index-page d-flex flex-column vh-100">
      <Helmet>
        <title>CarLine - Efficient School Dismissal</title>
        <meta name="description" content="Helping parents and teachers streamline end-of-day classroom pickups." />
      </Helmet>

      {/* Header Section */}
      <header className="container mt-3 text-center">
        <h1 className="display-4 text-primary fw-bold">Streamline School Pickups with CarLine</h1>
        <p className="lead text-secondary mt-2">
          Helping parents and teachers manage school dismissals quickly and efficiently.
        </p>
        <a href="#features" className="btn btn-primary btn-lg mt-4">
          Learn How
        </a>
      </header>

      {/* Main Section */}
      <main className="container flex-fill mt-5" id="features">
        <section className="text-center mb-5">
          <h2 className="fw-bold mb-4">How CarLine Helps</h2>
          <div className="row">
            <div className="col-md-6 mb-4">
              <div className="card shadow-sm">
                <div className="card-body">
                  <h5 className="card-title fw-bold">Efficient Parent Pickup</h5>
                  <h6 className="card-subtitle mb-2 text-body-secondary">A smoother experience for parents</h6>
                  <p className="card-text">
                    Quickly notify classrooms that you&#39;ve arrived so that kids are ready for pickup, reducing
                    waiting time.
                  </p>
                  <a href="#" className="card-link text-primary">
                    Learn more
                  </a>
                </div>
              </div>
            </div>
            <div className="col-md-6 mb-4">
              <div className="card shadow-sm">
                <div className="card-body">
                  <h5 className="card-title fw-bold">Organized Dismissals</h5>
                  <h6 className="card-subtitle mb-2 text-body-secondary">Simplified for teachers and staff</h6>
                  <p className="card-text">
                    Keep track of which parents have arrived and ensure kids are dismissed efficiently and safely.
                  </p>
                  <a href="#" className="card-link text-primary">
                    Learn more
                  </a>
                </div>
              </div>
            </div>
          </div>
        </section>

        {/* CTA Section */}
        <section className="text-center mt-5">
          <h2 className="fw-bold mb-4">Ready to improve your school&#39;s dismissal process?</h2>
          <p className="lead text-secondary mb-4">
            Join hundreds of schools in using CarLine to simplify school pickups for teachers, parents, and kids.
          </p>
          <a
            href={Path.REGISTER}
            onClick={event => {
              event.preventDefault();
              navigate(Path.REGISTER);
            }}
            className="btn btn-success btn-lg">
            Get Started Today
          </a>
        </section>
      </main>

      {/* Footer Section */}
      <Footer />
    </div>
  );
};

export default Index;
