import React from 'react';
import { Modal } from 'react-bootstrap';
import { Student } from '@domain/types/Student';

type Props = {
  show: boolean;
  onClose: () => void;
  student?: Student;
  onRemove: () => Promise<void>;
  isRemoving: boolean;
};

const RemoveStudentModal = (props: Props): React.ReactElement => {
  const {
    show,
    onClose,
    student = {
      id: '',
      tag_number: '',
      first_name: '',
      last_name: '',
      created_at: ''
    },
    onRemove,
    isRemoving
  } = props;

  return (
    <Modal
      className="remove-student-modal d-flex justify-content-center align-items-center"
      dialogClassName="w-75 h-auto"
      show={show}
      onHide={onClose}>
      <Modal.Header className="border-0">
        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={onClose} />
      </Modal.Header>
      <Modal.Body className="d-flex flex-column pt-0">
        <h5 className="remove-student-header">Remove Student</h5>
        <p className="remove-student-modal-caption">
          Are you sure you want to remove{' '}
          <strong>
            {student.first_name} {student.last_name}
          </strong>
          ?
        </p>
        <button
          id="remove-student-button"
          type="button"
          className="remove-student-button btn btn-primary font-size-sm font-weight-bold"
          disabled={isRemoving}
          onClick={onRemove}>
          Yes, Remove Student
        </button>
      </Modal.Body>
    </Modal>
  );
};

export default RemoveStudentModal;
