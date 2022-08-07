import React from "react";

import { languageOptions } from "../constants/languageOptions";

const LanguagesDropdown = ({ onSelectChange }) => {

    let optionsList = languageOptions.map((x) => (
        <option key={x.id} value={x.value}>{x.name}</option>
    ))

    return (
        <select className="form-select" aria-label="Default select example" onChange={(e) => onSelectChange(e)}>
            {optionsList}
        </select>
    )
};

export default LanguagesDropdown;
