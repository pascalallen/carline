import React, { ChangeEvent, FocusEvent, ReactElement, ReactNode, InputHTMLAttributes } from 'react';
import classNames from 'classnames';
import { StyleButtonVariants } from '@constants/StyleButtonVariants';

export type FileUploadProps = InputHTMLAttributes<HTMLInputElement> & {
  children?: ReactNode;
  id?: string;
  name?: string;
  className?: string;
  value?: string | string[] | number;
  placeholder?: string;
  tabIndex?: number;
  isValid?: boolean;
  required?: boolean;
  disabled?: boolean;
  variant?: StyleButtonVariants;
  ariaDescribedBy?: string;
  onChange?: (event: ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: FocusEvent<HTMLInputElement>) => void;
};

const FileUpload = (props: FileUploadProps): ReactElement => {
  const {
    children,
    id,
    name,
    className,
    value,
    placeholder,
    tabIndex,
    required,
    disabled,
    variant,
    onChange,
    onBlur,
    ...remaining
  } = props;

  return (
    <span className={classNames('btn btn-file', variant ? `btn-${variant}` : '')}>
      {children}
      <input
        id={id}
        className={className}
        type="file"
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
    </span>
  );
};

export default FileUpload;
