import React from 'react';

import { NavLink } from 'react-router-dom';
import { faInbox, faPaperPlane, faBolt, faTags, faUser } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function Sidebar() {

    let inboxIcon = <FontAwesomeIcon icon={faInbox} />
    let sentIcon = <FontAwesomeIcon icon={faPaperPlane} />
    let trendIcon = <FontAwesomeIcon icon={faBolt} />
    let tagIcon = <FontAwesomeIcon icon={faTags} />
    let peopleIcon = <FontAwesomeIcon icon={faUser} />

    return (
        <nav className="nav flex-column">
            <NavLink className="nav-link" to="/tags">{tagIcon} Tags</NavLink>
            <NavLink className="nav-link" to="/inbox">{inboxIcon} Inbox</NavLink>
            <NavLink className="nav-link" to="/sent">{sentIcon} Sent</NavLink>
            <NavLink className="nav-link" to="/trending">{trendIcon} Trending</NavLink>
            <NavLink className="nav-link" to="/people">{peopleIcon} People</NavLink>
        </nav >
    );
}

export default Sidebar;
