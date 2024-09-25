import React, { FormEvent, useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useNavigate } from 'react-router-dom';
import { DomainEvents } from '@domain/constants/DomainEvents';
import Path from '@domain/constants/Path';
import { School } from '@domain/types/School';
import useEvent from '@hooks/useEvents';
import useSchoolService from '@hooks/useSchoolService';
import { DomainEvent } from '@services/eventDispatcher';
import AddSchoolModal from '@components/AddSchoolModal';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';
import RemoveSchoolModal from '@components/RemoveSchoolModal';

type State = {
  loading: boolean;
  schools: School[];
  showAddSchoolModal: boolean;
  addingSchool: boolean;
  selectedSchool?: School;
  showRemoveSchoolModal: boolean;
  removingSchool: boolean;
};

const initialState: State = {
  loading: true,
  schools: [],
  showAddSchoolModal: false,
  addingSchool: false,
  showRemoveSchoolModal: false,
  removingSchool: false
};

const SchoolsIndex = (): React.ReactElement => {
  const schoolService = useSchoolService();
  const navigate = useNavigate();

  const [loading, setLoading] = useState(initialState.loading);
  const [schools, setSchools] = useState(initialState.schools);
  const [showAddSchoolModal, setShowAddSchoolModal] = useState(initialState.showAddSchoolModal);
  const [addingSchool, setAddingSchool] = useState(initialState.addingSchool);
  const [selectedSchool, setSelectedSchool] = useState(initialState.selectedSchool);
  const [showRemoveSchoolModal, setShowRemoveSchoolModal] = useState(initialState.showRemoveSchoolModal);
  const [removingSchool, setRemovingSchool] = useState(initialState.removingSchool);

  const schoolAddedEvent: DomainEvent | undefined = useEvent(DomainEvents.SCHOOL_ADDED);
  const schoolRemovedEvent: DomainEvent | undefined = useEvent(DomainEvents.SCHOOL_REMOVED);

  const fetchSchools = async (): Promise<void> => {
    setLoading(initialState.loading);

    try {
      const schools = await schoolService.getAll();
      setSchools(schools);
    } catch (error) {
      console.error(error);
    }

    setLoading(false);
  };

  useEffect(() => {
    fetchSchools();
  }, []);

  useEffect(() => {
    if (schoolAddedEvent?.id) {
      fetchSchools();
    }
  }, [schoolAddedEvent, schoolRemovedEvent]);

  const handleShowAddSchoolModal = (): void => setShowAddSchoolModal(true);
  const handleHideAddSchoolModal = (): void => setShowAddSchoolModal(initialState.showAddSchoolModal);

  const handleAddSchool = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();

    try {
      setAddingSchool(true);
      const formData = new FormData(event.currentTarget);
      const name = formData.get('name')?.toString() ?? '';
      await schoolService.add({ name });
      setAddingSchool(initialState.addingSchool);
      handleHideAddSchoolModal();
    } catch (error) {
      handleHideAddSchoolModal();
      setAddingSchool(initialState.addingSchool);
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
      await schoolService.remove(selectedSchool?.id ?? '');
      setSchools(prevSchools => prevSchools.filter(school => school.id !== selectedSchool?.id));
      setRemovingSchool(initialState.removingSchool);
      handleHideRemoveSchoolModal();
    } catch (error) {
      handleHideRemoveSchoolModal();
      setRemovingSchool(initialState.removingSchool);
    }
  };

  return (
    <div id="schools-page" className="schools-page d-flex flex-column vh-100">
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
                                navigate(`/schools/${school.id}/students`);
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
