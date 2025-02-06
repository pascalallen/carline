import React, { FormEvent } from 'react';
import { Modal } from 'react-bootstrap';

type Props = {
  show: boolean;
  onClose: () => void;
  onAdd: (event: FormEvent<HTMLFormElement>) => Promise<void>;
  isAdding: boolean;
};

const AddSchoolModal = (props: Props): React.ReactElement => {
  const { show, onClose, onAdd, isAdding } = props;

  return (
    <Modal
      className="add-school-modal d-flex justify-content-center align-items-center"
      dialogClassName="w-75 h-auto"
      show={show}
      onHide={onClose}>
      <Modal.Header className="border-0">
        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={onClose} />
      </Modal.Header>
      <Modal.Body className="d-flex flex-column pt-0">
        <h5 className="add-school-header">Add School</h5>
        <form id="add-school-form" className="add-school-form" onSubmit={onAdd}>
          <div className="mb-3">
            <label htmlFor="school-name" className="form-label">
              School name
            </label>
            <input id="school-name" className="school-name form-control" type="text" name="name" />
          </div>
          <div className="form-group">
            <button
              id="add-school-button"
              className="add-school-button btn btn-primary"
              type="submit"
              disabled={isAdding}>
              Add School
            </button>
          </div>
        </form>
      </Modal.Body>
    </Modal>
  );
};

export default AddSchoolModal;
