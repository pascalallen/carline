import React, { ReactElement, useState } from 'react';
import { Helmet } from 'react-helmet-async';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';

type State = {
  studentNumber: string;
  studentNumbers: Array<string>;
};

const initialState: State = {
  studentNumber: '',
  studentNumbers: []
};

const WalkerPage = (): ReactElement => {
  const [studentNumber, setStudentNumber] = useState(initialState.studentNumber);
  const [studentNumbers, setStudentNumbers] = useState(initialState.studentNumbers);

  const handleNumberClick = (value: string) => setStudentNumber(prevInput => prevInput + value);
  const submitStudentNumber = () => {
    setStudentNumbers(prevInput => [...prevInput, studentNumber]);
    setStudentNumber(initialState.studentNumber);
  };

  return (
    <div id="walker-page" className="walker-page d-flex flex-column vh-100">
      <Helmet>
        <title>CarLine - Walker</title>
      </Helmet>
      <header>
        <Navbar />
      </header>
      <main className="container flex-fill mt-3">
        <section>
          <div className="row">
            <div className="col">
              <h1>Walker View</h1>
              <p>
                Well look at you! Today, you&apos;re the <strong>walker</strong>. Your job is to walk through the
                parking lot and log the student numbers that you see on parents cars using the number pad below. This
                allows the faculty inside the school classrooms to see which students are ready to be picked up.
              </p>
            </div>
          </div>
        </section>
        <section>
          <div className="row row-cols-auto justify-content-center">
            <div className="col">
              <div className="row mb-3">
                <div className="col">
                  <label htmlFor="student-number" className="form-label">
                    Input Value
                  </label>
                  <input
                    id="student-number"
                    type="text"
                    value={studentNumber}
                    readOnly
                    className="form-control form-control-lg text-right"
                  />
                </div>
              </div>
              <div className="row mb-3">
                {[1, 2, 3].map(num => (
                  <div key={num} className="col">
                    <div
                      className="btn btn-lg btn-primary d-flex justify-content-center align-items-center"
                      onClick={() => handleNumberClick(num.toString())}>
                      {num}
                    </div>
                  </div>
                ))}
              </div>
              <div className="row mb-3">
                {[4, 5, 6].map(num => (
                  <div key={num} className="col">
                    <div
                      className="btn btn-lg btn-primary d-flex justify-content-center align-items-center"
                      onClick={() => handleNumberClick(num.toString())}>
                      {num}
                    </div>
                  </div>
                ))}
              </div>
              <div className="row mb-3">
                {[7, 8, 9].map(num => (
                  <div key={num} className="col">
                    <div
                      className="btn btn-lg btn-primary d-flex justify-content-center align-items-center"
                      onClick={() => handleNumberClick(num.toString())}>
                      {num}
                    </div>
                  </div>
                ))}
              </div>
              <div className="row mb-3">
                <div className="col">
                  <div
                    className="btn btn-lg btn-secondary d-flex justify-content-center align-items-center"
                    onClick={() => setStudentNumber('')}>
                    C
                  </div>
                </div>
                <div className="col">
                  <div
                    className="btn btn-lg btn-primary d-flex justify-content-center align-items-center"
                    onClick={() => handleNumberClick('0')}>
                    0
                  </div>
                </div>
                <div className="col">
                  <div
                    className="btn btn-lg btn-success d-flex justify-content-center align-items-center"
                    onClick={submitStudentNumber}>
                    +
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
        <section>
          <div className="row">
            <div className="col">
              <hr />
              <h1>History</h1>
              <ul className="student-list list-group list-group-flush overflow-y-scroll">
                {studentNumbers.length > 0 ? (
                  studentNumbers.map(num => (
                    <li key={num} className="list-group-item">
                      {num}
                    </li>
                  ))
                ) : (
                  <p className="text-center">No student numbers added yet</p>
                )}
              </ul>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default WalkerPage;
