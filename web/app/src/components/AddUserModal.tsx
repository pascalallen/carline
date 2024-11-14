import React, { FormEvent } from 'react';
import { Modal } from 'react-bootstrap';

type Props = {
  show: boolean;
  onClose: () => void;
  onAdd: (event: FormEvent<HTMLFormElement>) => Promise<void>;
  isAdding: boolean;
};

const AddUserModal = (props: Props): React.ReactElement => {
  const { show, onClose, onAdd, isAdding } = props;

  return (
    <Modal className="add-user-modal d-flex justify-content-center align-items-center" show={show} onHide={onClose}>
      <Modal.Header className="border-0">
        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={onClose} />
      </Modal.Header>
      <Modal.Body className="text-center d-flex flex-column align-items-center pt-0">
        <h5 className="add-user-header">Add User</h5>
        <form id="add-user-form" className="add-user-form" onSubmit={onAdd}>
          <div className="mb-3">
            <label htmlFor="user-first-name" className="form-label">
              First name
            </label>
            <input id="user-first-name" className="user-first-name form-control" type="text" name="first_name" />
          </div>
          <div className="mb-3">
            <label htmlFor="user-last-name" className="form-label">
              Last name
            </label>
            <input id="user-last-name" className="user-last-name form-control" type="text" name="last_name" />
          </div>
          <div className="mb-3">
            <label htmlFor="user-email-address" className="form-label">
              Email address
            </label>
            <input
              id="user-email-address"
              className="user-email-address form-control"
              type="email"
              name="email_address"
            />
          </div>
          <div className="form-group">
            <button id="add-user-button" className="add-user-button btn btn-primary" type="submit" disabled={isAdding}>
              Add User
            </button>
          </div>
        </form>
      </Modal.Body>
    </Modal>
  );
};

export default AddUserModal;
