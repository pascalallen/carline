import React, { ReactElement } from 'react';
import { Helmet } from 'react-helmet-async';
import Footer from '@components/Footer';

const IndexPage = (): ReactElement => {
  return (
    <div id="index-page" className="index-page container">
      <Helmet>
        <title>Carline - Home</title>
        <meta name="description" content="Welcome to the home page for Carline" />
      </Helmet>
      <header>
        <h1>Home</h1>
      </header>
      <section>
        <h2>
          Lorem ipsum dolor sit amet, consectetur adipisicing elit. Assumenda beatae blanditiis error exercitationem
          facere fuga fugit harum ipsa maiores nulla qui quod sit soluta, sunt totam? Accusamus laboriosam mollitia
          quas.
        </h2>
      </section>
      <Footer />
    </div>
  );
};

export default IndexPage;
