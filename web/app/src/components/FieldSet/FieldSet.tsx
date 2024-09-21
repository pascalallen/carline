import React, { ReactElement, ReactNode } from 'react';

export type FieldSetProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  legend?: string;
  disabled?: boolean;
};

const FieldSet = (props: FieldSetProps): ReactElement => {
  const { children, id, className, legend, disabled } = props;

  return (
    <fieldset id={id} className={className} disabled={disabled}>
      {legend && <legend>{legend}</legend>}
      {children}
    </fieldset>
  );
};

export default FieldSet;
