import React, { FormEvent, useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useNavigate } from 'react-router-dom';
import { DomainEvents } from '@domain/constants/DomainEvents';
import { School } from '@domain/types/School';
import useEvent from '@hooks/useEvents';
import useStore from '@hooks/useStore';
import { ErrorApiResponse, FailApiResponse } from '@services/ApiService';
import { DomainEvent } from '@services/eventDispatcher';
import SchoolService from '@services/SchoolService';
import AddSchoolModal from '@components/AddSchoolModal';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';
import RemoveSchoolModal from '@components/RemoveSchoolModal';
import Toast from '@components/Toast';

type State = {
  loading: boolean;
  schools: School[];
  showAddSchoolModal: boolean;
  addingSchool: boolean;
  selectedSchool?: School;
  showRemoveSchoolModal: boolean;
  removingSchool: boolean;
  errorMessage: string;
};

const initialState: State = {
  loading: false,
  schools: [],
  showAddSchoolModal: false,
  addingSchool: false,
  showRemoveSchoolModal: false,
  removingSchool: false,
  errorMessage: ''
};

const SchoolsIndex = (): React.ReactElement => {
  const authStore = useStore('authStore');
  const navigate = useNavigate();

  const [loading, setLoading] = useState(initialState.loading);
  const [schools, setSchools] = useState(initialState.schools);
  const [showAddSchoolModal, setShowAddSchoolModal] = useState(initialState.showAddSchoolModal);
  const [addingSchool, setAddingSchool] = useState(initialState.addingSchool);
  const [selectedSchool, setSelectedSchool] = useState(initialState.selectedSchool);
  const [showRemoveSchoolModal, setShowRemoveSchoolModal] = useState(initialState.showRemoveSchoolModal);
  const [removingSchool, setRemovingSchool] = useState(initialState.removingSchool);
  const [errorMessage, setErrorMessage] = useState(initialState.errorMessage);

  const schoolAddedEvent: DomainEvent | undefined = useEvent(DomainEvents.SCHOOL_ADDED);
  const schoolRemovedEvent: DomainEvent | undefined = useEvent(DomainEvents.SCHOOL_REMOVED);

  useEffect(() => {
    setLoading(true);
    const schoolService = new SchoolService(authStore);
    schoolService
      .getAll()
      .then((schools: School[]) => setSchools(schools))
      .catch(error => {
        let errorMessage = 'An unexpected error occurred.';

        if ((error as FailApiResponse)?.statusCode === 400) {
          errorMessage = 'Validation error.';
        } else if ((error as ErrorApiResponse)?.statusCode === 422) {
          errorMessage = (error as ErrorApiResponse).body.message ?? errorMessage;
        }

        setErrorMessage(errorMessage);
      })
      .finally(() => setLoading(initialState.loading));
  }, [authStore]);

  useEffect(() => {
    if (
      (schoolAddedEvent !== undefined && schoolAddedEvent.id !== undefined) ||
      (schoolRemovedEvent !== undefined && schoolRemovedEvent.id !== undefined)
    ) {
      setLoading(true);
      const schoolService = new SchoolService(authStore);
      schoolService
        .getAll()
        .then((schools: School[]) => setSchools(schools))
        .catch(error => {
          let errorMessage = 'An unexpected error occurred.';

          if ((error as FailApiResponse)?.statusCode === 400) {
            errorMessage = 'Validation error.';
          } else if ((error as ErrorApiResponse)?.statusCode === 422) {
            errorMessage = (error as ErrorApiResponse).body.message ?? errorMessage;
          }

          setErrorMessage(errorMessage);
        })
        .finally(() => setLoading(initialState.loading));
    }
  }, [authStore, schoolAddedEvent, schoolRemovedEvent]);

  const handleShowAddSchoolModal = (): void => setShowAddSchoolModal(true);
  const handleHideAddSchoolModal = (): void => setShowAddSchoolModal(initialState.showAddSchoolModal);

  const handleAddSchool = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    setErrorMessage(initialState.errorMessage);

    try {
      setAddingSchool(true);
      const formData = new FormData(event.currentTarget);
      const name = formData.get('name')?.toString() ?? '';
      const schoolService = new SchoolService(authStore);
      await schoolService.add({ name });
    } catch (error) {
      // 400 fail, 422 error, 500 error
      if ((error as FailApiResponse)?.statusCode === 400) {
        setErrorMessage('Validation error');
      }

      if ((error as ErrorApiResponse)?.statusCode === 422) {
        setErrorMessage((error as ErrorApiResponse).body.message);
      }

      if ((error as ErrorApiResponse)?.statusCode === 500) {
        setErrorMessage((error as ErrorApiResponse).body.message);
      }
    } finally {
      setAddingSchool(initialState.addingSchool);
      handleHideAddSchoolModal();
    }
  };

  const handleShowRemoveSchoolModal = (school: School): void => {
    setShowRemoveSchoolModal(true);
    setSelectedSchool(school);
  };

  const handleHideRemoveSchoolModal = (): void => {
    setShowRemoveSchoolModal(initialState.showRemoveSchoolModal);
    setSelectedSchool(initialState.selectedSchool);
  };

  const handleRemoveSchool = async (): Promise<void> => {
    try {
      setRemovingSchool(true);
      const schoolService = new SchoolService(authStore);
      await schoolService.remove(selectedSchool?.id ?? '');
      setSchools(prevSchools => prevSchools.filter(school => school.id !== selectedSchool?.id));
    } catch (error) {
      console.error('Error removing school: ', error);
    } finally {
      setRemovingSchool(initialState.removingSchool);
      handleHideRemoveSchoolModal();
    }
  };

  return (
    <div id="schools-page" className="schools-page d-flex flex-column vh-100">
      <div className="toast-container top-0 end-0 p-3">
        {errorMessage && <Toast className="text-bg-danger">{errorMessage}</Toast>}
      </div>
      <Helmet>
        <title>Carline - Schools</title>
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
                  <h1>My Schools</h1>
                  <p>Use this page to manage (create, update, and delete) your schools.</p>
                </div>
              </div>
            </section>
            <section>
              {showAddSchoolModal ? (
                <AddSchoolModal
                  show={showAddSchoolModal}
                  onClose={handleHideAddSchoolModal}
                  onAdd={handleAddSchool}
                  isAdding={addingSchool}
                />
              ) : null}
              {showRemoveSchoolModal ? (
                <RemoveSchoolModal
                  show={showRemoveSchoolModal}
                  onClose={handleHideRemoveSchoolModal}
                  school={selectedSchool}
                  onRemove={handleRemoveSchool}
                  isRemoving={removingSchool}
                />
              ) : null}
              <div className="row">
                <div className="col">
                  <table className="table">
                    <thead>
                      <tr className="align-middle">
                        <th scope="col">ID</th>
                        <th scope="col">Name</th>
                        <th scope="col">
                          <button type="button" className="btn btn-success" onClick={() => handleShowAddSchoolModal()}>
                            <i className="fa-solid fa-plus" />
                          </button>
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      {schools.map((school: School, index: number) => (
                        <tr key={`school-row-${index}`} className="align-middle">
                          <td>
                            <a
                              href={`/schools/${school.id}/students`}
                              onClick={event => {
                                event.preventDefault();
                                navigate(`/schools/${school.id}`);
                              }}>
                              {school.id}
                            </a>
                          </td>
                          <td>{school.name}</td>
                          <td>
                            <button
                              type="button"
                              className="btn btn-danger"
                              onClick={() => handleShowRemoveSchoolModal(school)}>
                              <i className="fa-solid fa-trash" />
                            </button>
                          </td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
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

export default SchoolsIndex;
