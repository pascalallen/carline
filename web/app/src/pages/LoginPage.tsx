import React, { FormEvent, ReactElement, useEffect, useState } from 'react';
import { Helmet } from 'react-helmet-async';
import { useLocation } from 'react-router';
import { useNavigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';
import { ErrorApiResponse, FailApiResponse } from '@services/ApiService';
import Footer from '@components/Footer';
import Toast from '@components/Toast';

export type LoginFormValues = {
  email_address: string;
  password: string;
};

type LocationState = { from?: Location };

const LoginPage = (): ReactElement => {
  const authService = useAuth();
  const navigate = useNavigate();
  const location = useLocation();
  const state: LocationState = location.state as LocationState;

  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    if (authService.isLoggedIn()) {
      navigate(Path.WALKER, { replace: true });
    }
  }, [authService, navigate]);

  const handleLogin = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    try {
      const formData = new FormData(event.currentTarget);
      const emailAddress = formData.get('email_address')?.toString() ?? '';
      const password = formData.get('password')?.toString() ?? '';
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
              <form id="login-form" className="login-form" onSubmit={handleLogin}>
                <div className="mb-3">
                  <label htmlFor="email-address" className="form-label">
                    Email address
                  </label>
                  <input id="email-address" className="email-address form-control" type="email" name="email_address" />
                </div>
                <div className="mb-3">
                  <label htmlFor="password" className="form-label">
                    Password
                  </label>
                  <input id="password" className="password form-control" type="password" name="password" />
                </div>
                <div className="form-group">
                  <button id="login-button" className="login-button btn btn-primary" type="submit">
                    Login
                  </button>
                </div>
              </form>
            </div>
          </div>
        </section>
      </main>
      <Footer />
    </div>
  );
};

export default LoginPage;
