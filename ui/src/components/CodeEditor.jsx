import React from "react";

import Editor from '@monaco-editor/react';

const CodeEditor = ({ language, handleCodeChange, theme, defaultValue, readOnly }) => {

    return (
        <div className="container">
            <Editor
                options={{ readOnly }}
                height={"70vh"}
                width={'100%'}
                language={language || "python"}
                // value={code}
                theme={theme}
                defaultValue={defaultValue}
                onChange={handleCodeChange}
            />
        </div>
    );
};

export default CodeEditor;
