import React from "react";
import {BrowserRouter, Route, Switch} from "react-router-dom";

import Menu from "./components/Menu/Menu";
import Loader from "./components/Loader/Loader";
import MainPage from "./layouts/Main";
import ResultPage from "./layouts/Recommendations";

const App = () => {
    return (
        <div className="ui container">
            <Menu/>
            <BrowserRouter>
                <Switch>
                    <Route path="/" exact={true}>
                        <MainPage/>
                    </Route>
                    <Route path="/recommendations/:movie_id" exact={true}>
                        <ResultPage/>
                    </Route>
                </Switch>
            </BrowserRouter>
            <Loader/>
        </div>
    );
}

export default App;
