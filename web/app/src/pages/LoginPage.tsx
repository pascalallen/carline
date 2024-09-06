import React, { FormEvent, ReactElement, useEffect } from 'react';
import { observer } from 'mobx-react-lite';
import { Helmet } from 'react-helmet-async';
import { useLocation } from 'react-router';
import { useNavigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';
import Footer from '@components/Footer';

export type LoginFormValues = {
  email_address: string;
  password: string;
};

type LocationState = { from?: Location };

const LoginPage = observer((): ReactElement => {
  const authService = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (authService.isLoggedIn()) {
      navigate(Path.WALKER, { replace: true });
    }
  }, [authService, navigate]);

  const location = useLocation();
  const state: LocationState = location.state as LocationState;

  const handleLogin = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const emailAddress = formData.get('email_address')?.toString() ?? '';
    const password = formData.get('password')?.toString() ?? '';
    await authService.login({ email_address: emailAddress, password });
    const from = state?.from?.pathname || Path.WALKER;
    navigate(from, { replace: true });
  };

  return (
    <div id="login-page" className="login-page d-flex flex-column vh-100">
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
});

export default LoginPage;
