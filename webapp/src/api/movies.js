import axios from "axios";

import {Constants} from "../Constants";

const instance = axios.create({
    baseURL: `${Constants.BaseURL}/movies`
});

export default instance;