import React from "react";
import {connect} from "react-redux";
import {Link, withRouter} from "react-router-dom";

import {clearRecommendations} from "../actions";
import "./Recommendations.css";
import Result from "../components/Results/Result";

class ResultPage extends React.Component {

    componentWillUnmount() {
        this.props.clearRecommendations();
    }

    renderRecommendations = () => {
        if (!this.props.recommendations.length) {
            return (
                <h2>What kinda taste do you have, I could not find any titles to recommend.</h2>
            )
        } else {
            return this.props.recommendations.map(title => {
                return (
                    <div className="column">
                        <Result key={title.id} title={title}/>
                    </div>
                );
            });
        }
    }

    render() {
        return (
            <div className="recommendations">
                <div className="ui three column grid container">
                    <div className="row navigation">
                        <Link to="/">← Back</Link>
                    </div>
                    <div className="row">
                        <div className="column">
                            <h1>Recommendations</h1>
                        </div>
                    </div>
                    <div className="row">
                        {this.renderRecommendations()}
                    </div>
                </div>
            </div>
        );
    }
}

const mapStateToProps = state => {
    return {
        recommendations: state.recommendations
    }
}

export default connect(mapStateToProps, {clearRecommendations})(withRouter(ResultPage));