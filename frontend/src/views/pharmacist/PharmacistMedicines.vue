<template>
  <div>
    <!-- Header -->
    <div class="mb-8 flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">Medicines</h1>
        <p class="text-gray-400 text-sm mt-1">Manage your medicine inventory</p>
      </div>
      <button
        @click="showForm = true"
        class="text-white text-sm font-semibold px-5 py-2.5 rounded-xl transition flex items-center gap-2"
        style="background-color: #1e3a5f"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="w-4 h-4"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 4v16m8-8H4"
          />
        </svg>
        Add Medicine
      </button>
    </div>

    <!-- Stat Cards -->
    <div class="grid grid-cols-3 gap-5 mb-8">
      <div class="bg-white rounded-2xl p-5 border border-gray-100 shadow-sm">
        <div class="flex items-start justify-between mb-4">
          <div
            class="w-10 h-10 rounded-xl flex items-center justify-center"
            style="background-color: #e8f0fe"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5"
              style="color: #1e3a5f"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >Active</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Medicines</p>
        <p class="text-3xl font-bold text-gray-900">{{ medicines.length }}</p>
        <div
          class="mt-3 h-0.5 w-12 rounded-full"
          style="background-color: #1e3a5f"
        ></div>
      </div>

      <div class="bg-white rounded-2xl p-5 border border-gray-100 shadow-sm">
        <div class="flex items-start justify-between mb-4">
          <div
            class="w-10 h-10 rounded-xl flex items-center justify-center"
            style="background-color: #e8f0fe"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5"
              style="color: #1e3a5f"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M7 7h.01M7 3h5l5 5v5M7 3H5a2 2 0 00-2 2v14a2 2 0 002 2h14a2 2 0 002-2v-5"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >Unique</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Categories</p>
        <p class="text-3xl font-bold text-gray-900">{{ uniqueCategories }}</p>
        <div
          class="mt-3 h-0.5 w-12 rounded-full"
          style="background-color: #1e3a5f"
        ></div>
      </div>

      <div class="bg-white rounded-2xl p-5 border border-gray-100 shadow-sm">
        <div class="flex items-start justify-between mb-4">
          <div
            class="w-10 h-10 rounded-xl flex items-center justify-center bg-orange-50"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 text-orange-500"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full bg-orange-50 text-orange-500"
            >Monitor</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Reorder Items</p>
        <p class="text-3xl font-bold text-gray-900">{{ lowStock }}</p>
        <div class="mt-3 h-0.5 w-12 rounded-full bg-orange-400"></div>
      </div>
    </div>

    <!-- Search -->
    <div class="mb-4">
      <input
        v-model="search"
        type="text"
        placeholder="Search by name or category..."
        class="w-full px-4 py-2.5 bg-white border border-gray-200 rounded-xl text-sm text-gray-800 placeholder-gray-400 focus:outline-none transition"
      />
    </div>

    <!-- Table -->
    <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
      <div class="px-6 py-4 border-b border-gray-100">
        <h2 class="text-base font-bold text-gray-900">
          Medicine List
          <span class="text-gray-400 font-normal text-sm"
            >({{ filteredMedicines.length }})</span
          >
        </h2>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-gray-100">
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Name
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Generic Name
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Category
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Unit
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Price
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Reorder Level
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="med in filteredMedicines"
              :key="med.id"
              class="border-b border-gray-50 hover:bg-gray-50 transition"
            >
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                {{ med.name }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ med.generic_name || "—" }}
              </td>
              <td class="px-6 py-4">
                <span
                  class="px-2.5 py-1 rounded-full text-xs font-semibold"
                  style="background-color: #e8f0fe; color: #1e3a5f"
                >
                  {{ med.category || "General" }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ med.unit }}</td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                ₹{{ med.price.toFixed(2) }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ med.reorder_level }}
              </td>
            </tr>
            <tr v-if="filteredMedicines.length === 0">
              <td
                colspan="6"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                {{
                  search
                    ? "No medicines match your search"
                    : "No medicines added yet"
                }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Add Medicine Modal -->
    <div
      v-if="showForm"
      class="fixed inset-0 bg-black/40 flex items-center justify-center z-50"
      @click.self="showForm = false"
    >
      <div class="bg-white rounded-2xl shadow-xl w-full max-w-md p-6">
        <h2 class="text-xl font-bold text-gray-900 mb-1">Add New Medicine</h2>
        <p class="text-sm text-gray-400 mb-5">
          Fill in the details to add a medicine
        </p>

        <div class="space-y-4">
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Name *</label
            >
            <input
              v-model="form.name"
              type="text"
              placeholder="e.g. Paracetamol"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Generic Name</label
            >
            <input
              v-model="form.generic_name"
              type="text"
              placeholder="e.g. Acetaminophen"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Category</label
            >
            <input
              v-model="form.category"
              type="text"
              placeholder="e.g. Analgesic"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
          </div>
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label
                class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
                >Unit *</label
              >
              <select
                v-model="form.unit"
                class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
              >
                <option value="">Select unit</option>
                <option value="tablet">Tablet</option>
                <option value="capsule">Capsule</option>
                <option value="ml">ml</option>
                <option value="mg">mg</option>
                <option value="syrup">Syrup</option>
                <option value="injection">Injection</option>
                <option value="cream">Cream</option>
                <option value="drops">Drops</option>
              </select>
            </div>
            <div>
              <label
                class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
                >Price (₹) *</label
              >
              <input
                v-model="form.price"
                type="number"
                min="0"
                step="0.01"
                placeholder="0.00"
                class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
              />
            </div>
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Reorder Level</label
            >
            <input
              v-model="form.reorder_level"
              type="number"
              min="0"
              placeholder="e.g. 10"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
          </div>
        </div>

        <p v-if="formError" class="text-red-500 text-xs mt-3">
          {{ formError }}
        </p>

        <div class="flex gap-3 mt-6">
          <button
            @click="showForm = false"
            class="flex-1 px-4 py-2.5 border border-gray-200 rounded-xl text-sm font-medium text-gray-600 hover:bg-gray-50 transition"
          >
            Cancel
          </button>
          <button
            @click="addMedicine"
            :disabled="saving"
            class="flex-1 px-4 py-2.5 disabled:opacity-50 text-white rounded-xl text-sm font-medium transition"
            style="background-color: #1e3a5f"
          >
            {{ saving ? "Saving..." : "Add Medicine" }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import api from "../../api/axios";

const medicines = ref([]);
const search = ref("");
const showForm = ref(false);
const saving = ref(false);
const formError = ref("");

const form = ref({
  name: "",
  generic_name: "",
  category: "",
  unit: "",
  price: "",
  reorder_level: 10,
});

const filteredMedicines = computed(() => {
  if (!search.value) return medicines.value;
  const q = search.value.toLowerCase();
  return medicines.value.filter(
    (m) =>
      m.name.toLowerCase().includes(q) ||
      (m.category && m.category.toLowerCase().includes(q)) ||
      (m.generic_name && m.generic_name.toLowerCase().includes(q)),
  );
});

const uniqueCategories = computed(
  () => new Set(medicines.value.map((m) => m.category).filter(Boolean)).size,
);

const lowStock = computed(
  () => medicines.value.filter((m) => m.reorder_level > 0).length,
);

onMounted(loadMedicines);

async function loadMedicines() {
  try {
    const res = await api.get("/medicines");
    medicines.value = res.data.data || [];
  } catch (err) {
    console.error(err);
  }
}

async function addMedicine() {
  formError.value = "";
  if (!form.value.name) return (formError.value = "Name is required");
  if (!form.value.unit) return (formError.value = "Unit is required");
  if (!form.value.price) return (formError.value = "Price is required");

  saving.value = true;
  try {
    await api.post("/medicines", {
      name: form.value.name,
      generic_name: form.value.generic_name,
      category: form.value.category,
      unit: form.value.unit,
      price: parseFloat(form.value.price),
      reorder_level: parseInt(form.value.reorder_level) || 0,
    });
    showForm.value = false;
    form.value = {
      name: "",
      generic_name: "",
      category: "",
      unit: "",
      price: "",
      reorder_level: 10,
    };
    await loadMedicines();
  } catch (err) {
    formError.value = err.response?.data?.message || "Failed to add medicine";
  } finally {
    saving.value = false;
  }
}
</script>
