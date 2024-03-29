import axios from "axios";

export const apiUrl = process.env.REACT_APP_API_URL || "http://localhost:3000";

export async function GetStatus() {
    return axios.get(apiUrl + "/status");
}

