import React from "react";
import {connect} from "react-redux";

import Result from "./Result";
import "./Style.css";

class Results extends React.Component {

    render() {
        if (this.props.titles.length === 0) {
            return "";
        }
        return (
            <div className="results">
                <div className="ui four column grid container">
                    {
                        this.props.titles.map(title => {
                            return <Result key={title.id} title={title}/>
                        })
                    }
                </div>
            </div>
        );
    }
}

const mapStateToProps = state => {
    return {
        titles: state.titles
    }
};

export default connect(mapStateToProps, null)(Results);