import React, { FormEvent, ReactElement, useEffect, useState } from 'react';
import { FormGroup } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useLocation } from 'react-router';
import { useNavigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';
import { ErrorApiResponse, FailApiResponse } from '@services/ApiService';
import Footer from '@components/Footer';
import Form from '@components/Form/Form';
import InputControl from '@components/InputControl/InputControl';
import Toast from '@components/Toast';

export type LoginFormValues = {
  email_address: string;
  password: string;
};

type LocationState = { from?: Location };

type State = {
  emailAddress: string;
  password: string;
  touched: {
    emailAddress: boolean;
    password: boolean;
  };
  errorMessage: string;
};

const initialState: State = {
  emailAddress: '',
  password: '',
  touched: {
    emailAddress: false,
    password: false
  },
  errorMessage: ''
};

const LoginPage = (): ReactElement => {
  const authService = useAuth();
  const navigate = useNavigate();
  const location = useLocation();
  const state: LocationState = location.state as LocationState;

  const [emailAddress, setEmailAddress] = useState(initialState.emailAddress);
  const [password, setPassword] = useState(initialState.password);
  const [touched, setTouched] = useState(initialState.touched);
  const [errorMessage, setErrorMessage] = useState(initialState.errorMessage);

  useEffect(() => {
    if (authService.isLoggedIn()) {
      navigate(Path.WALKER, { replace: true });
    }
  }, [authService, navigate]);

  const handleLogin = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    setErrorMessage(initialState.errorMessage);
    try {
      await authService.login({ email_address: emailAddress, password });
      const from = state?.from?.pathname || Path.WALKER;
      navigate(from, { replace: true });
    } catch (error) {
      if ((error as FailApiResponse)?.statusCode === 400) {
        setErrorMessage('Validation error');
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
    <div id="login-page" className="login-page d-flex flex-column vh-100">
      <div className="toast-container top-0 end-0 p-3">
        {errorMessage && <Toast className="text-bg-danger">{errorMessage}</Toast>}
      </div>
      <Helmet>
        <title>Carline - Login</title>
      </Helmet>
      <main className="container flex-fill mt-3">
        <section>
          <div className="row row-cols-auto justify-content-center">
            <div className="col">
              <h1>Login</h1>
              <Form id="login-form" className="login-form" onSubmit={handleLogin}>
                <InputControl
                  inputId="email-address"
                  className="email-address mb-3"
                  name="email_address"
                  type="email"
                  label="Email address"
                  tabIndex={1}
                  value={emailAddress}
                  isValid={touched.emailAddress ? emailAddress.length > 0 : true}
                  required
                  autoFocus
                  error={touched.emailAddress && emailAddress.length < 1 ? 'Email address is required' : undefined}
                  onChange={e => setEmailAddress(e.target.value)}
                  onBlur={() => handleBlur('emailAddress')}
                />
                <InputControl
                  inputId="password"
                  className="password mb-3"
                  name="password"
                  type="password"
                  label="Password"
                  tabIndex={2}
                  value={password}
                  isValid={touched.password ? password.length > 0 : true}
                  required
                  error={touched.password && password.length < 1 ? 'Password is required' : undefined}
                  onChange={e => setPassword(e.target.value)}
                  onBlur={() => handleBlur('password')}
                />
                <FormGroup>
                  <button id="login-button" className="login-button btn btn-primary" type="submit" tabIndex={3}>
                    Login
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

export default LoginPage;
