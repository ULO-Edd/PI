<script setup lang="ts">
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Select from 'primevue/select';
import Checkbox from 'primevue/checkbox';
import RadioButton from 'primevue/radiobutton';
import Card from 'primevue/card';
import Dialog from 'primevue/dialog';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Tag from 'primevue/tag';
import Badge from 'primevue/badge';
import Toast from 'primevue/toast';
import Menu from 'primevue/menu';
import Breadcrumb from 'primevue/breadcrumb';
import { ref } from 'vue';

// Sandbox State
const visible = ref(true);
const selectedRadio = ref('option1');
const selectedCheckbox = ref(false);
const textValue = ref('');
const selectedOption = ref(null);
const menuItems = ref([{ label: 'Menu Item 1', icon: 'pi pi-user' }, { label: 'Menu Item 2', icon: 'pi pi-search' }]);
const breadcrumbItems = ref([{ label: 'Home', command: () => {} }, { label: 'Dashboard' }]);
const products = ref([
    { id: '1', name: 'Laptop', price: 1200 },
    { id: '2', name: 'Phone', price: 800 },
    { id: '3', name: 'Tablet', price: 500 },
]);

const toast = ref();
const menu = ref();
</script>

<template>
  <div class="p-md bg-white text-black min-h-screen">
    <h1 class="text-xl font-bold mb-lg">Design Token Sandbox</h1>

    <!-- Existing Token Sections ... -->
    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">Colors</h2>
      <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-5 gap-md">
        <div v-for="color in [
          { name: 'primary-50', class: 'bg-primary-50' },
          { name: 'primary-100', class: 'bg-primary-100' },
          { name: 'primary-200', class: 'bg-primary-200' },
          { name: 'primary-300', class: 'bg-primary-300' },
          { name: 'primary-400', class: 'bg-primary-400' },
          { name: 'primary-500', class: 'bg-primary-500' },
          { name: 'primary-600', class: 'bg-primary-600' },
          { name: 'primary-700', class: 'bg-primary-700' },
          { name: 'primary-800', class: 'bg-primary-800' },
          { name: 'primary-900', class: 'bg-primary-900' },
          { name: 'secondary-500', class: 'bg-secondary-500' },
          { name: 'success-500', class: 'bg-success-500' },
          { name: 'warning-500', class: 'bg-warning-500' },
          { name: 'danger-500', class: 'bg-danger-500' },
          { name: 'info-500', class: 'bg-info-500' }
        ]" :key="color.name" class="flex flex-col items-center">
          <div :class="[color.class, 'w-full h-12 rounded-sm mb-xs shadow-sm']"></div>
          <span class="text-xs">{{ color.name }}</span>
        </div>
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">Spacing</h2>
      <div class="flex flex-wrap gap-md items-end">
        <div v-for="s in ['xs', 'sm', 'md', 'lg', 'xl']" :key="s" class="flex flex-col items-center">
          <div :class="[`p-${s}`, 'bg-primary-200 border border-primary-500 w-20 text-center text-xs']">
            p-{{ s }}
          </div>
          <span class="text-xs mt-xs">{{ s }}</span>
        </div>
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">Border Radius</h2>
      <div class="flex flex-wrap gap-md items-center">
        <div v-for="r in ['sm', 'md', 'lg', 'xl']" :key="r" class="flex flex-col items-center">
          <div :class="[`rounded-${r}`, 'bg-secondary-500 w-16 h-16 flex items-center justify-center text-white text-xs']">
            rounded-{{ r }}
          </div>
          <span class="text-xs mt-xs">{{ r }}</span>
        </div>
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">Shadows</h2>
      <div class="flex flex-wrap gap-xl items-center">
        <div v-for="s in ['soft', 'medium', 'hard']" :key="s" class="flex flex-col items-center">
          <div :class="[`shadow-${s}`, 'bg-white w-24 h-24 rounded-lg border border-gray-200 flex items-center justify-center text-xs']">
            shadow-{{ s }}
          </div>
          <span class="text-xs mt-xs">{{ s }}</span>
        </div>
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">Typography</h2>
      <div class="space-y-md">
        <p class="text-xs">Text xs (default small)</p>
        <p class="text-sm">Text sm</p>
        <p class="text-base">Text base</p>
        <p class="text-lg">Text lg</p>
        <p class="text-xl font-bold">Text xl bold</p>
      </div>
    </section>

    <!-- PrimeVue Components Section -->
    <hr class="my-xl border-slate-200" />
    <h1 class="text-2xl font-bold mb-xl">PrimeVue Component Sandbox</h1>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">T012: Buttons</h2>
      <div class="flex flex-wrap gap-md items-center">
        <Button label="Primary" />
        <Button label="Secondary" severity="secondary" />
        <Button label="Success" severity="success" />
        <Button label="Warning" severity="warning" />
        <Button label="Danger" severity="danger" />
        <Button label="Info" severity="info" />
        <Button label="Outlined" outlined />
        <Button label="Link" link />
        <Button label="Disabled" disabled />
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">T013: Inputs</h2>
      <div class="flex flex-wrap gap-md items-center">
        <InputText v-model="textValue" placeholder="Type here..." />
        <Select v-model="selectedOption" :options="['Option 1', 'Option 2', 'Option 3']" placeholder="Select an option" class="w-48" />
        <div class="flex items-center gap-xs">
          <Checkbox v-model="selectedCheckbox" :binary="true" />
          <label>Checkbox</label>
        </div>
        <div class="flex items-center gap-xs">
          <RadioButton v-model="selectedRadio" value="option1" />
          <label>Option 1</label>
        </div>
        <div class="flex items-center gap-xs">
          <RadioButton v-model="selectedRadio" value="option2" />
          <label>Option 2</label>
        </div>
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">T014: Layout & Containers</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-xl">
        <Card>
          <template #title>Card Title</template>
          <template #subtitle>Card Subtitle</template>
          <template #content>
            <p class="text-sm">This is the main content of the card.</p>
          </template>
          <template #footer>
            <Button label="Card Action" size="small" />
          </template>
        </Card>

        <Panel header="Panel Header">
          <p class="text-sm">This is the content of the panel.</p>
        </Panel>

        <Dialog v-model:visible="visible" header="Dialog Header" style="width: 25vw">
          <p class="text-sm">This is the content of the dialog.</p>
          <template #footer>
            <Button label="Close" @click="visible = false" outlined />
          </template>
        </Dialog>
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">T015: Data Display</h2>
      <div class="space-y-xl">
        <div class="flex gap-md items-center">
          <Tag value="New" severity="success" />
          <Tag value="Warning" severity="warning" />
          <Tag value="Error" severity="danger" />
          <Tag value="Info" severity="info" />
          <Badge value="2" severity="primary" />
        </div>

        <DataTable :value="products" class="w-full">
          <Column field="name" header="Name"></Column>
          <Column field="price" header="Price"></Column>
        </DataTable>
      </div>
    </section>

    <section class="mb-xl">
      <h2 class="text-lg font-semibold mb-md">T016: Navigation & Feedback</h2>
      <div class="space-y-xl">
        <div class="flex items-center gap-md">
          <Breadcrumb :model="breadcrumbItems" />
        </div>
        
        <div class="flex gap-md items-center">
          <Button label="Open Menu" @click="menu.toggle($event)" />
          <Menu ref="menu" :model="menuItems" />
        </div>

        <div class="flex gap-md items-center">
          <Button label="Show Toast" @click="() => toast.add({ severity: 'success', summary: 'Success', detail: 'Message Content', life: 3000 })" />
          <Toast ref="toast" />
        </div>
      </div>
    </section>
  </div>
</template>
