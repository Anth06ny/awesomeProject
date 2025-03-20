import axios from "axios";

const apiClient = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL, // Récupère l'URL depuis .env
    headers: {
        "Content-Type": "application/json",
    },
});

export default apiClient;