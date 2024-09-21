import React, { ChangeEvent, FocusEvent, ReactElement, ReactNode } from 'react';
import classNames from 'classnames';
import { StyleButtonVariants } from '@constants/StyleButtonVariants';

export type RadioButtonProps = {
  children?: ReactNode;
  id: string;
  name?: string;
  className?: string;
  variant?: StyleButtonVariants;
  value?: string | string[] | number;
  tabIndex?: number;
  isValid?: boolean;
  checked?: boolean;
  required?: boolean;
  disabled?: boolean;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const RadioButton = (props: RadioButtonProps): ReactElement => {
  const {
    children,
    id,
    name,
    className,
    variant = 'primary',
    value,
    tabIndex,
    isValid = true,
    checked,
    required,
    disabled,
    onChange,
    onBlur
  } = props;

  return (
    <>
      <input
        id={id}
        className={classNames('btn-check', className, !isValid ? 'is-invalid' : '')}
        type="radio"
        autoComplete="off"
        name={name}
        value={value}
        checked={checked}
        required={required}
        tabIndex={tabIndex}
        disabled={disabled}
        onChange={onChange}
        onBlur={onBlur}
      />
      <label className={classNames('btn', `btn-${variant}`)} htmlFor={id}>
        {children}
      </label>
    </>
  );
};

export default RadioButton;
