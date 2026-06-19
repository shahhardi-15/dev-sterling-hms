<template>
  <div>
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-gray-900">Book Appointment</h1>
      <p class="text-gray-500 text-sm mt-1">
        Schedule a new appointment with a doctor
      </p>
    </div>

    <div
      class="bg-white rounded-2xl border border-gray-100 shadow-sm p-4 sm:p-8 max-w-2xl"
    >
      <div
        v-if="success"
        class="bg-green-50 border border-green-200 text-green-600 text-sm px-4 py-3 rounded-lg mb-6"
      >
        Appointment booked successfully!
      </div>

      <div
        v-if="error"
        class="bg-red-50 border border-red-200 text-red-600 text-sm px-4 py-3 rounded-lg mb-6"
      >
        {{ error }}
      </div>

      <div class="mb-5">
        <label
          class="block text-xs font-semibold text-gray-600 uppercase tracking-widest mb-2"
          >Select Doctor</label
        >
        <select
          v-model="form.doctor_id"
          class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:border-blue-500 transition"
        >
          <option value="">Choose a doctor</option>
          <option v-for="doctor in doctors" :key="doctor.id" :value="doctor.id">
            {{ doctor.full_name }} — {{ doctor.specialization || "General" }}
          </option>
        </select>
      </div>

      <div class="mb-5">
        <label
          class="block text-xs font-semibold text-gray-600 uppercase tracking-widest mb-2"
          >Date & Time</label
        >
        <input
          v-model="form.scheduled_at"
          type="datetime-local"
          :min="minDateTime"
          class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:border-blue-500 transition"
        />
        <p v-if="!isEmergency" class="text-xs text-gray-400 mt-1">
          Working hours: Monday – Saturday, 9:00 AM to 6:00 PM
        </p>
        <p v-else class="text-xs text-orange-500 mt-1">
          Emergency appointments are available 24/7
        </p>
      </div>

      <div class="mb-5">
        <label
          class="block text-xs font-semibold text-gray-600 uppercase tracking-widest mb-2"
          >Appointment Type</label
        >
        <select
          v-model="form.type"
          class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:border-blue-500 transition"
        >
          <option value="">Select type</option>
          <option value="online">Online</option>
          <option value="walk_in">Walk In</option>
          <option value="follow_up">Follow Up</option>
          <option value="emergency">Emergency (24/7)</option>
        </select>
      </div>

      <div class="mb-8">
        <label
          class="block text-xs font-semibold text-gray-600 uppercase tracking-widest mb-2"
          >Reason</label
        >
        <textarea
          v-model="form.reason"
          rows="3"
          placeholder="Describe your symptoms or reason for visit"
          class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:border-blue-500 transition resize-none"
        ></textarea>
      </div>

      <button
        @click="bookAppointment"
        :disabled="booking"
        class="w-full text-white font-semibold py-3 rounded-xl flex items-center justify-center gap-2 transition disabled:opacity-60"
        style="background-color: #1e3a5f"
      >
        {{ booking ? "Booking..." : "Book Appointment" }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useAuthStore } from "../../stores/auth";
import api from "../../api/axios";

const auth = useAuthStore();
const doctors = ref([]);
const booking = ref(false);
const success = ref(false);
const error = ref("");

const form = ref({
  doctor_id: "",
  scheduled_at: "",
  type: "",
  reason: "",
});

const isEmergency = computed(() => form.value.type === "emergency");

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
  if (isEmergency.value) return null;
  const day = date.getDay();
  if (day === 0) return "We are closed on Sundays. Please choose another day.";
  const hours = date.getHours();
  const minutes = date.getMinutes();
  const timeInMinutes = hours * 60 + minutes;
  if (timeInMinutes < 9 * 60 || timeInMinutes > 18 * 60)
    return "Please book between 9:00 AM and 6:00 PM (Monday to Saturday).";
  return null;
}

onMounted(async () => {
  try {
    const res = await api.get("/doctors");
    doctors.value = res.data.data || [];
  } catch (err) {
    console.error(err);
  }
});

async function bookAppointment() {
  if (!form.value.doctor_id || !form.value.scheduled_at || !form.value.type) {
    error.value = "Please fill in all required fields.";
    return;
  }
  const timeError = validateDateTime(form.value.scheduled_at);
  if (timeError) {
    error.value = timeError;
    return;
  }
  booking.value = true;
  error.value = "";
  success.value = false;
  try {
    const patientsRes = await api.get("/patients");
    const myPatient = patientsRes.data.data.find(
      (p) => p.email === auth.user.email,
    );
    if (!myPatient) {
      error.value = "Patient profile not found.";
      return;
    }
    await api.post("/appointments", {
      patient_id: myPatient.id,
      doctor_id: form.value.doctor_id,
      scheduled_at: new Date(form.value.scheduled_at).toISOString(),
      type: form.value.type,
      reason: form.value.reason,
    });
    success.value = true;
    form.value = { doctor_id: "", scheduled_at: "", type: "", reason: "" };
  } catch (err) {
    error.value = err.response?.data?.error || "Failed to book appointment.";
  } finally {
    booking.value = false;
  }
}
</script>
