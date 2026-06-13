<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Patient Dashboard</h1>
      <p class="text-gray-400 text-sm mt-1">
        Welcome back, {{ auth.userName }}. Here's your health summary.
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
                d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >All Time</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Appointments</p>
        <p class="text-3xl font-bold text-gray-900">
          {{ appointments.length }}
        </p>
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
                d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >Active</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Upcoming</p>
        <p class="text-3xl font-bold text-gray-900">{{ upcoming }}</p>
        <div
          class="mt-3 h-0.5 w-12 rounded-full"
          style="background-color: #1e3a5f"
        ></div>
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
        <p class="text-sm text-gray-400 mb-1">Completed</p>
        <p class="text-3xl font-bold text-gray-900">{{ completed }}</p>
        <div class="mt-3 h-0.5 w-12 rounded-full bg-green-400"></div>
      </div>
    </div>

    <!-- Two column layout -->
    <div class="grid grid-cols-2 gap-6">
      <!-- Appointments -->
      <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
        <div
          class="px-6 py-4 border-b border-gray-100 flex items-center justify-between"
        >
          <h2 class="text-base font-bold text-gray-900">My Appointments</h2>
          <RouterLink
            to="/patient/book"
            class="text-white text-xs font-semibold px-4 py-2 rounded-xl transition"
            style="background-color: #1e3a5f"
          >
            + Book
          </RouterLink>
        </div>
        <div class="overflow-x-auto">
          <table class="w-full">
            <thead>
              <tr class="border-b border-gray-100">
                <th
                  class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
                >
                  Doctor
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
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="appt in appointments.slice(0, 5)"
                :key="appt.id"
                class="border-b border-gray-50 hover:bg-gray-50 transition"
              >
                <td class="px-6 py-4">
                  <div class="flex items-center gap-3">
                    <div
                      class="w-8 h-8 rounded-full flex items-center justify-center text-white text-xs font-bold flex-shrink-0"
                      style="background-color: #1e3a5f"
                    >
                      {{ appt.doctor_name?.charAt(0) }}
                    </div>
                    <p class="text-sm font-medium text-gray-900">
                      {{ appt.doctor_name }}
                    </p>
                  </div>
                </td>
                <td class="px-6 py-4 text-sm text-gray-500">
                  {{ formatDate(appt.scheduled_at) }}
                </td>
                <td class="px-6 py-4">
                  <span
                    class="px-2.5 py-1 rounded-full text-xs font-semibold capitalize"
                    :class="{
                      'bg-yellow-50 text-yellow-600': appt.status === 'pending',
                      'bg-blue-50 text-blue-600': appt.status === 'approved',
                      'bg-green-50 text-green-600': appt.status === 'completed',
                      'bg-red-50 text-red-600':
                        appt.status === 'rejected' ||
                        appt.status === 'cancelled',
                    }"
                  >
                    {{ appt.status }}
                  </span>
                </td>
              </tr>
              <tr v-if="appointments.length === 0">
                <td
                  colspan="3"
                  class="px-6 py-8 text-center text-sm text-gray-400"
                >
                  No appointments yet
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Prescriptions -->
      <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
        <div
          class="px-6 py-4 border-b border-gray-100 flex items-center justify-between"
        >
          <h2 class="text-base font-bold text-gray-900">My Prescriptions</h2>
          <RouterLink
            to="/patient/prescriptions"
            class="text-sm font-medium flex items-center gap-1"
            style="color: #1e3a5f"
          >
            View All
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
                d="M17 8l4 4m0 0l-4 4m4-4H3"
              />
            </svg>
          </RouterLink>
        </div>
        <div class="divide-y divide-gray-50">
          <div
            v-for="prescription in prescriptions.slice(0, 5)"
            :key="prescription.id"
            class="px-6 py-4 hover:bg-gray-50 transition"
          >
            <div class="flex items-start justify-between">
              <div>
                <p class="text-sm font-semibold text-gray-900">
                  {{ prescription.diagnosis }}
                </p>
                <p class="text-xs text-gray-400 mt-0.5">
                  {{ formatDate(prescription.created_at) }}
                </p>
              </div>
              <span
                class="text-xs font-semibold px-2.5 py-1 rounded-full"
                :class="
                  prescription.dispensed
                    ? 'bg-green-50 text-green-600'
                    : 'bg-orange-50 text-orange-500'
                "
              >
                {{ prescription.dispensed ? "Dispensed" : "Pending" }}
              </span>
            </div>
          </div>
          <div
            v-if="prescriptions.length === 0"
            class="px-6 py-8 text-center text-sm text-gray-400"
          >
            No prescriptions yet
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { RouterLink } from "vue-router";
import { useAuthStore } from "../../stores/auth";
import api from "../../api/axios";

const auth = useAuthStore();
const appointments = ref([]);
const prescriptions = ref([]);

const upcoming = computed(
  () =>
    appointments.value.filter(
      (a) => a.status === "approved" || a.status === "pending",
    ).length,
);
const completed = computed(
  () => appointments.value.filter((a) => a.status === "completed").length,
);

onMounted(async () => {
  try {
    const patientsRes = await api.get("/patients");
    const myPatient = patientsRes.data.data.find(
      (p) => p.email === auth.user.email,
    );
    if (myPatient) {
      const [apptRes, prescRes] = await Promise.all([
        api.get("/appointments"),
        api.get(`/patients/${myPatient.id}/prescriptions`),
      ]);
      appointments.value = apptRes.data.data || [];
      prescriptions.value = prescRes.data.data || [];
    }
  } catch (err) {
    console.error(err);
  }
});

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}
</script>
