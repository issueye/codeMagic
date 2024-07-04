import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/home/index.vue";
import Root from "./views/Root.vue";
import Dataset from './views/dataset/index.vue';
import Template from './views/template/index.vue';
import Login from "./views/Login.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "root", component: Root },
    { path: "/login", name: "login", component: Login },
    { path: "/home", name: "home", component: Home },
    { path: "/datasets", component: Dataset },
    { path: "/templates", component: Template },
    { path: "/setup", component: () => import("./views/Setup.vue") },
  ],
});

export default router;
