import React, { ChangeEvent, FocusEvent, ReactElement } from 'react';
import classNames from 'classnames';

export type SwitchProps = {
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

const Switch = (props: SwitchProps): ReactElement => {
  const { id, name, className, value, tabIndex, isValid = true, checked, required, disabled, onChange, onBlur } = props;

  return (
    <input
      id={id}
      className={classNames('form-check-input', className, !isValid ? 'is-invalid' : '')}
      type="checkbox"
      role="switch"
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

export default Switch;
