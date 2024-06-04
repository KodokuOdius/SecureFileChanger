import React from "react";

const Loader = ({ msg }) => {
    return (
        <div className="loading__panel">
            <h3 className="loading__msg">
                {msg}
            </h3>
        </div>
    )
}

export default Loader