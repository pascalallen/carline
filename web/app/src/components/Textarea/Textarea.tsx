import React, { ChangeEvent, FocusEvent, ReactElement } from 'react';
import classNames from 'classnames';

export type TextareaProps = {
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  placeholder?: string;
  rows?: number;
  tabIndex?: number;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  onChange?: (event: ChangeEvent<HTMLTextAreaElement>) => void;
  onBlur?: (event: FocusEvent<HTMLTextAreaElement>) => void;
};

const Textarea = (props: TextareaProps): ReactElement => {
  const {
    id,
    name,
    className,
    value,
    placeholder,
    rows,
    tabIndex,
    isValid = true,
    required,
    disabled,
    onChange,
    onBlur
  } = props;

  return (
    <textarea
      id={id}
      className={classNames('form-control', className, !isValid ? 'is-invalid' : '')}
      name={name}
      value={value}
      placeholder={placeholder}
      required={required}
      tabIndex={tabIndex}
      rows={rows}
      disabled={disabled}
      onChange={onChange}
      onBlur={onBlur}
    />
  );
};

export default Textarea;
