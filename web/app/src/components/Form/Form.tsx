import React, { FormEvent, ReactElement, ReactNode } from 'react';

export type FormProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  autoComplete?: boolean;
  noValidate?: boolean;
  onSubmit?: (event: FormEvent<HTMLFormElement>) => void;
};

const Form = (props: FormProps): ReactElement => {
  const { children, id, className, autoComplete = false, noValidate = true, onSubmit } = props;

  return (
    <form
      id={id}
      className={className}
      autoComplete={autoComplete ? 'on' : 'off'}
      noValidate={noValidate}
      onSubmit={onSubmit}>
      {children}
    </form>
  );
};

export default Form;
