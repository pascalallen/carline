import React from 'react';
import { Helmet } from 'react-helmet-async';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';

const StudentsPage = (): React.ReactElement => {
  return (
    <div id="students-page" className="students-page d-flex flex-column vh-100">
      <Helmet>
        <title>Carline - Students</title>
      </Helmet>
      <header>
        <Navbar />
      </header>
      <main className="container flex-fill mt-3">
        <section>
          <div className="row">
            <div className="col">
              <h1>Students</h1>
              <p>Use this page to manage (create, update, and delete) your students.</p>
            </div>
          </div>
        </section>
        <section>
          <div className="row">
            <div className="col">
              <table className="table">
                <thead>
                  <tr className="align-middle">
                    <th scope="col">ID</th>
                    <th scope="col">Tag Number</th>
                    <th scope="col">First Name</th>
                    <th scope="col">Last Name</th>
                    <th scope="col">School</th>
                    <th scope="col">
                      <button type="button" className="btn btn-success">
                        <i className="fa-solid fa-plus" />
                      </button>
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr className="align-middle">
                    <td>1</td>
                    <td>C246</td>
                    <td>Bobby</td>
                    <td>Tables</td>
                    <td>Weiss Elementary</td>
                    <td>
                      <button type="button" className="btn btn-danger">
                        <i className="fa-solid fa-trash" />
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default StudentsPage;
