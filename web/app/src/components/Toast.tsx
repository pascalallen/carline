import React, { ReactNode } from 'react';

export type ToastProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
};

const Toast = (props: ToastProps) => {
  const { children, id, className } = props;

  return (
    <div
      id={id}
      className={`toast align-items-center border-0 show ${className}`}
      role="alert"
      aria-live="assertive"
      aria-atomic="true">
      <div className="d-flex">
        <div className="toast-body">{children}</div>
        <button
          type="button"
          className="btn-close btn-close-white me-2 m-auto"
          data-bs-dismiss="toast"
          aria-label="Close"
        />
      </div>
    </div>
  );
};

export default Toast;
