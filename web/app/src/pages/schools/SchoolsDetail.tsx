import React, { useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useParams } from 'react-router-dom';
import { School } from '@domain/types/School';
import useSchoolService from '@hooks/useSchoolService';
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
                  Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ab alias consequatur delectus, dolore eos
                  exercitationem iusto laboriosam minus, molestiae obcaecati pariatur possimus rerum sed suscipit
                  tempora ullam unde veritatis voluptatem.
                </div>
                <div className="col">
                  Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ab alias consequatur delectus, dolore eos
                  exercitationem iusto laboriosam minus, molestiae obcaecati pariatur possimus rerum sed suscipit
                  tempora ullam unde veritatis voluptatem.
                </div>
              </div>
            </section>
          </>
        )}
      </main>
    </div>
  );
};

export default SchoolsDetail;
