import React, { FormEvent, ReactElement, useEffect, useState } from 'react';
import { Form, InputControl } from '@pascalallen/react-form-components';
import { FormGroup } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useNavigate, useParams } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';
import { ErrorApiResponse, FailApiResponse } from '@services/ApiService';
import Footer from '@components/Footer';
import Toast from '@components/Toast';

export type ActivateFormValues = {
  token: string;
  password: string;
  confirm_password: string;
};

type State = {
  password: string;
  confirmPassword: string;
  touched: {
    password: boolean;
    confirmPassword: boolean;
  };
  errorMessage: string;
};

const initialState: State = {
  password: '',
  confirmPassword: '',
  touched: {
    password: false,
    confirmPassword: false
  },
  errorMessage: ''
};

const Activate = (): ReactElement => {
  const authService = useAuth();
  const navigate = useNavigate();
  const { token } = useParams();

  const [password, setPassword] = useState(initialState.password);
  const [confirmPassword, setConfirmPassword] = useState(initialState.confirmPassword);
  const [touched, setTouched] = useState(initialState.touched);
  const [errorMessage, setErrorMessage] = useState(initialState.errorMessage);

  useEffect(() => {
    if (authService.isLoggedIn()) {
      navigate(Path.SCHOOLS, { replace: true });
    }
  }, [authService, navigate]);

  const handleActivate = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    setErrorMessage(initialState.errorMessage);
    try {
      await authService.activate({
        token: token ?? '',
        password: password,
        confirm_password: confirmPassword
      });
    } catch (error) {
      if ((error as FailApiResponse)?.statusCode === 400) {
        setErrorMessage('Validation error');
      }

      if ((error as FailApiResponse)?.statusCode === 401) {
        setErrorMessage('Invalid credentials');
      }

      if ((error as ErrorApiResponse)?.statusCode === 422) {
        setErrorMessage((error as ErrorApiResponse).body.message);
      }
    }
  };

  const handleBlur = (field: string) => {
    setTouched({
      ...touched,
      [field]: true
    });
  };

  return (
    <div id="activate-page" className="activate-page d-flex flex-column vh-100">
      <div className="toast-container top-0 end-0 p-3">
        {errorMessage && <Toast className="text-bg-danger">{errorMessage}</Toast>}
      </div>
      <Helmet>
        <title>CarLine - Activate Account</title>
      </Helmet>
      <main className="container flex-fill mt-3">
        <section>
          <div className="row row-cols-auto justify-content-center">
            <div className="col">
              <h1>Activate Account</h1>
              <Form id="activate-form" className="activate-form" onSubmit={handleActivate}>
                <input name="token" type="hidden" value={token} required />
                <InputControl
                  inputId="password"
                  className="password mb-3"
                  name="password"
                  type="password"
                  label="Password"
                  tabIndex={1}
                  value={password}
                  isValid={touched.password ? password.length > 0 : true}
                  required
                  autoFocus
                  error={touched.password && password.length < 1 ? 'Password is required' : undefined}
                  onChange={e => setPassword(e.target.value)}
                  onBlur={() => handleBlur('password')}
                />
                <InputControl
                  inputId="confirm-password"
                  className="confirm-password mb-3"
                  name="confirm_password"
                  type="password"
                  label="Confirm password"
                  tabIndex={2}
                  value={confirmPassword}
                  isValid={touched.confirmPassword ? confirmPassword.length > 0 : true}
                  required
                  error={
                    touched.confirmPassword && confirmPassword.length < 1
                      ? 'Password confirmation is required'
                      : undefined
                  }
                  onChange={e => setConfirmPassword(e.target.value)}
                  onBlur={() => handleBlur('confirmPassword')}
                />
                <FormGroup className="mb-3">
                  <button id="activate-button" className="activate-button btn btn-primary" type="submit" tabIndex={3}>
                    Activate
                  </button>
                </FormGroup>
              </Form>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default Activate;
