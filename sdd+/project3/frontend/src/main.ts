import { createApp } from "vue";
import { createPinia } from "pinia";
import PrimeVue from "primevue/config";
import "./assets/css/main.css";
// @ts-ignore
import presets from "./assets/css/primevue-presets.js";
import App from "./App.vue";
import { router } from "./router";

const app = createApp(App);
const pinia = createPinia();

app.use(pinia);
app.use(router);
app.use(PrimeVue, { unstyled: true, pt: presets });

app.mount("#app");
