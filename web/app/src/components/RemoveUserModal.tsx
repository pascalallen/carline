import React from 'react';
import { Modal } from 'react-bootstrap';
import { User } from '@domain/types/User';

type Props = {
  show: boolean;
  onClose: () => void;
  user?: User;
  onRemove: () => Promise<void>;
  isRemoving: boolean;
};

const RemoveUserModal = (props: Props): React.ReactElement => {
  const {
    show,
    onClose,
    user = {
      id: '',
      first_name: '',
      last_name: '',
      email_address: '',
      created_at: '',
      modified_at: ''
    },
    onRemove,
    isRemoving
  } = props;

  return (
    <Modal
      className="remove-user-modal d-flex justify-content-center align-items-center"
      dialogClassName="w-75 h-auto"
      show={show}
      onHide={onClose}>
      <Modal.Header className="border-0">
        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={onClose} />
      </Modal.Header>
      <Modal.Body className="d-flex flex-column pt-0">
        <h5 className="remove-user-header">Remove User</h5>
        <p className="remove-user-modal-caption">
          Are you sure you want to remove{' '}
          <strong>
            {user.first_name} {user.last_name}
          </strong>
          ?
        </p>
        <button
          id="remove-user-button"
          type="button"
          className="remove-user-button btn btn-primary font-size-sm font-weight-bold"
          disabled={isRemoving}
          onClick={onRemove}>
          Yes, Remove User
        </button>
      </Modal.Body>
    </Modal>
  );
};

export default RemoveUserModal;
