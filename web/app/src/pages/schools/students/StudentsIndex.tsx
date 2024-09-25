import React, { FormEvent, useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useNavigate, useParams } from 'react-router-dom';
import { DomainEvents } from '@domain/constants/DomainEvents';
import { Student } from '@domain/types/Student';
import useEvent from '@hooks/useEvents';
import useStudentService from '@hooks/useStudentService';
import { DomainEvent } from '@services/eventDispatcher';
import Footer from '@components/Footer';
import ImportStudentsModal from '@components/ImportStudentsModal';
import Navbar from '@components/Navbar';
import RemoveStudentModal from '@components/RemoveStudentModal';

type State = {
  loading: boolean;
  students: Student[];
  showImportStudentsModal: boolean;
  importingStudents: boolean;
  selectedStudent?: Student;
  showRemoveStudentModal: boolean;
  removingStudent: boolean;
};

const initialState: State = {
  loading: true,
  students: [],
  showImportStudentsModal: false,
  importingStudents: false,
  showRemoveStudentModal: false,
  removingStudent: false
};

const StudentsIndex = (): React.ReactElement => {
  const studentService = useStudentService();
  const { schoolId } = useParams();
  const navigate = useNavigate();

  const [loading, setLoading] = useState(initialState.loading);
  const [students, setStudents] = useState(initialState.students);
  const [showImportStudentsModal, setShowImportStudentsModal] = useState(initialState.showImportStudentsModal);
  const [importingStudents, setImportingStudents] = useState(initialState.importingStudents);
  const [selectedStudent, setSelectedStudent] = useState(initialState.selectedStudent);
  const [showRemoveStudentModal, setShowRemoveStudentModal] = useState(initialState.showRemoveStudentModal);
  const [removingStudent, setRemovingStudent] = useState(initialState.removingStudent);

  const studentsImportedEvent: DomainEvent | undefined = useEvent(DomainEvents.STUDENTS_IMPORTED);
  const studentRemovedEvent: DomainEvent | undefined = useEvent(DomainEvents.STUDENT_REMOVED);

  const fetchStudents = async (): Promise<void> => {
    setLoading(initialState.loading);

    try {
      const students = await studentService.getAll(schoolId ?? '', { include_deleted: false });
      setStudents(students);
    } catch (error) {
      console.error(error);
    }

    setLoading(false);
  };

  useEffect(() => {
    fetchStudents();
  }, []);

  useEffect(() => {
    if (studentsImportedEvent?.id) {
      fetchStudents();
    }
  }, [studentsImportedEvent, studentRemovedEvent]);

  const handleShowImportStudentsModal = (): void => setShowImportStudentsModal(true);
  const handleHideImportStudentsModal = (): void => setShowImportStudentsModal(initialState.showImportStudentsModal);

  const handleImportStudents = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();

    try {
      setImportingStudents(true);
      const formData = new FormData(event.currentTarget);
      await studentService.import(schoolId ?? '', formData);
      setImportingStudents(initialState.importingStudents);
      handleHideImportStudentsModal();
    } catch (error) {
      handleHideImportStudentsModal();
      setImportingStudents(initialState.importingStudents);
    }
  };

  const handleShowRemoveStudentModal = (student: Student): void => {
    setShowRemoveStudentModal(true);
    setSelectedStudent(student);
  };

  const handleHideRemoveStudentModal = (): void => {
    setShowRemoveStudentModal(initialState.showRemoveStudentModal);
    setSelectedStudent(initialState.selectedStudent);
  };

  const handleRemoveStudent = async (): Promise<void> => {
    try {
      setRemovingStudent(true);
      await studentService.remove(schoolId ?? '', selectedStudent?.id ?? '');
      setStudents(prevStudents => prevStudents.filter(student => student.id !== selectedStudent?.id));
      setRemovingStudent(initialState.removingStudent);
      handleHideRemoveStudentModal();
    } catch (error) {
      handleHideRemoveStudentModal();
      setRemovingStudent(initialState.removingStudent);
    }
  };

  return (
    <div id="students-page" className="students-page d-flex flex-column vh-100">
      <Helmet>
        <title>Carline - Students</title>
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
                  <h1>Students</h1>
                  <p>Use this page to manage (create, update, and delete) your students.</p>
                </div>
              </div>
            </section>
            <section>
              {showImportStudentsModal ? (
                <ImportStudentsModal
                  show={showImportStudentsModal}
                  onClose={handleHideImportStudentsModal}
                  onImport={handleImportStudents}
                  isImporting={importingStudents}
                />
              ) : null}
              {showRemoveStudentModal ? (
                <RemoveStudentModal
                  show={showRemoveStudentModal}
                  onClose={handleHideRemoveStudentModal}
                  student={selectedStudent}
                  onRemove={handleRemoveStudent}
                  isRemoving={removingStudent}
                />
              ) : null}
              <div className="row">
                <div className="col">
                  <table className="table">
                    <thead>
                      <tr className="align-middle">
                        <th scope="col">ID</th>
                        <th scope="col">Tag Number</th>
                        <th scope="col">First Name</th>
                        <th scope="col">Last Name</th>
                        <th scope="col">School ID</th>
                        <th scope="col">
                          <button
                            type="button"
                            className="btn btn-success"
                            onClick={() => handleShowImportStudentsModal()}>
                            <i className="fa-solid fa-plus" />
                          </button>
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      {students.map((student: Student, index: number) => (
                        <tr key={`student-row-${index}`} className="align-middle">
                          <td>{student.id}</td>
                          <td>{student.tag_number}</td>
                          <td>{student.first_name}</td>
                          <td>{student.last_name}</td>
                          <td>
                            <a
                              href={`/schools/${student.school_id}`}
                              onClick={event => {
                                event.preventDefault();
                                navigate(`/schools/${student.school_id}`);
                              }}>
                              {student.school_id}
                            </a>
                          </td>

                          <td>
                            <button
                              type="button"
                              className="btn btn-danger"
                              onClick={() => handleShowRemoveStudentModal(student)}>
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

export default StudentsIndex;
