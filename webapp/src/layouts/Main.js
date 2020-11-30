import React from "react";

import Search from "../components/Search/Search";
import Results from "../components/Results/Results";

const MainPage = () => {
    return (
        <div className="ui one column grid container">
            <div className="row">
                <div className="column">
                    <Search/>
                </div>
            </div>
            <div className="row">
                <div className="column">
                    <Results/>
                </div>
            </div>
        </div>
    );
}

export default MainPage;