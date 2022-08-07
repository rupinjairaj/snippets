import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getURL, tagsPath } from "../settings/api";

function Tags() {

    // the useHistory hook returns the current history
    // context that is used to programmatically navigate 
    // to a different component.
    let history = useHistory();
    let [tags, setTags] = useState([]);

    function inboxItemClickHandler(tagName, e) {
        e.preventDefault();
        history.push("/snippetView/" + tagName)
    }

    useEffect(() => {
        const url = getURL(window.location.hostname);
        fetch(`${url + tagsPath}`)
            .then(response => response.json())
            .then(data => {
                if (data === null) data = [];
                setTags(data)
            })
    }, []);

    return (
        <div className="container">
            <div className="list-group">
                {tags.length !== 0 &&
                    tags.map((tag) => {
                        return (
                            <li className="list-group-item d-flex justify-content-between align-items-center"
                                key={tag.id}
                                onClick={(e) => inboxItemClickHandler(tag.name, e)}>
                                <b>{tag.name}</b>
                                {/* <span class="badge bg-primary rounded-pill">14</span> */}
                                <span className="badge rounded-pill bg-light text-dark">{tag.count}</span>
                            </li>
                            // <li className="list-group-item list-group-item-action"
                            //     key={tag.id}
                            //     onClick={(e) => inboxItemClickHandler(tag.name, e)}>
                            //     <div className="row">
                            //         <div className="col-1">
                            //             <span class="badge rounded-pill bg-light text-dark">{tag.count}</span>
                            //         </div>
                            //         <div className="col-3">
                            //             <b>{tag.name}</b>
                            //         </div>
                            //     </div>
                            // </li>
                        )
                    })
                }
            </div>
        </div>
    );
}

export default Tags;
