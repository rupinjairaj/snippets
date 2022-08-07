import React from 'react';

import { NavLink } from 'react-router-dom';
import { faTags, faCode } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function Navbar() {

    let codeIcon = <FontAwesomeIcon icon={faCode} />
    // let sentIcon = <FontAwesomeIcon icon={faPaperPlane} />
    // let trendIcon = <FontAwesomeIcon icon={faBolt} />
    let tagIcon = <FontAwesomeIcon icon={faTags} />
    // let peopleIcon = <FontAwesomeIcon icon={faUser} />

    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
            <div className="container-fluid">
                <a className="navbar-brand" href="/tags">Snippets</a>
                <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarToggler" aria-controls="navbarToggler" aria-expanded="false" aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarToggler">
                    <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                        <li className="nav-item">
                            <NavLink className="nav-link" to="/tags">{tagIcon} Tags</NavLink>
                        </li>
                        {/* <li className="nav-item">
                            <NavLink className="nav-link" to="/publish">{codeIcon} Publish</NavLink>
                        </li> */}
                        {/* <li className="nav-item">
                            <NavLink className="nav-link" to="/sent">{sentIcon} Sent</NavLink>
                        </li>
                        <li className="nav-item">
                            <NavLink className="nav-link" to="/trending">{trendIcon} Trending</NavLink>
                        </li>
                        <li className="nav-item">
                            <NavLink className="nav-link" to="/people">{peopleIcon} People</NavLink>
                        </li> */}
                    </ul>
                </div>
            </div>
        </nav>
    );
}

export default Navbar;
