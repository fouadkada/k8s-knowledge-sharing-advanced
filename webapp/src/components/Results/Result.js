import React from "react";
import {connect} from "react-redux";
import {withRouter} from 'react-router-dom';

import {getRecommendations} from "../../actions";
import "./Style.css";

class Result extends React.Component {

    getRecommendations = (movieID) => {
        this.props.getRecommendations(movieID, success => {
            if (success) {
                this.props.history.push(`/recommendations/${movieID}`)
            }
        })
    }

    render() {
        const title = this.props.title
        return (
            <div className="single-result" onClick={() => this.getRecommendations(title.id)}>
                <div className="ui people shape">
                    <div className="sides">
                        <div className="side active">
                            <div className="ui card">
                                <Image title={title}/>
                                <div className="content">
                                    <div className="header">{title.title}</div>
                                    <div className="meta">
                                        {title.release_date.substr(0, 4)}
                                    </div>
                                    <div className="description">
                                        <p>{title.overview}</p>
                                    </div>
                                </div>
                                <div className="extra content">
                                    <span className="right floated">Rated {title.vote_average} / 10</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

const Image = ({title}) => {
    const poster_path = title.poster_path;
    if (poster_path) {
        const postURL = `https://image.tmdb.org/t/p/w780${title.poster_path}`
        return (
            <div className="poster image">
                <img src={postURL} alt="movie poster"/>
            </div>
        );
    } else {
        return (
            <div className="ui placeholder">
                <div className="image"/>
            </div>
        );
    }
}

export default connect(null, {getRecommendations})(withRouter(Result));