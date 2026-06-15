<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">My Appointments</h1>
      <p class="text-gray-400 text-sm mt-1">Manage your patient appointments</p>
    </div>

    <!-- Table -->
    <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
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
                Date
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Type
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Reason
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
              v-for="appt in appointments"
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
                  <p class="text-sm font-medium text-gray-900">
                    {{ appt.patient_name }}
                  </p>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ formatDate(appt.scheduled_at) }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500 capitalize">
                {{ appt.type }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ appt.reason || "—" }}
              </td>
              <td class="px-6 py-4">
                <span
                  class="px-2.5 py-1 rounded-full text-xs font-semibold capitalize"
                  :class="{
                    'bg-yellow-50 text-yellow-600': appt.status === 'pending',
                    'bg-blue-50 text-blue-600': appt.status === 'approved',
                    'bg-green-50 text-green-600': appt.status === 'completed',
                    'bg-red-50 text-red-600':
                      appt.status === 'rejected' || appt.status === 'cancelled',
                  }"
                >
                  {{ appt.status }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex gap-2">
                  <button
                    v-if="appt.status === 'pending'"
                    @click="updateStatus(appt.id, 'approve')"
                    class="text-xs text-white px-3 py-1.5 rounded-lg font-medium transition"
                    style="background-color: #1e3a5f"
                  >
                    Approve
                  </button>
                  <button
                    v-if="appt.status === 'pending'"
                    @click="updateStatus(appt.id, 'reject')"
                    class="text-xs bg-red-50 text-red-500 hover:bg-red-100 px-3 py-1.5 rounded-lg font-medium transition"
                  >
                    Reject
                  </button>
                  <button
                    v-if="appt.status === 'approved'"
                    @click="updateStatus(appt.id, 'complete')"
                    class="text-xs bg-green-50 text-green-600 hover:bg-green-100 px-3 py-1.5 rounded-lg font-medium transition"
                  >
                    Complete
                  </button>
                  <button
                    v-if="appt.status === 'completed'"
                    @click="goToPrescription(appt)"
                    class="text-xs bg-blue-50 text-blue-600 hover:bg-blue-100 px-3 py-1.5 rounded-lg font-medium transition"
                  >
                    Prescribe
                  </button>
                  <span
                    v-if="
                      appt.status === 'rejected' || appt.status === 'cancelled'
                    "
                    class="text-xs text-gray-400"
                    >—</span
                  >
                </div>
              </td>
            </tr>
            <tr v-if="appointments.length === 0">
              <td
                colspan="6"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                No appointments yet
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../../stores/auth";
import api from "../../api/axios";

const auth = useAuthStore();
const router = useRouter();
const appointments = ref([]);

onMounted(async () => {
  await loadAppointments();
});

async function loadAppointments() {
  try {
    const res = await api.get("/appointments");
    appointments.value = (res.data.data || []).filter(
      (a) => a.doctor_name === auth.userName,
    );
  } catch (err) {
    console.error(err);
  }
}

async function updateStatus(id, action) {
  try {
    await api.patch(`/appointments/${id}/${action}`);
    await loadAppointments();
  } catch (err) {
    console.error(err);
  }
}

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}
function goToPrescription(appt) {
  router.push({
    name: "WritePrescription",
    query: {
      appointment_id: appt.id,
      patient_id: appt.patient_id,
      patient_name: appt.patient_name,
    },
  });
}
</script>
