<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Hospital Command Center</h1>
      <p class="text-gray-400 text-sm mt-1">
        Morning, {{ auth.userName }}. Here's the clinical status for today.
      </p>
    </div>

    <!-- Stat Cards -->
    <div class="grid grid-cols-4 gap-5 mb-8">
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
                d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >+4 this month</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Doctors</p>
        <p class="text-3xl font-bold text-gray-900">{{ stats.totalDoctors }}</p>
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
                d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >Active Care</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Patients</p>
        <p class="text-3xl font-bold text-gray-900">
          {{ stats.totalPatients }}
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
                d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"
              />
            </svg>
          </div>
          <span
            class="text-xs font-semibold px-2 py-1 rounded-full"
            style="background-color: #e8f0fe; color: #1e3a5f"
            >All Shifts</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Receptionists</p>
        <p class="text-3xl font-bold text-gray-900">
          {{ stats.totalReceptionists }}
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
            style="background-color: #fff0f0"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-5 h-5 text-red-500"
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
            class="text-xs font-semibold px-2 py-1 rounded-full bg-red-50 text-red-500"
            >High Volume</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Appointments</p>
        <p class="text-3xl font-bold text-gray-900">
          {{ stats.totalAppointments }}
        </p>
        <div class="mt-3 h-0.5 w-12 rounded-full bg-red-400"></div>
      </div>
    </div>

    <!-- Two column layout -->
    <div class="grid grid-cols-2 gap-6">
      <!-- Upcoming Appointments -->
      <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
        <div
          class="px-6 py-4 border-b border-gray-100 flex items-center justify-between"
        >
          <h2 class="text-base font-bold text-gray-900">
            Upcoming Appointments
          </h2>
          <RouterLink
            to="/admin/users"
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
        <div class="overflow-x-auto">
          <table class="w-full">
            <thead>
              <tr class="border-b border-gray-100">
                <th
                  class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
                >
                  Patient
                </th>
                <th
                  class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
                >
                  Doctor
                </th>
                <th
                  class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
                >
                  Schedule
                </th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="appt in upcomingAppointments"
                :key="appt.id"
                class="border-b border-gray-50 hover:bg-gray-50 transition"
              >
                <td class="px-6 py-4">
                  <div class="flex items-center gap-3">
                    <div
                      class="w-8 h-8 rounded-full flex items-center justify-center text-white text-xs font-bold flex-shrink-0"
                      style="background-color: #1e3a5f"
                    >
                      {{ appt.patient_name?.charAt(0) }}
                    </div>
                    <div>
                      <p class="text-sm font-medium text-gray-900">
                        {{ appt.patient_name }}
                      </p>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 text-sm text-gray-500">
                  {{ appt.doctor_name }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-500">
                  {{ formatDate(appt.scheduled_at) }}
                </td>
              </tr>
              <tr v-if="upcomingAppointments.length === 0">
                <td
                  colspan="3"
                  class="px-6 py-8 text-center text-sm text-gray-400"
                >
                  No upcoming appointments
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Recent Completions -->
      <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
        <div
          class="px-6 py-4 border-b border-gray-100 flex items-center justify-between"
        >
          <h2 class="text-base font-bold text-gray-900">Recent Completions</h2>
          <span
            class="text-xs font-semibold text-gray-400 uppercase tracking-widest"
            >Today</span
          >
        </div>
        <div class="divide-y divide-gray-50">
          <div
            v-for="appt in recentCompletions"
            :key="appt.id"
            class="px-6 py-4 hover:bg-gray-50 transition"
          >
            <div class="flex items-start justify-between">
              <div>
                <p class="text-sm font-semibold text-gray-900">
                  {{ appt.patient_name }}
                </p>
                <p class="text-xs text-gray-400 mt-0.5">
                  {{ appt.type }} • {{ appt.doctor_name }}
                </p>
                <div class="flex items-center gap-1 mt-1">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="w-3 h-3 text-gray-400"
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
                  <p class="text-xs text-gray-400">
                    {{ formatDate(appt.scheduled_at) }}
                  </p>
                </div>
              </div>
              <span
                class="text-xs font-bold px-2.5 py-1 rounded-full"
                style="background-color: #e8f0fe; color: #1e3a5f"
                >COMPLETED</span
              >
            </div>
          </div>
          <div
            v-if="recentCompletions.length === 0"
            class="px-6 py-8 text-center text-sm text-gray-400"
          >
            No completions today
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
const users = ref([]);
const appointments = ref([]);

const stats = ref({
  totalDoctors: 0,
  totalPatients: 0,
  totalReceptionists: 0,
  totalAppointments: 0,
});

const upcomingAppointments = computed(() =>
  appointments.value
    .filter((a) => a.status === "pending" || a.status === "approved")
    .slice(0, 5),
);

const recentCompletions = computed(() =>
  appointments.value.filter((a) => a.status === "completed").slice(0, 5),
);

onMounted(async () => {
  try {
    const [usersRes, apptRes] = await Promise.all([
      api.get("/users"),
      api.get("/appointments"),
    ]);
    users.value = usersRes.data.data;
    appointments.value = apptRes.data.data || [];
    stats.value.totalDoctors = users.value.filter(
      (u) => u.role === "doctor",
    ).length;
    stats.value.totalPatients = users.value.filter(
      (u) => u.role === "patient",
    ).length;
    stats.value.totalReceptionists = users.value.filter(
      (u) => u.role === "receptionist",
    ).length;
    stats.value.totalAppointments = appointments.value.length;
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
