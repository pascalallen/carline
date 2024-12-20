import React, { FormEvent, ReactElement, useEffect, useState } from 'react';
import { Form, InputControl } from '@pascalallen/react-form-components';
import { FormGroup } from 'react-bootstrap';
import { Helmet } from 'react-helmet-async';
import { useNavigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';
import { ErrorApiResponse, FailApiResponse } from '@services/ApiService';
import Footer from '@components/Footer';
import Toast from '@components/Toast';

export type RegisterFormValues = {
  first_name: string;
  last_name: string;
  email_address: string;
};

type State = {
  firstName: string;
  lastName: string;
  emailAddress: string;
  touched: {
    firstName: boolean;
    lastName: boolean;
    emailAddress: boolean;
  };
  errorMessage: string;
};

const initialState: State = {
  firstName: '',
  lastName: '',
  emailAddress: '',
  touched: {
    firstName: false,
    lastName: false,
    emailAddress: false
  },
  errorMessage: ''
};

const Register = (): ReactElement => {
  const authService = useAuth();
  const navigate = useNavigate();

  const [firstName, setFirstName] = useState(initialState.firstName);
  const [lastName, setLastName] = useState(initialState.lastName);
  const [emailAddress, setEmailAddress] = useState(initialState.emailAddress);
  const [touched, setTouched] = useState(initialState.touched);
  const [errorMessage, setErrorMessage] = useState(initialState.errorMessage);

  useEffect(() => {
    if (authService.isLoggedIn()) {
      navigate(Path.SCHOOLS, { replace: true });
    }
  }, [authService, navigate]);

  const handleRegister = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
    event.preventDefault();
    setErrorMessage(initialState.errorMessage);
    try {
      await authService.register({
        first_name: firstName,
        last_name: lastName,
        email_address: emailAddress
      });
      // TODO: Redirect to "check your email" success page
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
    <div id="register-page" className="register-page d-flex flex-column vh-100">
      <div className="toast-container top-0 end-0 p-3">
        {errorMessage && <Toast className="text-bg-danger">{errorMessage}</Toast>}
      </div>
      <Helmet>
        <title>CarLine - Register</title>
      </Helmet>
      <main className="container flex-fill mt-3">
        <section>
          <div className="row row-cols-auto justify-content-center">
            <div className="col">
              <h1>Register</h1>
              <Form id="register-form" className="register-form" onSubmit={handleRegister}>
                <InputControl
                  inputId="first-name"
                  className="first-name mb-3"
                  name="first_name"
                  type="text"
                  label="First name"
                  tabIndex={1}
                  value={firstName}
                  isValid={touched.firstName ? firstName.length > 0 : true}
                  required
                  autoFocus
                  error={touched.firstName && firstName.length < 1 ? 'First name is required' : undefined}
                  onChange={e => setFirstName(e.target.value)}
                  onBlur={() => handleBlur('firstName')}
                />
                <InputControl
                  inputId="last-name"
                  className="last-name mb-3"
                  name="last_name"
                  type="text"
                  label="Last name"
                  tabIndex={2}
                  value={lastName}
                  isValid={touched.lastName ? lastName.length > 0 : true}
                  required
                  error={touched.lastName && lastName.length < 1 ? 'Last name is required' : undefined}
                  onChange={e => setLastName(e.target.value)}
                  onBlur={() => handleBlur('lastName')}
                />
                <InputControl
                  inputId="email-address"
                  className="email-address mb-3"
                  name="email_address"
                  type="email"
                  label="Email address"
                  tabIndex={3}
                  value={emailAddress}
                  isValid={touched.emailAddress ? emailAddress.length > 0 : true}
                  required
                  error={touched.emailAddress && emailAddress.length < 1 ? 'Email address is required' : undefined}
                  onChange={e => setEmailAddress(e.target.value)}
                  onBlur={() => handleBlur('emailAddress')}
                />
                <FormGroup className="mb-3">
                  <button id="register-button" className="register-button btn btn-primary" type="submit" tabIndex={6}>
                    Register
                  </button>
                </FormGroup>
                <FormGroup>
                  <p>
                    Already have an account?{' '}
                    <a
                      href={Path.LOGIN}
                      onClick={event => {
                        event.preventDefault();
                        navigate(Path.LOGIN);
                      }}>
                      Log in
                    </a>
                  </p>
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

export default Register;
