<template>
  <div>
    <div class="mb-8">
      <h1 class="text-2xl font-bold text-gray-900">Write Prescription</h1>
      <p class="text-gray-500 text-sm mt-1">
        Create a prescription for a patient
      </p>
    </div>

    <div
      class="bg-white rounded-2xl border border-gray-100 shadow-sm p-8 max-w-3xl"
    >
      <div
        v-if="success"
        class="bg-green-50 border border-green-200 text-green-600 text-sm px-4 py-3 rounded-lg mb-6"
      >
        Prescription created successfully!
      </div>

      <div
        v-if="error"
        class="bg-red-50 border border-red-200 text-red-600 text-sm px-4 py-3 rounded-lg mb-6"
      >
        {{ error }}
      </div>

      <!-- Appointment -->
      <div class="mb-5">
        <label
          class="block text-xs font-semibold text-gray-600 uppercase tracking-widest mb-2"
          >Select Appointment</label
        >
        <select
          v-model="form.appointment_id"
          @change="onAppointmentChange"
          class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:border-blue-500 transition"
        >
          <option value="">Choose an appointment</option>
          <option v-for="appt in appointments" :key="appt.id" :value="appt.id">
            {{ appt.patient_name }} — {{ formatDate(appt.scheduled_at) }}
          </option>
        </select>
      </div>

      <!-- Diagnosis -->
      <div class="mb-5">
        <label
          class="block text-xs font-semibold text-gray-600 uppercase tracking-widest mb-2"
          >Diagnosis</label
        >
        <input
          v-model="form.diagnosis"
          type="text"
          placeholder="e.g. Tension headache"
          class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:border-blue-500 transition"
        />
      </div>

      <!-- Notes -->
      <div class="mb-6">
        <label
          class="block text-xs font-semibold text-gray-600 uppercase tracking-widest mb-2"
          >Notes</label
        >
        <textarea
          v-model="form.notes"
          rows="2"
          placeholder="Additional notes for the patient"
          class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none focus:border-blue-500 transition resize-none"
        ></textarea>
      </div>

      <!-- Medicine items -->
      <div class="mb-6">
        <div class="flex items-center justify-between mb-3">
          <label
            class="text-xs font-semibold text-gray-600 uppercase tracking-widest"
            >Medicines</label
          >
          <button
            @click="addItem"
            class="text-xs font-semibold"
            style="color: #1e3a5f"
          >
            + Add Medicine
          </button>
        </div>

        <div
          v-for="(item, index) in form.items"
          :key="index"
          class="bg-gray-50 rounded-xl p-4 mb-3 border border-gray-100"
        >
          <div class="grid grid-cols-2 gap-3 mb-3">
            <div>
              <label class="block text-xs text-gray-500 mb-1">Medicine</label>
              <select
                v-model="item.medicine_id"
                class="w-full px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm focus:outline-none focus:border-blue-500"
              >
                <option value="">Select medicine</option>
                <option v-for="med in medicines" :key="med.id" :value="med.id">
                  {{ med.name }}
                </option>
              </select>
            </div>
            <div>
              <label class="block text-xs text-gray-500 mb-1">Dosage</label>
              <input
                v-model="item.dosage"
                type="text"
                placeholder="e.g. 500mg"
                class="w-full px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm focus:outline-none focus:border-blue-500"
              />
            </div>
          </div>
          <div class="grid grid-cols-3 gap-3">
            <div>
              <label class="block text-xs text-gray-500 mb-1">Frequency</label>
              <input
                v-model="item.frequency"
                type="text"
                placeholder="e.g. Twice daily"
                class="w-full px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm focus:outline-none focus:border-blue-500"
              />
            </div>
            <div>
              <label class="block text-xs text-gray-500 mb-1"
                >Duration (days)</label
              >
              <input
                v-model="item.duration_days"
                type="number"
                placeholder="5"
                class="w-full px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm focus:outline-none focus:border-blue-500"
              />
            </div>
            <div>
              <label class="block text-xs text-gray-500 mb-1">Quantity</label>
              <input
                v-model="item.quantity"
                type="number"
                placeholder="10"
                class="w-full px-3 py-2 bg-white border border-gray-200 rounded-lg text-sm focus:outline-none focus:border-blue-500"
              />
            </div>
          </div>
          <button
            v-if="form.items.length > 1"
            @click="removeItem(index)"
            class="mt-2 text-xs text-red-500 hover:text-red-600"
          >
            Remove
          </button>
        </div>
      </div>

      <button
        @click="submitPrescription"
        :disabled="submitting"
        class="w-full text-white font-semibold py-3 rounded-xl transition disabled:opacity-60"
        style="background-color: #1e3a5f"
      >
        {{ submitting ? "Creating..." : "Create Prescription" }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRoute } from "vue-router";
import api from "../../api/axios";

const route = useRoute();
const appointments = ref([]);
const medicines = ref([]);
const submitting = ref(false);
const success = ref(false);
const error = ref("");
const selectedPatientID = ref("");
const selectedDoctorID = ref("");

const form = ref({
  appointment_id: "",
  patient_id: "",
  doctor_id: "",
  diagnosis: "",
  notes: "",
  items: [
    {
      medicine_id: "",
      dosage: "",
      frequency: "",
      duration_days: 5,
      quantity: 1,
      instructions: "",
    },
  ],
});

onMounted(async () => {
  try {
    const [apptRes, medRes] = await Promise.all([
      api.get("/appointments"),
      api.get("/medicines"),
    ]);
    appointments.value = apptRes.data.data.filter(
      (a) => a.status === "approved" || a.status === "completed",
    );
    medicines.value = medRes.data.data;

    // Auto-select if coming from MyAppointments
    if (route.query.appointment_id) {
      form.value.appointment_id = route.query.appointment_id;
      form.value.patient_id = route.query.patient_id;
      onAppointmentChange();
    }
  } catch (err) {
    console.error(err);
  }
});

function onAppointmentChange() {
  const appt = appointments.value.find(
    (a) => a.id === form.value.appointment_id,
  );
  if (appt) {
    form.value.patient_id = appt.patient_id;
    form.value.doctor_id = appt.doctor_id;
  }
}

function addItem() {
  form.value.items.push({
    medicine_id: "",
    dosage: "",
    frequency: "",
    duration_days: 5,
    quantity: 1,
    instructions: "",
  });
}

function removeItem(index) {
  form.value.items.splice(index, 1);
}

async function submitPrescription() {
  if (
    !form.value.appointment_id ||
    !form.value.diagnosis ||
    form.value.items.some((i) => !i.medicine_id)
  ) {
    error.value = "Please fill in all required fields.";
    return;
  }

  submitting.value = true;
  error.value = "";
  success.value = false;

  try {
    await api.post("/prescriptions", {
      appointment_id: form.value.appointment_id,
      doctor_id: form.value.doctor_id,
      patient_id: form.value.patient_id,
      diagnosis: form.value.diagnosis,
      notes: form.value.notes,
      items: form.value.items.map((i) => ({
        ...i,
        duration_days: parseInt(i.duration_days),
        quantity: parseInt(i.quantity),
      })),
    });
    success.value = true;
    form.value = {
      appointment_id: "",
      patient_id: "",
      doctor_id: "",
      diagnosis: "",
      notes: "",
      items: [
        {
          medicine_id: "",
          dosage: "",
          frequency: "",
          duration_days: 5,
          quantity: 1,
          instructions: "",
        },
      ],
    };
  } catch (err) {
    error.value = err.response?.data?.error || "Failed to create prescription.";
  } finally {
    submitting.value = false;
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
