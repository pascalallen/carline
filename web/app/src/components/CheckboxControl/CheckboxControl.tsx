import React, { ReactElement, ChangeEvent, FocusEvent } from 'react';
import Checkbox from '@components/Checkbox/Checkbox';
import Feedback from '@components/Feedback/Feedback';
import FormGroup from '@components/FormGroup/FormGroup';
import FormLabel from '@components/FormLabel/FormLabel';
import FormText from '@components/FormText/FormText';

export type CheckboxControlProps = {
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  label?: string;
  tabIndex?: number;
  error?: string;
  tip?: string;
  inputId?: string;
  labelId?: string;
  errorId?: string;
  tipId?: string;
  checked?: boolean;
  inline?: boolean;
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

const CheckboxControl = (props: CheckboxControlProps): ReactElement => {
  const {
    id,
    name,
    className,
    value,
    label,
    tabIndex,
    error,
    tip,
    inputId,
    labelId,
    errorId,
    tipId,
    checked,
    inline,
    isValid,
    required,
    disabled,
    theme = {},
    onChange,
    onBlur
  } = props;

  return (
    <FormGroup id={id} style="checkbox" className={className} inline={inline}>
      <Checkbox
        id={inputId}
        name={name}
        className={theme.input}
        value={value}
        tabIndex={tabIndex}
        isValid={isValid}
        checked={checked}
        required={required}
        disabled={disabled}
        onChange={onChange}
        onBlur={onBlur}
      />
      <FormLabel id={labelId} className={theme.label} style="checkbox" htmlFor={inputId} required={required}>
        {label}
      </FormLabel>
      {error && (
        <Feedback id={errorId} className={theme.error} isValid={false}>
          {error}
        </Feedback>
      )}
      {!error && tip && (
        <FormText id={tipId} className={theme.tip}>
          {tip}
        </FormText>
      )}
    </FormGroup>
  );
};

export default CheckboxControl;
