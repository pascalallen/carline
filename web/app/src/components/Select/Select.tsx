import React, { ChangeEvent, FocusEvent, ReactElement, ReactNode } from 'react';
import classNames from 'classnames';

export type SelectProps = {
  children?: ReactNode;
  id?: string;
  name?: string;
  className?: string;
  defaultValue?: string | string[] | number;
  value?: string | string[] | number;
  tabIndex?: number;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  onChange?: (event: ChangeEvent<HTMLSelectElement>) => void;
  onBlur?: (event: FocusEvent<HTMLSelectElement>) => void;
};

const Select = (props: SelectProps): ReactElement => {
  const {
    children,
    id,
    name,
    className,
    defaultValue,
    value,
    tabIndex,
    isValid = true,
    required,
    disabled,
    onChange,
    onBlur
  } = props;

  return (
    <select
      id={id}
      className={classNames('form-control', className, !isValid ? 'is-invalid' : '')}
      name={name}
      defaultValue={defaultValue}
      value={value}
      required={required}
      tabIndex={tabIndex}
      disabled={disabled}
      onChange={onChange}
      onBlur={onBlur}>
      {children}
    </select>
  );
};

export default Select;
