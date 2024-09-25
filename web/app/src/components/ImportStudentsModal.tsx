import React, { FormEvent } from 'react';
import { Modal } from 'react-bootstrap';

type Props = {
  show: boolean;
  onClose: () => void;
  onImport: (event: FormEvent<HTMLFormElement>) => Promise<void>;
  isImporting: boolean;
};

const ImportStudentsModal = (props: Props): React.ReactElement => {
  const { show, onClose, onImport, isImporting } = props;

  return (
    <Modal
      className="import-students-modal d-flex justify-content-center align-items-center"
      show={show}
      onHide={onClose}>
      <Modal.Header className="border-0">
        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={onClose} />
      </Modal.Header>
      <Modal.Body className="d-flex flex-column align-items-center pt-0">
        <h5 className="import-students-header">Import Students</h5>
        <form id="import-students-form" className="import-students-form" onSubmit={onImport}>
          <div className="mb-3">
            <label htmlFor="students-file" className="form-label">
              File input
            </label>
            <input className="form-control" type="file" id="students-file" name="file" />
          </div>
          <div className="form-group text-center">
            <button
              id="import-students-button"
              className="import-students-button btn btn-primary"
              type="submit"
              disabled={isImporting}>
              Import
            </button>
          </div>
        </form>
      </Modal.Body>
    </Modal>
  );
};

export default ImportStudentsModal;
