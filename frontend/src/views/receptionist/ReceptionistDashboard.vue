<template>
  <div>
    <!-- Header -->
    <div class="mb-8 flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">Receptionist Dashboard</h1>
        <p class="text-gray-400 text-sm mt-1">
          Welcome back, {{ auth.userName }}. Manage today's appointments.
        </p>
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
        Book Appointment
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
        <p class="text-sm text-gray-400 mb-1">Pending</p>
        <p class="text-3xl font-bold text-gray-900">{{ pending }}</p>
        <div class="mt-3 h-0.5 w-12 rounded-full bg-orange-400"></div>
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
            >Registered</span
          >
        </div>
        <p class="text-sm text-gray-400 mb-1">Total Patients</p>
        <p class="text-3xl font-bold text-gray-900">{{ patients.length }}</p>
        <div
          class="mt-3 h-0.5 w-12 rounded-full"
          style="background-color: #1e3a5f"
        ></div>
      </div>
    </div>

    <!-- Appointments Table -->
    <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
      <div class="px-6 py-4 border-b border-gray-100">
        <h2 class="text-base font-bold text-gray-900">Appointment Queue</h2>
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
                >
                  {{ appt.status }}
                </span>
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
            <tr v-if="appointments.length === 0">
              <td
                colspan="6"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                No appointments
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Book Appointment Modal -->
    <div
      v-if="showForm"
      class="fixed inset-0 bg-black/40 flex items-center justify-center z-50"
    >
      <div
        class="bg-white rounded-2xl shadow-xl w-full max-w-md p-6 overflow-y-auto max-h-[90vh]"
      >
        <h2 class="text-xl font-bold text-gray-900 mb-1">Book Appointment</h2>
        <p class="text-sm text-gray-400 mb-5">
          Book an appointment on behalf of a patient
        </p>

        <div class="space-y-4">
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Patient *</label
            >
            <select
              v-model="form.patient_id"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            >
              <option value="">Select patient</option>
              <option v-for="p in patients" :key="p.id" :value="p.id">
                {{ p.full_name }}
              </option>
            </select>
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Doctor *</label
            >
            <select
              v-model="form.doctor_id"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            >
              <option value="">Select doctor</option>
              <option v-for="d in doctors" :key="d.id" :value="d.id">
                {{ d.full_name }} — {{ d.specialization || "General" }}
              </option>
            </select>
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Type *</label
            >
            <select
              v-model="form.type"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            >
              <option value="">Select type</option>
              <option value="online">Online</option>
              <option value="walk_in">Walk In</option>
              <option value="follow_up">Follow Up</option>
              <option value="emergency">Emergency (24/7)</option>
            </select>
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Date & Time *</label
            >
            <input
              v-model="form.scheduled_at"
              type="datetime-local"
              :min="minDateTime"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
            <p
              v-if="form.type !== 'emergency'"
              class="text-xs text-gray-400 mt-1"
            >
              Working hours: Monday – Saturday, 9:00 AM to 6:00 PM
            </p>
            <p v-else class="text-xs text-orange-500 mt-1">
              Emergency appointments are available 24/7
            </p>
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Reason</label
            >
            <textarea
              v-model="form.reason"
              rows="2"
              placeholder="Reason for visit"
              class="w-full px-4 py-2.5 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition resize-none"
            ></textarea>
          </div>
        </div>

        <p v-if="formError" class="text-red-500 text-xs mt-3">
          {{ formError }}
        </p>

        <div class="flex gap-3 mt-6">
          <button
            @click="closeForm"
            class="flex-1 px-4 py-2.5 border border-gray-200 rounded-xl text-sm font-medium text-gray-600 hover:bg-gray-50 transition"
          >
            Cancel
          </button>
          <button
            @click="bookAppointment"
            :disabled="booking"
            class="flex-1 px-4 py-2.5 disabled:opacity-50 text-white rounded-xl text-sm font-medium transition"
            style="background-color: #1e3a5f"
          >
            {{ booking ? "Booking..." : "Book Appointment" }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useAuthStore } from "../../stores/auth";
import api from "../../api/axios";

const auth = useAuthStore();
const appointments = ref([]);
const patients = ref([]);
const doctors = ref([]);
const showForm = ref(false);
const booking = ref(false);
const formError = ref("");

const form = ref({
  patient_id: "",
  doctor_id: "",
  type: "",
  scheduled_at: "",
  reason: "",
});

const pending = computed(
  () => appointments.value.filter((a) => a.status === "pending").length,
);

const minDateTime = computed(() => {
  const now = new Date();
  now.setMinutes(now.getMinutes() + 30);
  return now.toISOString().slice(0, 16);
});

function validateDateTime(dateTimeStr) {
  if (!dateTimeStr) return "Please select a date and time.";
  const date = new Date(dateTimeStr);
  const now = new Date();
  if (date < now) return "Cannot book an appointment in the past.";
  if (form.value.type === "emergency") return null;
  const day = date.getDay();
  if (day === 0) return "We are closed on Sundays. Please choose another day.";
  const timeInMinutes = date.getHours() * 60 + date.getMinutes();
  if (timeInMinutes < 9 * 60 || timeInMinutes > 18 * 60)
    return "Please book between 9:00 AM and 6:00 PM (Monday to Saturday).";
  return null;
}

onMounted(async () => {
  try {
    const [apptRes, patientRes, doctorRes] = await Promise.all([
      api.get("/appointments"),
      api.get("/patients"),
      api.get("/doctors"),
    ]);
    appointments.value = apptRes.data.data;
    patients.value = patientRes.data.data;
    doctors.value = doctorRes.data.data;
  } catch (err) {
    console.error(err);
  }
});

function closeForm() {
  showForm.value = false;
  formError.value = "";
  form.value = {
    patient_id: "",
    doctor_id: "",
    type: "",
    scheduled_at: "",
    reason: "",
  };
}

async function bookAppointment() {
  formError.value = "";
  if (!form.value.patient_id)
    return (formError.value = "Please select a patient.");
  if (!form.value.doctor_id)
    return (formError.value = "Please select a doctor.");
  if (!form.value.type)
    return (formError.value = "Please select appointment type.");
  const timeError = validateDateTime(form.value.scheduled_at);
  if (timeError) return (formError.value = timeError);

  booking.value = true;
  try {
    await api.post("/appointments", {
      patient_id: form.value.patient_id,
      doctor_id: form.value.doctor_id,
      scheduled_at: new Date(form.value.scheduled_at).toISOString(),
      type: form.value.type,
      reason: form.value.reason,
    });
    const apptRes = await api.get("/appointments");
    appointments.value = apptRes.data.data;
    closeForm();
  } catch (err) {
    formError.value =
      err.response?.data?.error || "Failed to book appointment.";
  } finally {
    booking.value = false;
  }
}

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
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}
</script>
