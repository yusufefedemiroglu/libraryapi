import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:3000",
});

API.interceptors.request.use((config) => {
  const userdata = localStorage.getItem("userdata");
  if (userdata) {
    const token = JSON.parse(userdata).token
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default API;
