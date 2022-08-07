import React, { useState } from "react";
import CodeEditor from "./CodeEditor";
import LanguagesDropdown from "./LanguagesDropdown";
import { faPlus } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { getURL, snippetPath } from "../settings/api";


const Publish = () => {

    const [code, setCode] = useState("");
    const [defaultValue] = useState("# some comment");
    const [language, setLanguage] = useState("");
    const [tags, setTags] = useState([]);
    const [tag, setTag] = useState("");
    const [title, setTitle] = useState("");

    let plusIcon = <FontAwesomeIcon icon={faPlus} />

    const onSelectChangeLanguage = (e) => {
        setLanguage(e.target.value);
    };

    const onChangeTitleInput = (e) => {
        setTitle(e.target.value);
    };

    const handleCodeChange = (val) => {
        setCode(val);
    };

    const onChangeTagsInput = (e) => {
        setTag(e.target.value);
    };

    const handleAddTag = (e) => {
        e.preventDefault();
        setTags([...tags, tag])
        setTag("");
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (tags.length === 0 || code === "" || title === "") {
            return;
        }
        let payload = {
            "name": title,
            "tags": tags,
            "content": btoa(code)
        };
        console.log(payload);
        const url = getURL(window.location.hostname);
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload)
        };
        fetch(`${url + snippetPath}`, requestOptions)
            .then(response => response.json())
            .then(data => {
                if (data === null) data = [];
                setTags(data)
            })
    };

    return (
        <div className="container">
            <div className="row" style={{ marginBottom: "12px" }}>
                <div className="col-3">
                    <LanguagesDropdown onSelectChange={onSelectChangeLanguage} />
                </div>
                <div className="col-6">
                    <div className="input-group mb-3">
                        <input
                            type="text"
                            className="form-control"
                            placeholder="title"
                            onChange={onChangeTitleInput}
                            value={title} />
                    </div>
                </div>
            </div>
            <div className="row" style={{ marginBottom: "12px" }}>
                <CodeEditor
                    language={language}
                    handleCodeChange={handleCodeChange}
                    defaultValue={defaultValue} />
            </div>
            <div className="row" style={{ marginBottom: "12px" }}>
                {
                    tags.map((tag, idx) => (
                        <div key={idx} className="col">
                            <span className="badge bg-primary rounded-pill">{tag}</span>
                        </div>
                    ))
                }
            </div>
            <div className="row">
                <div className="col-3">
                    <div className="input-group mb-3">
                        <input
                            type="text"
                            className="form-control"
                            placeholder="tags"
                            aria-label="tags"
                            onChange={onChangeTagsInput}
                            value={tag} />
                    </div>
                </div>
                <div className="col-3">
                    <button type="button"
                        className="btn btn-outline-secondary"
                        onClick={handleAddTag}>{plusIcon}</button>
                </div>
            </div>
            <div className="row">
                <div className="col-3">
                    <button type="button"
                        onClick={handleSubmit}
                        className="btn btn-outline-dark">Submit</button>
                </div>
            </div>
        </div>
    );
};

export default Publish;
