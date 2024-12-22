import React, { useState, useEffect } from 'react';

function App() {
    const [message, setMessage] = useState('');

    useEffect(() => {
        fetch("http://localhost:8332/hello")
            .then((response) => response.text())
            .then((data) => setMessage(data));
    }, []);

    return <div>{message || "Loading..."}</div>;
}

export default App;
