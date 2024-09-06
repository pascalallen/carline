import React, { ReactElement } from 'react';
import { Helmet } from 'react-helmet-async';
import Footer from '@components/Footer';

const IndexPage = (): ReactElement => {
  return (
    <div id="index-page" className="index-page d-flex flex-column vh-100">
      <Helmet>
        <title>Carline - Home</title>
        <meta name="description" content="Welcome to the home page for Carline" />
      </Helmet>
      <header>
        <h1>Home</h1>
      </header>
      <main className="container flex-fill mt-3">
        <section>
          <h2>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Assumenda beatae blanditiis error exercitationem
            facere fuga fugit harum ipsa maiores nulla qui quod sit soluta, sunt totam? Accusamus laboriosam mollitia
            quas.
          </h2>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default IndexPage;
