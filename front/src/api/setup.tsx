//set variable apiUrl to the url of the backend using the environment variable REACT_APP_API_URL

import {apiUrl, GetStatus} from "./api";
import axios from "axios";

export function GetSetup() {
    return axios.get(apiUrl + "/setup");
}

export function IsSetupRequired() {
    const status = GetStatus();

    return status.then((response) => {
        if (response.status === 200) {
            return false;
        }
    }).catch((error) => {
        if (error.response.status === 503) {
            return true;
        }
    });
}

export function CurrentSetupStep() {
    const setup = GetSetup();

    return setup.then((response) => {
        console.log(response.data);

        return response.data.nextStep;
    });
}