import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate} from "react-router-dom";
import Room from './components/Room.js';
import Main from './components/Main.js';
import './components/css/styles.css';

class App extends React.Component {

    render() {
        return (
            <>
                <Router>
                    <Routes>
                        <Route
                            path={"/"}
                            element={<Main />} />
                        <Route
                            path={"/room"}
                            element={<Navigate to={"/"} replace />} />
                        <Route
                            path={"/room/*"}
                            element={<Room />} />
                        <Route
                            path={"*"}
                            element={<Navigate to="/" replace />} />
                    </Routes>
                </Router>
            </>
        )
    }
}

export default App;