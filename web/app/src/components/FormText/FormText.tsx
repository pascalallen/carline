import React, { ReactElement, ReactNode } from 'react';
import classNames from 'classnames';

export type FormTextProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  block?: boolean;
};

const FormText = (props: FormTextProps): ReactElement => {
  const { children, id, className, block = true } = props;

  return React.createElement(
    block ? 'div' : 'span',
    { id: id, className: classNames('form-text', className) },
    children
  );
};

export default FormText;
