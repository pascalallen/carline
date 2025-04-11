import React, { ReactElement, useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useParams } from 'react-router-dom';
import env, { EnvKey } from '@utilities/env';
import { Student } from '@domain/types/Student';
import useStore from '@hooks/useStore';
import StudentService from '@services/StudentService';
import Navbar from '@components/Navbar';
import Toast from '@components/Toast';

type State = {
  loading: boolean;
  students: Array<Student>;
  errorMessage: string;
};

const initialState: State = {
  loading: false,
  students: [],
  errorMessage: ''
};

const MarshalIndex = (): ReactElement => {
  const authStore = useStore('authStore');
  const { schoolId } = useParams();

  const [loading, setLoading] = useState(initialState.loading);
  const [students, setStudents] = useState(initialState.students);
  const [errorMessage, setErrorMessage] = useState(initialState.errorMessage);

  useEffect(() => {
    setLoading(true);
    const studentService = new StudentService(authStore);
    studentService
      .getAll(schoolId ?? '', { dismissed: true })
      .then((students: Student[]) => setStudents(students))
      .catch(error => setErrorMessage(error))
      .finally(() => setLoading(false));
  }, [authStore, schoolId]);

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
      <div className="toast-container top-0 end-0 p-3">
        {errorMessage && <Toast className="text-bg-danger">{errorMessage}</Toast>}
      </div>
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
        {loading ? (
          <Spinner />
        ) : (
          <section>
            <div className="row">
              <div className="col">
                <hr />
                <h1>Dismissed Students</h1>
                <ul className="student-list list-group list-group-flush overflow-y-scroll">
                  {students.length > 0 ? (
                    students.map(student => (
                      <li key={student.id} className="list-group-item">
                        {student.tag_number} {student.first_name} {student.last_name}
                      </li>
                    ))
                  ) : (
                    <p className="text-center">No students dismissed yet</p>
                  )}
                </ul>
              </div>
            </div>
          </section>
        )}
      </main>
    </div>
  );
};

export default MarshalIndex;
