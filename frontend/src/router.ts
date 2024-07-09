import { createRouter, createWebHistory } from "vue-router";
import Home from "./views/home/index.vue";
import Root from "./views/Root.vue";
import Dataset from './views/dataset/index.vue';
import DataModel from './views/data_model/index.vue';
import MakeCode from './views/make_code/index.vue';
import Template from './views/template/index.vue';
import ScriptEdit from './views/script_edit/index.vue';
import Login from "./views/Login.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "root", component: Root },
    { path: "/login", name: "login", component: Login },
    { path: "/home", name: "home", component: Home },
    { path: "/datasets", component: Dataset },
    { path: "/data_model", component: DataModel, meta: { transition: 'slide-left' }, },
    { path: "/make_code", component: MakeCode, meta: { transition: 'slide-left' }, },
    { path: "/templates", component: Template },
    { path: "/script_edit", component: ScriptEdit },
    { path: "/setup", component: () => import("./views/Setup.vue") },
  ],
});

export default router;
