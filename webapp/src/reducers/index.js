import {combineReducers} from "redux";

import {CLEAR_RECOMMENDATIONS, RECOMMENDATION_SUCCESS, SEARCH_SUCCESS, START_NETWORK, STOP_NETWORK} from "../actions";

const INITIAL_MOVIES_STATE = []
const titlesReducer = (state = INITIAL_MOVIES_STATE, action) => {
    switch (action.type) {
        case SEARCH_SUCCESS:
            return action.payload;
        default:
            return state;
    }
};

const INITIAL_RECOMMENDATIONS_STATE = []
const recommendationsReducer = (state = INITIAL_RECOMMENDATIONS_STATE, action) => {
    switch (action.type) {
        case RECOMMENDATION_SUCCESS:
            return action.payload;
        case CLEAR_RECOMMENDATIONS:
            return INITIAL_RECOMMENDATIONS_STATE;
        default:
            return state;
    }
};


const INITIAL_LOADING_STATE = {loading: false};
const loadingReducer = (state = INITIAL_LOADING_STATE, action) => {
    switch (action.type) {
        case START_NETWORK:
            return {...state, loading: true};
        case STOP_NETWORK:
            return {...state, loading: false};
        default:
            return state;
    }
};

export default combineReducers({
    titles: titlesReducer,
    recommendations: recommendationsReducer,
    loading: loadingReducer,
});