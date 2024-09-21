import React, { ChangeEvent, FocusEvent, ReactElement } from 'react';
import classNames from 'classnames';

export type RangeProps = {
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  tabIndex?: number;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  list?: string;
  max?: string | number;
  min?: string | number;
  step?: string | number;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const Range = (props: RangeProps): ReactElement => {
  const {
    id,
    name,
    className,
    value,
    tabIndex,
    isValid = true,
    required,
    disabled,
    list,
    max,
    min,
    step,
    onChange,
    onBlur
  } = props;

  return (
    <input
      id={id}
      className={classNames('form-range', className, !isValid ? 'is-invalid' : '')}
      type="range"
      name={name}
      value={value}
      required={required}
      tabIndex={tabIndex}
      disabled={disabled}
      list={list}
      max={max}
      min={min}
      step={step}
      onChange={onChange}
      onBlur={onBlur}
    />
  );
};

export default Range;
