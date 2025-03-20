import {createRouter, createWebHistory} from 'vue-router';
import HomeView from '../views/HomeView.vue';
import Jeux from "@/views/Jeux.vue";
import Livres from "@/views/Livres.vue";

const routes = [
    { path: '/', component: HomeView },
    { path: '/livres',component: Livres },
     { path: '/jeux', component: Jeux },
    // { path: '/connexion', component: () => import('../views/ConnexionView.vue') },
];

const router = createRouter({
    history: createWebHistory("/static"),
    routes,
});

export default router;