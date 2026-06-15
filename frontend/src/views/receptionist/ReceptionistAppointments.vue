<template>
  <div>
    <div class="mb-8 flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">Appointments</h1>
        <p class="text-gray-400 text-sm mt-1">
          View and manage all appointments
        </p>
      </div>
    </div>

    <!-- Filter Tabs -->
    <div class="flex gap-2 mb-6 flex-wrap">
      <button
        v-for="f in filters"
        :key="f.value"
        @click="activeFilter = f.value"
        class="px-4 py-2 rounded-xl text-sm font-medium transition"
        :class="
          activeFilter === f.value
            ? 'text-white'
            : 'bg-white border border-gray-200 text-gray-500 hover:bg-gray-50'
        "
        :style="activeFilter === f.value ? 'background-color: #1e3a5f;' : ''"
      >
        {{ f.label }}
        <span class="ml-1 text-xs opacity-70">({{ getCount(f.value) }})</span>
      </button>
    </div>

    <!-- Appointments Table -->
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
                Type
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
              v-for="appt in filteredAppointments"
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
                {{ appt.doctor_name }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ formatDate(appt.scheduled_at) }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500 capitalize">
                {{ appt.type }}
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
                  >{{ appt.status }}</span
                >
              </td>
              <td class="px-6 py-4">
                <button
                  v-if="appt.status === 'pending'"
                  @click="cancelAppointment(appt.id)"
                  class="text-xs text-red-500 hover:text-red-600 font-medium hover:underline"
                >
                  Cancel
                </button>
                <span v-else class="text-xs text-gray-400">—</span>
              </td>
            </tr>
            <tr v-if="filteredAppointments.length === 0">
              <td
                colspan="6"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                No appointments found
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

const appointments = ref([]);
const activeFilter = ref("all");

const filters = [
  { label: "All", value: "all" },
  { label: "Pending", value: "pending" },
  { label: "Approved", value: "approved" },
  { label: "Completed", value: "completed" },
  { label: "Cancelled", value: "cancelled" },
];

const filteredAppointments = computed(() => {
  if (activeFilter.value === "all") return appointments.value;
  return appointments.value.filter((a) => a.status === activeFilter.value);
});

function getCount(filter) {
  if (filter === "all") return appointments.value.length;
  return appointments.value.filter((a) => a.status === filter).length;
}

onMounted(async () => {
  try {
    const res = await api.get("/appointments");
    appointments.value = res.data?.data || [];
  } catch (err) {
    console.error(err);
    appointments.value = [];
  }
});

async function cancelAppointment(id) {
  if (!confirm("Cancel this appointment?")) return;
  try {
    await api.patch(`/appointments/${id}/cancel`);
    appointments.value = appointments.value.map((a) =>
      a.id === id ? { ...a, status: "cancelled" } : a,
    );
  } catch (err) {
    console.error(err);
  }
}

function formatDate(dateStr) {
  if (!dateStr) return "—";
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}
</script>
