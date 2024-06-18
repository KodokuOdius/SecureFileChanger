import React, { useEffect, useState } from "react";

const Loader = ({ msg }) => {
    const delayMs = 500
    const [isShow, setIsShow] = useState(false)

    useEffect(() => {
        const timer = setTimeout(() => {
            setIsShow(true)
        }, delayMs)

        return () => clearTimeout(timer)
    }, [delayMs])

    return (
        <>
            {isShow &&
                <div className="loading__panel">
                    <div className="loader__body">
                        <span class="loader"></span>
                        <h3 className="loading__msg">
                            {msg}
                        </h3>
                    </div>
                </div>
            }
        </>
    )
}

export default Loader