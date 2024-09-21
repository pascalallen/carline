import React, { ChangeEvent, FocusEvent } from 'react';
import Feedback from '@components/Feedback/Feedback';
import FormGroup from '@components/FormGroup/FormGroup';
import FormLabel from '@components/FormLabel/FormLabel';
import FormText from '@components/FormText/FormText';
import Input from '@components/Input/Input';

export type DatePickerControlProps = {
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  label?: string;
  placeholder?: string;
  tabIndex?: number;
  error?: string;
  tip?: string;
  inputId?: string;
  labelId?: string;
  errorId?: string;
  tipId?: string;
  autoFocus?: boolean;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  theme?: {
    input?: string;
    label?: string;
    tip?: string;
    error?: string;
  };
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const DatePickerControl = (props: DatePickerControlProps): React.ReactElement => {
  const {
    id,
    name,
    className,
    value,
    label,
    placeholder,
    tabIndex,
    error,
    tip,
    inputId,
    labelId,
    errorId,
    tipId,
    autoFocus,
    isValid,
    required,
    disabled,
    theme = {},
    onChange,
    onBlur
  } = props;

  return (
    <FormGroup id={id} style="default" className={className}>
      {label && (
        <FormLabel id={labelId} className={theme.label} style="default" htmlFor={inputId} required={required}>
          {label}
        </FormLabel>
      )}
      <Input
        id={inputId}
        name={name}
        className={theme.input}
        type="date"
        value={value}
        placeholder={placeholder}
        tabIndex={tabIndex}
        autoFocus={autoFocus}
        isValid={isValid}
        required={required}
        disabled={disabled}
        onChange={onChange}
        onBlur={onBlur}
      />
      {error && (
        <Feedback id={errorId} className={theme.error} isValid={isValid}>
          {error}
        </Feedback>
      )}
      {!error && tip && (
        <FormText id={tipId} className={theme.tip} block>
          {tip}
        </FormText>
      )}
    </FormGroup>
  );
};

export default DatePickerControl;
