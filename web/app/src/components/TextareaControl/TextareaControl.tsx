import React, { ChangeEvent, FocusEvent, ReactElement, ReactNode } from 'react';
import Feedback from '@components/Feedback/Feedback';
import FormGroup from '@components/FormGroup/FormGroup';
import FormLabel from '@components/FormLabel/FormLabel';
import FormText from '@components/FormText/FormText';
import Textarea from '@components/Textarea/Textarea';

export type TextareaControlProps = {
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  label?: string;
  placeholder?: string;
  rows?: number;
  tabIndex?: number;
  error?: string;
  tip?: ReactNode;
  textareaId?: string;
  labelId?: string;
  errorId?: string;
  tipId?: string;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  theme?: {
    textarea?: string;
    label?: string;
    tip?: string;
    error?: string;
  };
  onChange?: (event: ChangeEvent<HTMLTextAreaElement>) => void;
  onBlur?: (event: FocusEvent<HTMLTextAreaElement>) => void;
};

const TextareaControl = (props: TextareaControlProps): ReactElement => {
  const {
    id,
    name,
    className,
    value,
    label,
    placeholder,
    rows,
    tabIndex,
    error,
    tip,
    textareaId,
    labelId,
    errorId,
    tipId,
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
        <FormLabel id={labelId} className={theme.label} style="default" htmlFor={textareaId} required={required}>
          {label}
        </FormLabel>
      )}
      <Textarea
        id={textareaId}
        name={name}
        className={theme.textarea}
        value={value}
        placeholder={placeholder}
        rows={rows}
        tabIndex={tabIndex}
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

export default TextareaControl;
