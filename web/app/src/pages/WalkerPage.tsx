import React, { ReactElement, useState } from 'react';

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
    setStudentNumbers(prevNumbers => [...prevNumbers, studentNumber]);
    setStudentNumber(initialState.studentNumber);
  };

  return (
    <div id="walker-page" className="walker-page container">
      <div className="row row-cols-auto justify-content-center align-items-center mt-3 mb-3">
        <div className="col">
          <h1>Add Student</h1>
          <input
            type="text"
            value={studentNumber}
            readOnly
            placeholder="Student number"
            className="form-control form-control-lg text-right"
          />
        </div>
      </div>
      <div className="row row-cols-auto justify-content-center align-items-center mb-3">
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
      <div className="row row-cols-auto justify-content-center align-items-center mb-3">
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
      <div className="row row-cols-auto justify-content-center align-items-center mb-3">
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
      <div className="row row-cols-auto justify-content-center align-items-center mb-3">
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
      <div className="row row-cols-auto justify-content-center align-items-center">
        <div className="col">
          <hr />
          <h1>Student Numbers</h1>
          <ul className="list-group list-group-flush">
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
    </div>
  );
};

export default WalkerPage;
