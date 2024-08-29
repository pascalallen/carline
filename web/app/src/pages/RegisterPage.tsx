import React, { FormEvent, ReactElement, useEffect } from 'react';
import { observer } from 'mobx-react-lite';
import { Helmet } from 'react-helmet-async';
import { useLocation } from 'react-router';
import { useNavigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';
import Footer from '@components/Footer';

export type RegisterFormValues = {
  first_name: string;
  last_name: string;
  email_address: string;
  password: string;
  confirm_password: string;
};

type LocationState = { from?: Location };

const RegisterPage = observer((): ReactElement => {
  const authService = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (authService.isLoggedIn()) {
      navigate(Path.WALKER, { replace: true });
    }
  }, [authService, navigate]);

  const location = useLocation();
  const state: LocationState = location.state as LocationState;

  const handleRegister = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const firstName = formData.get('first_name')?.toString() ?? '';
    const lastName = formData.get('last_name')?.toString() ?? '';
    const emailAddress = formData.get('email_address')?.toString() ?? '';
    const password = formData.get('password')?.toString() ?? '';
    const confirmPassword = formData.get('confirm_password')?.toString() ?? '';
    await authService.register({
      first_name: firstName,
      last_name: lastName,
      email_address: emailAddress,
      password,
      confirm_password: confirmPassword
    });
    await authService.login({ email_address: emailAddress, password });
    const from = state?.from?.pathname || Path.WALKER;
    navigate(from, { replace: true });
  };

  return (
    <div id="register-page" className="register-page">
      <Helmet>
        <title>Carline - Register</title>
      </Helmet>
      <main className="container vh-100 mt-3">
        <section>
          <div className="row row-cols-auto justify-content-center">
            <div className="col">
              <h1>Register</h1>
              <form id="register-form" className="register-form" onSubmit={handleRegister}>
                <div className="mb-3">
                  <label htmlFor="first-name" className="form-label">
                    First name
                  </label>
                  <input id="first-name" className="first-name form-control" type="text" name="first_name" />
                </div>
                <div className="mb-3">
                  <label htmlFor="last-name" className="form-label">
                    Last name
                  </label>
                  <input id="last-name" className="last-name form-control" type="text" name="last_name" />
                </div>
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
                <div className="mb-3">
                  <label htmlFor="confirm-password" className="form-label">
                    Confirm password
                  </label>
                  <input
                    id="confirm-password"
                    className="confirm-password form-control"
                    type="password"
                    name="confirm_password"
                  />
                </div>
                <div className="form-group">
                  <button id="register-button" className="register-button btn btn-primary" type="submit">
                    Register
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

export default RegisterPage;
