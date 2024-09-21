import React, { ReactElement, ReactNode } from 'react';
import classNames from 'classnames';

export type FormLabelProps = {
  children?: ReactNode;
  id?: string;
  htmlFor?: string;
  className?: string;
  style?: 'default' | 'checkbox' | 'radio' | 'switch';
  required?: boolean;
};

const FormLabel = (props: FormLabelProps): ReactElement => {
  const { children, id, htmlFor, className, style = 'default', required } = props;

  const classes: string[] = ['form-label'];

  if (style !== 'default') {
    classes.push('form-check-label');
  }

  return (
    <label id={id} className={classNames(...classes, className)} htmlFor={htmlFor}>
      {children}
      {required && <span className="text-danger">*</span>}
    </label>
  );
};

export default FormLabel;
