<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Medicines</h1>
      <p class="text-gray-400 text-sm mt-1">
        Reference list of available medicines
      </p>
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
          Available Medicines
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
              <td class="px-6 py-4 text-sm text-gray-500 capitalize">
                {{ med.unit }}
              </td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                ₹{{ med.price.toFixed(2) }}
              </td>
            </tr>
            <tr v-if="filteredMedicines.length === 0">
              <td
                colspan="5"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                {{
                  search
                    ? "No medicines match your search"
                    : "No medicines available"
                }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import api from "../../api/axios";

const medicines = ref([]);
const search = ref("");

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

onMounted(async () => {
  try {
    const res = await api.get("/medicines");
    medicines.value = res.data.data || [];
  } catch (err) {
    console.error(err);
  }
});
</script>
