import React, { useEffect, useState } from 'react';
import { useHistory, useParams } from 'react-router';
import Prism from 'prismjs';
import { faArrowLeft, faTags, faCode, faArrowUp, faArrowDown } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import './SnippetView.css';


function SnippetView() {

    let params = useParams();
    let history = useHistory();

    let [snippets, setSnippets] = useState([]);

    // useEffect hook is used to achieve the 
    // componentDidMount lifecycle. This is required
    // to trigger Prims to look at the <pre>/<code>
    // html element and syntax highlight the code with using the 
    // css class that specifies the target highlight language 
    useEffect(() => {
        Prism.highlightAll();
        fetch(`http://localhost:9000/snippet/${params.tagName}`)
            .then(response => response.json())
            .then(data => setSnippets(data))
    }, []);

    function backClick(e) {
        e.preventDefault();
        history.push("/inbox");
    }

    let backIcon = <FontAwesomeIcon icon={faArrowLeft} />
    let codeIcon = <FontAwesomeIcon icon={faCode} />
    let upvoteIcon = <FontAwesomeIcon icon={faArrowUp} />
    let downvoteIcon = <FontAwesomeIcon icon={faArrowDown} />

    return (
        <div className="row">
            {/* <ul className="list-group list-group-horizontal"> */}
            {/* <button onClick={backClick} className="list-group-item">{backIcon}</button> */}
            {/* <li className="list-group-item">{topicIcon} {topicName} { }</li> */}
            {/* </ul> */}
            <h1>{params.tagName}</h1>
            <div className="row row-cols-1 row-cols-md-1 g-4 add-scroll">
                {snippets.length !== 0 && snippets.map((snippet) => (
                    <div className="col" key={snippet.id}>
                        <div className="card">
                            <div className="card-header">
                                {snippet.name} {codeIcon}
                            </div>
                            <div className="card-body">
                                <pre className="lang-clike"
                                    style={{ border: "none", padding: "2px", marginBottom: "10px", marginTop: "1px" }}>
                                    <code>
                                        {Buffer.from(snippet.content, 'base64').toString('utf-8')}
                                    </code>
                                </pre>
                            </div>
                            <div className="card-footer text-muted">
                                {/* {upvoteIcon} 12  {downvoteIcon} 3  By: {data.Author} ({data.DateCreated}) | File name: {data.FileName} | Language: {data.Language */}
                                {/* } */}
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default SnippetView;