import React, { FormEvent, useEffect, useState } from 'react';
import { Spinner } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useNavigate, useParams } from 'react-router-dom';
import { DomainEvents } from '@domain/constants/DomainEvents';
import { User } from '@domain/types/User';
import useEvent from '@hooks/useEvents';
import useUserService from '@hooks/useUserService';
import { ErrorApiResponse, FailApiResponse } from '@services/ApiService';
import { DomainEvent } from '@services/eventDispatcher';
import AddUserModal from '@components/AddUserModal';
import Footer from '@components/Footer';
import Navbar from '@components/Navbar';
import RemoveUserModal from '@components/RemoveUserModal';
import Toast from '@components/Toast';

type State = {
  loading: boolean;
  users: User[];
  showAddUserModal: boolean;
  addingUser: boolean;
  selectedUser?: User;
  showRemoveUserModal: boolean;
  removingUser: boolean;
  errorMessage: string;
};

const initialState: State = {
  loading: true,
  users: [],
  showAddUserModal: false,
  addingUser: false,
  showRemoveUserModal: false,
  removingUser: false,
  errorMessage: ''
};

const UsersIndex = (): React.ReactElement => {
  const userService = useUserService();
  const { schoolId } = useParams();
  const navigate = useNavigate();

  const [loading, setLoading] = useState(initialState.loading);
  const [users, setUsers] = useState(initialState.users);
  const [showAddUserModal, setShowAddUserModal] = useState(initialState.showAddUserModal);
  const [addingUser, setAddingUser] = useState(initialState.addingUser);
  const [selectedUser, setSelectedUser] = useState(initialState.selectedUser);
  const [showRemoveUserModal, setShowRemoveUserModal] = useState(initialState.showRemoveUserModal);
  const [removingUser, setRemovingUser] = useState(initialState.removingUser);
  const [errorMessage, setErrorMessage] = useState(initialState.errorMessage);

  const userAddedEvent: DomainEvent | undefined = useEvent(DomainEvents.USER_ADDED);
  const userRemovedEvent: DomainEvent | undefined = useEvent(DomainEvents.USER_REMOVED);

  useEffect(() => {
    setLoading(initialState.loading);
    userService
      .getAll(schoolId ?? '')
      .then((users: User[]) => setUsers(users))
      .catch(error => setErrorMessage(error))
      .finally(() => setLoading(false));
  }, [schoolId]);

  useEffect(() => {
    if (userAddedEvent?.id) {
      setLoading(initialState.loading);
      userService
        .getAll(schoolId ?? '')
        .then((users: User[]) => setUsers(users))
        .catch(error => setErrorMessage(error))
        .finally(() => setLoading(false));
    }
  }, [userAddedEvent, userRemovedEvent, schoolId]);

  const handleShowAddUserModal = (): void => setShowAddUserModal(true);
  const handleHideAddUserModal = (): void => setShowAddUserModal(initialState.showAddUserModal);

  const handleAddUser = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    setErrorMessage(initialState.errorMessage);

    try {
      setAddingUser(true);
      const formData = new FormData(event.currentTarget);
      const firstName = formData.get('first_name')?.toString() ?? '';
      const lastName = formData.get('last_name')?.toString() ?? '';
      const emailAddress = formData.get('email_address')?.toString() ?? '';
      await userService.add(schoolId ?? '', {
        first_name: firstName,
        last_name: lastName,
        email_address: emailAddress
      });
      setAddingUser(initialState.addingUser);
      handleHideAddUserModal();
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
      handleHideAddUserModal();
      setAddingUser(initialState.addingUser);
    }
  };

  const handleShowRemoveUserModal = (user: User): void => {
    setShowRemoveUserModal(true);
    setSelectedUser(user);
  };

  const handleHideRemoveUserModal = (): void => {
    setShowRemoveUserModal(initialState.showRemoveUserModal);
    setSelectedUser(initialState.selectedUser);
  };

  const handleRemoveUser = async (): Promise<void> => {
    try {
      setRemovingUser(true);
      await userService.remove(schoolId ?? '', selectedUser?.id ?? '');
      setUsers(prevUsers => prevUsers.filter(user => user.id !== selectedUser?.id));
      setRemovingUser(initialState.removingUser);
      handleHideRemoveUserModal();
    } catch (error) {
      handleHideRemoveUserModal();
      setRemovingUser(initialState.removingUser);
    }
  };

  return (
    <div id="users-page" className="users-page d-flex flex-column vh-100">
      <div className="toast-container top-0 end-0 p-3">
        {errorMessage && <Toast className="text-bg-danger">{errorMessage}</Toast>}
      </div>
      <Helmet>
        <title>CarLine - Users</title>
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
                  <h1>Users</h1>
                  <p>Use this page to manage (create, update, and delete) your users.</p>
                </div>
              </div>
            </section>
            <section>
              {showAddUserModal ? (
                <AddUserModal
                  show={showAddUserModal}
                  onClose={handleHideAddUserModal}
                  onAdd={handleAddUser}
                  isAdding={addingUser}
                />
              ) : null}
              {showRemoveUserModal ? (
                <RemoveUserModal
                  show={showRemoveUserModal}
                  onClose={handleHideRemoveUserModal}
                  user={selectedUser}
                  onRemove={handleRemoveUser}
                  isRemoving={removingUser}
                />
              ) : null}
              <div className="row">
                <div className="col">
                  <table className="table">
                    <thead>
                      <tr className="align-middle">
                        <th scope="col">ID</th>
                        <th scope="col">First Name</th>
                        <th scope="col">Last Name</th>
                        <th scope="col">Email Address</th>
                        <th scope="col">School ID</th>
                        <th scope="col">
                          <button type="button" className="btn btn-success" onClick={() => handleShowAddUserModal()}>
                            <i className="fa-solid fa-plus" />
                          </button>
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      {users.map((user: User, index: number) => (
                        <tr key={`user-row-${index}`} className="align-middle">
                          <td>{user.id}</td>
                          <td>{user.first_name}</td>
                          <td>{user.last_name}</td>
                          <td>{user.email_address}</td>
                          <td>
                            <a
                              href={`/schools/${user.school_id}`}
                              onClick={event => {
                                event.preventDefault();
                                navigate(`/schools/${user.school_id}`);
                              }}>
                              {user.school_id}
                            </a>
                          </td>
                          <td>
                            <button
                              type="button"
                              className="btn btn-danger"
                              onClick={() => handleShowRemoveUserModal(user)}>
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

export default UsersIndex;
