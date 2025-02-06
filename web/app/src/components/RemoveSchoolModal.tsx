import React from 'react';
import { Modal } from 'react-bootstrap';
import { School } from '@domain/types/School';

type Props = {
  show: boolean;
  onClose: () => void;
  school?: School;
  onRemove: () => Promise<void>;
  isRemoving: boolean;
};

const RemoveSchoolModal = (props: Props): React.ReactElement => {
  const {
    show,
    onClose,
    school = {
      id: '',
      name: ''
    },
    onRemove,
    isRemoving
  } = props;

  return (
    <Modal
      className="remove-school-modal d-flex justify-content-center align-items-center"
      dialogClassName="w-75 h-auto"
      show={show}
      onHide={onClose}>
      <Modal.Header className="border-0">
        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={onClose} />
      </Modal.Header>
      <Modal.Body className="d-flex flex-column pt-0">
        <h5 className="remove-school-header">Remove School</h5>
        <p className="remove-school-modal-caption">
          Are you sure you want to remove <strong>{school.name}</strong>? Students associated with this school will be
          detached but not removed.
        </p>
        <button
          id="remove-school-button"
          type="button"
          className="remove-school-button btn btn-primary font-size-sm font-weight-bold"
          disabled={isRemoving}
          onClick={onRemove}>
          Yes, Remove School
        </button>
      </Modal.Body>
    </Modal>
  );
};

export default RemoveSchoolModal;
