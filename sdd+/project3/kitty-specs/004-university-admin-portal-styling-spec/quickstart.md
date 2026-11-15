# Quickstart: Implementing the Styling System

This guide provides developers with the necessary steps to apply the new styling system to the University Admin Portal.

## Prerequisites

- Vue 3 project with Tailwind CSS and PostCSS installed.
- PrimeVue installed.

## 1. Configure Tailwind CSS

Copy the proposed configuration from `contracts/tailwind-config.json` into your project's `tailwind.config.js` file.

```bash
# Ensure your tailwind.config.js includes the custom design tokens
```

## 2. Set Up PrimeVue Unstyled Mode

Configure PrimeVue to use unstyled mode in your main entry file (e.g., `main.js` or `main.ts`).

```javascript
import { createApp } from 'vue';
import PrimeVue from 'primevue/config';
import App from './App.vue';

const app = createApp(App);

app.use(PrimeVue, {
    unstyled: true
});

app.mount('#app');
```

## 3. Implement Component Presets

To style PrimeVue components using Tailwind, use the `pt` (Pass Through) property. Create a centralized preset file to avoid repetition.

**Example: Styling a Button**

In `src/assets/css/primevue-presets.js`:

```javascript
export const myButtonPreset = {
    root: ({ props }) => ({
        class: [
            'px-4 py-2 rounded-md font-medium transition-colors duration-200',
            props.severity === 'primary' ? 'bg-primary-500 text-white hover:bg-primary-600' : '',
            props.disabled ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer'
        ]
    })
};
```

Then apply it globally or per component:

```javascript
// Global application in main.js
app.use(PrimeVue, {
    unstyled: true,
    pt: {
        button: myButtonPreset
    }
});
```

## 4. Verifying Accessibility

After implementing a component, verify its accessibility using the following checklist:
- [ ] Does it have a visible focus ring?
- [ ] Does it meet the required color contrast (4.5:1)?
- [ ] Is it navigable via keyboard?

Use **Lighthouse** or **Axe-core** for automated validation.
