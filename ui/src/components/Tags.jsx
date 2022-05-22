import React, { useEffect, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { feedData } from '../data/test';

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
        fetch('http://localhost:9000/tag')
            .then(response => response.json())
            .then(data => setTags(data))
    }, []);

    return (
        <div className="container">
            <div className="list-group">
                {tags.length !== 0 &&
                    tags.map((tag) => {
                        return (
                            <li className="list-group-item list-group-item-action"
                                key={tag.id}
                                onClick={(e) => inboxItemClickHandler(tag.name, e)}>
                                <div className="row" >
                                    <div className="col-3">
                                        <b>{tag.name}</b>
                                        <span style={{ marginLeft: "4px" }}
                                            className="badge bg-primary rounded-pill">+{10}</span>
                                    </div>
                                    <div className="col-6">
                                        <em>{"Some description"}</em>
                                    </div>
                                </div>
                            </li>
                        )
                    })
                }
            </div>
        </div>
    );
}

export default Tags;
