import api from "../api/movies";

export const START_NETWORK = "START_NETWORK";
export const STOP_NETWORK = "STOP_NETWORK";

export const SEARCH_SUCCESS = "SEARCH_SUCCESS";
export const SEARCH_FAILURE = "SEARCH_FAILURE";

export const search = (title) => async dispatch => {
    try {
        dispatch({type: START_NETWORK});
        const response = await api.get(`/search?title=${title}`);
        dispatch({type: SEARCH_SUCCESS, payload: response.data});
        dispatch({type: STOP_NETWORK});
    } catch (e) {
        dispatch({type: SEARCH_FAILURE});
        dispatch({type: STOP_NETWORK});
    }
};

export const RECOMMENDATION_SUCCESS = "RECOMMENDATION_SUCCESS";
export const RECOMMENDATION_FAILURE = "RECOMMENDATION_FAILURE";

export const getRecommendations = (movieID, cb) => async dispatch => {
    try {
        dispatch({type: START_NETWORK});
        const response = await api.get(`/recommendations/${movieID}`);
        dispatch({type: RECOMMENDATION_SUCCESS, payload: response.data});
        dispatch({type: STOP_NETWORK});
        cb && cb(true)
    } catch (e) {
        dispatch({type: RECOMMENDATION_FAILURE});
        dispatch({type: STOP_NETWORK});
        cb && cb(false)
    }
};

export const CLEAR_RECOMMENDATIONS = "CLEAR_RECOMMENDATIONS";

export const clearRecommendations = () => dispatch => {
    dispatch({type: CLEAR_RECOMMENDATIONS});
};