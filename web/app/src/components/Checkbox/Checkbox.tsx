import React, { ChangeEvent, FocusEvent, ReactElement, useEffect, useRef } from 'react';
import classNames from 'classnames';

export type CheckboxProps = {
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  tabIndex?: number;
  isValid?: boolean;
  checked?: boolean | 'indeterminate';
  required?: boolean;
  disabled?: boolean;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const Checkbox = (props: CheckboxProps): ReactElement => {
  const { id, name, className, value, tabIndex, isValid = true, checked, required, disabled, onChange, onBlur } = props;

  const checkboxElement = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (checkboxElement.current !== null) {
      checkboxElement.current.indeterminate = checked === 'indeterminate';
    }
  }, [checked]);

  return (
    <input
      id={id}
      className={classNames('form-check-input', className, !isValid ? 'is-invalid' : '')}
      ref={checkboxElement}
      type="checkbox"
      name={name}
      value={value}
      checked={checked === 'indeterminate' || checked}
      required={required}
      tabIndex={tabIndex}
      disabled={disabled}
      onChange={onChange}
      onBlur={onBlur}
    />
  );
};

export default Checkbox;
