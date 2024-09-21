import React, { ChangeEvent, FocusEvent, ReactElement } from 'react';
import classNames from 'classnames';

export type RadioProps = {
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  tabIndex?: number;
  isValid?: boolean;
  checked?: boolean;
  required?: boolean;
  disabled?: boolean;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const Radio = (props: RadioProps): ReactElement => {
  const { id, name, className, value, tabIndex, isValid = true, checked, required, disabled, onChange, onBlur } = props;

  return (
    <input
      id={id}
      className={classNames('form-check-input', className, !isValid ? 'is-invalid' : '')}
      type="radio"
      name={name}
      value={value}
      checked={checked}
      required={required}
      tabIndex={tabIndex}
      disabled={disabled}
      onChange={onChange}
      onBlur={onBlur}
    />
  );
};

export default Radio;
