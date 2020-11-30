import React from "react";
import {connect} from "react-redux"

import {search} from "../../actions";
import "./Search.css";

class Search extends React.Component {

    search = (e) => {
        e.preventDefault()
        const value = document.getElementById("movie-title").value;
        if (value.trim() !== "") {
            this.props.search(value);
        }
    }

    render() {
        return (
            <div className="search">
                <form className="ui form" onSubmit={this.search}>
                    <div className="ui two column grid container">
                        <div className="row">
                            <div className="fifteen wide column">
                                <div className="ui form">
                                    <div className="field">
                                        <input id="movie-title" type="text" placeholder="Movie Title"/>
                                    </div>
                                </div>
                            </div>
                            <div className="one wide column">
                                <button className="ui button" type="submit">Search</button>
                            </div>
                        </div>
                    </div>
                </form>
            </div>
        );
    }
}

export default connect(null, {search})(Search);