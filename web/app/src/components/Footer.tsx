import React, { ReactElement } from 'react';

const Footer = (): ReactElement => {
  return (
    <footer id="footer" className="footer p-3">
      <div className="row justify-content-center align-items-center">
        <div className="col-auto d-none d-sm-block">
          <p className="copyright-desktop mb-0">© 2024 Pascal Allen & Crimson Drive Design LLC</p>
        </div>
        <div className="col-auto d-block d-sm-none">
          <p className="copyright-mobile mb-0">© 2024 Pascal Allen & CDD LLC</p>
        </div>
        <div className="col-auto">
          <a
            href="https://www.termsfeed.com/live/acb0c6cd-2718-465b-9339-7997e07f7ca9"
            target="_blank"
            rel="noreferrer">
            Terms
          </a>
        </div>
        <div className="col-auto">
          <a
            href="https://www.privacypolicies.com/live/47848157-ff99-4a14-a488-f5e0fb1b89af"
            target="_blank"
            rel="noreferrer">
            Privacy
          </a>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
