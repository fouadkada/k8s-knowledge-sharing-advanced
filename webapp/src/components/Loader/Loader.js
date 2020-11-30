import React from "react";
import {connect} from "react-redux";

import "./Loader.css";

class Loader extends React.Component {

    render() {
        const loading = (this.props.loading) ? "active" : "";
        return (
            <div className="loader">
                <div className={`ui ${loading} dimmer`}>
                    <div className="ui text loader">Loading</div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = state => {
    return {
        loading: state.loading.loading
    }
}

export default connect(mapStateToProps, null)(Loader);