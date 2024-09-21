import React, { ChangeEvent, FocusEvent, ReactElement, InputHTMLAttributes } from 'react';
import classNames from 'classnames';

export type InputProps = InputHTMLAttributes<HTMLInputElement> & {
  id?: string;
  name?: string;
  type: 'text' | 'password' | 'email' | 'color' | 'date' | 'number' | 'search' | 'tel' | 'url';
  className?: string;
  value?: string | string[] | number;
  placeholder?: string;
  tabIndex?: number;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const Input = (props: InputProps): ReactElement => {
  const {
    id,
    name,
    type,
    className,
    value,
    placeholder,
    tabIndex,
    isValid = true,
    required,
    disabled,
    onChange,
    onBlur,
    ...remaining
  } = props;

  return (
    <input
      id={id}
      className={classNames('form-control', className, !isValid ? 'is-invalid' : '')}
      type={type}
      name={name}
      value={value}
      placeholder={placeholder}
      required={required}
      tabIndex={tabIndex}
      disabled={disabled}
      onChange={onChange}
      onBlur={onBlur}
      {...remaining}
    />
  );
};

export default Input;
