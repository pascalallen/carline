import React, { useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useNavigate, useParams } from 'react-router-dom';
import { School } from '@domain/types/School';
import useSchoolService from '@hooks/useSchoolService';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';
import Toast from '@components/Toast';

type State = {
  loading: boolean;
  school?: School;
  errorMessage: string;
};

const initialState: State = {
  loading: true,
  errorMessage: ''
};

const SchoolsDetail = (): React.ReactElement => {
  const schoolService = useSchoolService();
  const navigate = useNavigate();
  const { schoolId } = useParams();

  const [loading, setLoading] = useState(initialState.loading);
  const [school, setSchool] = useState(initialState.school);
  const [errorMessage, setErrorMessage] = useState(initialState.errorMessage);

  useEffect(() => {
    setLoading(initialState.loading);
    schoolService
      .getById(schoolId ?? '')
      .then((school?: School) => setSchool(school))
      .catch(error => setErrorMessage(error))
      .finally(() => setLoading(false));
  }, [schoolId]);

  return (
    <div id="schools-detail-page" className="schools-detail-page d-flex flex-column vh-100">
      <div className="toast-container top-0 end-0 p-3">
        {errorMessage && <Toast className="text-bg-danger">{errorMessage}</Toast>}
      </div>
      <Helmet>
        <title>CarLine - {school?.name ?? ''}</title>
      </Helmet>
      <header>
        <Navbar />
      </header>
      <main className="container flex-fill mt-3">
        {loading ? (
          <Spinner />
        ) : (
          <>
            <section>
              <div className="row">
                <div className="col">
                  <h1>{school?.name}</h1>
                </div>
              </div>
            </section>
            <section>
              <div className="row">
                <div className="col-12 col-md-6 mb-3 mb-md-0">
                  <div className="card h-100">
                    <div className="card-body">
                      <h5 className="card-title">Manage Students</h5>
                      <p className="card-text">Create, update, or delete {school?.name} students.</p>
                      <a
                        className="btn btn-primary"
                        href={`/schools/${schoolId}/students`}
                        onClick={event => {
                          event.preventDefault();
                          navigate(`/schools/${schoolId}/students`);
                        }}>
                        Manage Students
                      </a>
                    </div>
                  </div>
                </div>
                <div className="col-12 col-md-6">
                  <div className="card h-100">
                    <div className="card-body">
                      <h5 className="card-title">Manage Users</h5>
                      <p className="card-text">Create, update, or delete your CarLine users for {school?.name}.</p>
                      <a
                        className="btn btn-primary"
                        href={`/schools/${schoolId}/users`}
                        onClick={event => {
                          event.preventDefault();
                          navigate(`/schools/${schoolId}/users`);
                        }}>
                        Manage Users
                      </a>
                    </div>
                  </div>
                </div>
              </div>
            </section>
          </>
        )}
      </main>
      <Footer />
    </div>
  );
};

export default SchoolsDetail;
