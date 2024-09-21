import React, { ReactElement, ReactNode } from 'react';
import classNames from 'classnames';

export type FormGroupProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  style?: 'default' | 'checkbox' | 'radio' | 'switch';
  inline?: boolean;
};

const FormGroup = (props: FormGroupProps): ReactElement => {
  const { children, id, className, style = 'default', inline } = props;

  const classes: string[] = ['form-group'];

  if (style === 'checkbox' || style === 'radio') {
    classes.push('form-check');
  }

  if (style === 'switch') {
    classes.push('form-check form-switch');
  }

  if (inline) {
    classes.push('form-check-inline');
  }

  return (
    <div id={id} className={classNames(...classes, className)}>
      {children}
    </div>
  );
};

export default FormGroup;
