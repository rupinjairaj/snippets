import React from 'react';

import { NavLink } from 'react-router-dom';
import { faInbox, faPaperPlane, faBolt, faTags, faUser } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function Navbar() {

    let inboxIcon = <FontAwesomeIcon icon={faInbox} />
    let sentIcon = <FontAwesomeIcon icon={faPaperPlane} />
    let trendIcon = <FontAwesomeIcon icon={faBolt} />
    let tagIcon = <FontAwesomeIcon icon={faTags} />
    let peopleIcon = <FontAwesomeIcon icon={faUser} />

    return (
        <nav class="navbar navbar-expand-lg navbar-light bg-light">
            <div class="container-fluid">
                <a class="navbar-brand" href="#">Snippets</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarToggler" aria-controls="navbarToggler" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarToggler">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <NavLink className="nav-link" to="/tags">{tagIcon} Tags</NavLink>
                        </li>
                        <li class="nav-item">
                            <NavLink className="nav-link" to="/inbox">{inboxIcon} Inbox</NavLink>
                        </li>
                        <li class="nav-item">
                            <NavLink className="nav-link" to="/sent">{sentIcon} Sent</NavLink>
                        </li>
                        <li class="nav-item">
                            <NavLink className="nav-link" to="/trending">{trendIcon} Trending</NavLink>
                        </li>
                        <li class="nav-item">
                            <NavLink className="nav-link" to="/people">{peopleIcon} People</NavLink>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    );
}

export default Navbar;
