import React, { ReactElement, ReactNode } from 'react';
import classNames from 'classnames';

export type FeedbackProps = {
  children?: ReactNode;
  id?: string;
  className?: string;
  isValid?: boolean;
};

const Feedback = (props: FeedbackProps): ReactElement => {
  const { children, id, className, isValid = true } = props;

  return (
    <div id={id} className={classNames(isValid ? 'valid-feedback' : 'invalid-feedback', className)}>
      {children}
    </div>
  );
};

export default Feedback;
