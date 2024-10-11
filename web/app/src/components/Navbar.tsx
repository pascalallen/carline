import React, { ReactElement } from 'react';
import { useLocation } from 'react-router';
import { useNavigate } from 'react-router-dom';
import Path from '@domain/constants/Path';
import useAuth from '@hooks/useAuth';

const Navbar = (): ReactElement => {
  const authService = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = async (): Promise<void> => {
    await authService.logout();
    authService.logout().finally(() => navigate(Path.INDEX));
  };

  return (
    <nav className="navbar navbar-expand-lg bg-body-tertiary">
      <div className="container-fluid">
        <a
          className="navbar-brand"
          href={Path.INDEX}
          onClick={event => {
            event.preventDefault();
            navigate(Path.INDEX);
          }}>
          CarLine
        </a>
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbar-content"
          aria-controls="navbar-content"
          aria-expanded="false"
          aria-label="Toggle navigation">
          <i className="fa-solid fa-bars" />
        </button>
        <div id="navbar-content" className="collapse navbar-collapse">
          <ul className="navbar-nav me-auto mb-2 mb-lg-0">
            <li className="nav-item">
              <a
                className={`nav-link ${location.pathname === Path.SCHOOLS ? 'active' : ''}`}
                href={Path.SCHOOLS}
                onClick={event => {
                  event.preventDefault();
                  navigate(Path.SCHOOLS);
                }}>
                Schools
              </a>
            </li>
            {authService.isLoggedIn() && (
              <li className="nav-item">
                <a className="nav-link" onClick={handleLogout} role="button">
                  <i className="fa-solid fa-right-from-bracket"></i>
                </a>
              </li>
            )}
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
