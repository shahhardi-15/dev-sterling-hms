<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Pharmacist Dashboard</h1>
      <p class="text-gray-400 text-sm mt-1">
        Welcome back, {{ auth.userName }}. Here's your dispensing queue.
      </p>
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
                d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >All Time</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Prescriptions</p>
        <p class="text-3xl font-bold text-gray-900">
          {{ prescriptions.length }}
        </p>
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
                d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full bg-orange-50 text-orange-500"
            >Action Needed</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Pending Dispense</p>
        <p class="text-3xl font-bold text-gray-900">{{ pending }}</p>
        <div class="mt-3 h-0.5 w-12 rounded-full bg-orange-400"></div>
      </div>

      <div class="bg-white rounded-2xl p-5 border border-gray-100 shadow-sm">
        <div class="flex items-start justify-between mb-4">
          <div
            class="w-10 h-10 rounded-xl flex items-center justify-center bg-green-50"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 text-green-600"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full bg-green-50 text-green-600"
            >Done</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Dispensed</p>
        <p class="text-3xl font-bold text-gray-900">{{ dispensed }}</p>
        <div class="mt-3 h-0.5 w-12 rounded-full bg-green-400"></div>
      </div>
    </div>

    <!-- Prescriptions Table -->
    <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
      <div class="px-6 py-4 border-b border-gray-100">
        <h2 class="text-base font-bold text-gray-900">Prescription Queue</h2>
      </div>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-gray-100">
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Prescription ID
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Diagnosis
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Date
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Status
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="prescription in prescriptions"
              :key="prescription.id"
              class="border-b border-gray-50 hover:bg-gray-50 transition"
            >
              <td class="px-6 py-4 text-sm font-mono text-gray-400">
                {{ prescription.id.slice(0, 8) }}...
              </td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                {{ prescription.diagnosis }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ formatDate(prescription.created_at) }}
              </td>
              <td class="px-6 py-4">
                <span
                  class="px-2.5 py-1 rounded-full text-xs font-semibold"
                  :class="
                    prescription.dispensed
                      ? 'bg-green-50 text-green-600'
                      : 'bg-orange-50 text-orange-500'
                  "
                >
                  {{ prescription.dispensed ? "Dispensed" : "Pending" }}
                </span>
              </td>
              <td class="px-6 py-4">
                <button
                  v-if="!prescription.dispensed"
                  @click="dispense(prescription.id)"
                  class="text-xs text-white px-3 py-1.5 rounded-lg font-medium transition"
                  style="background-color: #1e3a5f"
                >
                  Dispense
                </button>
                <span v-else class="text-xs text-gray-400">—</span>
              </td>
            </tr>
            <tr v-if="prescriptions.length === 0">
              <td
                colspan="5"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                No prescriptions yet
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
import { useAuthStore } from "../../stores/auth";
import api from "../../api/axios";

const auth = useAuthStore();
const prescriptions = ref([]);
const loading = ref(false);

const pending = computed(
  () => prescriptions.value.filter((p) => !p.dispensed).length,
);
const dispensed = computed(
  () => prescriptions.value.filter((p) => p.dispensed).length,
);

onMounted(async () => {
  await loadPrescriptions();
});

async function loadPrescriptions() {
  if (loading.value) return;
  loading.value = true;
  try {
    const patientsRes = await api.get("/patients");
    const allPrescriptions = [];
    for (const patient of patientsRes.data.data) {
      const res = await api.get(`/patients/${patient.id}/prescriptions`);
      if (res.data.data) allPrescriptions.push(...res.data.data);
    }
    prescriptions.value = allPrescriptions;
  } catch (err) {
    console.error(err);
  } finally {
    loading.value = false;
  }
}

async function dispense(id) {
  try {
    await api.patch(`/prescriptions/${id}/dispense`);
    const prescription = prescriptions.value.find((p) => p.id === id);
    if (prescription) prescription.dispensed = true;
  } catch (err) {
    console.error(err);
  }
}

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}
</script>
