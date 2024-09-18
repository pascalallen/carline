import React, { useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { School } from '@domain/types/School';
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

const SchoolsPage = (): React.ReactElement => {
  const [loading, setLoading] = useState(initialState.loading);
  const [schools, setSchools] = useState(initialState.schools);
  const [showAddSchoolModal, setShowAddSchoolModal] = useState(initialState.showAddSchoolModal);
  const [addingSchool, setAddingSchool] = useState(initialState.addingSchool);
  const [selectedSchool, setSelectedSchool] = useState(initialState.selectedSchool);
  const [showRemoveSchoolModal, setShowRemoveSchoolModal] = useState(initialState.showRemoveSchoolModal);
  const [removingSchool, setRemovingSchool] = useState(initialState.removingSchool);

  useEffect(() => {
    // TODO: call service to fetch schools
    setLoading(initialState.loading);
    setSchools([
      {
        id: '1',
        name: 'Weiss Elementary'
      },
      {
        id: '2',
        name: 'Hutto Elementary'
      }
    ]);
    setLoading(false);
  }, []);

  const handleShowAddSchoolModal = (): void => {
    setShowAddSchoolModal(true);
  };

  const handleHideAddSchoolModal = (): void => {
    setShowAddSchoolModal(initialState.showAddSchoolModal);
  };

  const handleAddSchool = async (): Promise<void> => {
    try {
      setAddingSchool(true);
      // TODO: call service to add school and trigger re-fetching of schools
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
      if (!selectedSchool?.id) {
        // TODO: throw error if this function is triggered without having set a school ID
      }
      // TODO: call service to remove school
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
                          <td>{school.id}</td>
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

export default SchoolsPage;
