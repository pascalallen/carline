import React, { ReactElement } from 'react';
import { Helmet } from 'react-helmet-async';
import Footer from '@components/Footer';

const IndexPage = (): ReactElement => {
  return (
    <div id="index-page" className="index-page d-flex flex-column vh-100">
      <Helmet>
        <title>CarLine - Home</title>
        <meta name="description" content="Welcome to the home page for CarLine" />
      </Helmet>
      <header className="container mt-3">
        <h1>CarLine</h1>
      </header>
      <main className="container flex-fill mt-3">
        <section>
          <h2>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Assumenda beatae blanditiis error exercitationem
            facere fuga fugit harum ipsa maiores nulla qui quod sit soluta, sunt totam? Accusamus laboriosam mollitia
            quas.
          </h2>
          <div className="row mt-3">
            <div className="col">
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">Card title</h5>
                  <h6 className="card-subtitle mb-2 text-body-secondary">Card subtitle</h6>
                  <p className="card-text">
                    Some quick example text to build on the card title and make up the bulk of the card's content.
                  </p>
                  <a href="#" className="card-link">
                    Card link
                  </a>
                  <a href="#" className="card-link">
                    Another link
                  </a>
                </div>
              </div>
            </div>
            <div className="col">
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">Card title</h5>
                  <h6 className="card-subtitle mb-2 text-body-secondary">Card subtitle</h6>
                  <p className="card-text">
                    Some quick example text to build on the card title and make up the bulk of the card's content.
                  </p>
                  <a href="#" className="card-link">
                    Card link
                  </a>
                  <a href="#" className="card-link">
                    Another link
                  </a>
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default IndexPage;
