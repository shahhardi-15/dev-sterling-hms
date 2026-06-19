<template>
  <div>
    <!-- Header -->
    <div class="mb-8">
      <h1 class="text-2xl sm:text-3xl font-bold text-gray-900">
        My Prescriptions
      </h1>
      <p class="text-gray-400 text-sm mt-1">
        View all your prescriptions and medicines
      </p>
    </div>

    <div class="space-y-4">
      <div
        v-for="prescription in prescriptions"
        :key="prescription.id"
        class="bg-white rounded-2xl border border-gray-100 shadow-sm p-6"
      >
        <!-- Header -->
        <div
          class="flex flex-col sm:flex-row sm:items-start justify-between gap-2 mb-4"
        >
          <div>
            <h2 class="text-base font-bold text-gray-900">
              {{ prescription.diagnosis }}
            </h2>
            <p class="text-sm text-gray-400 mt-0.5">
              Dr. {{ prescription.doctor_name }} &mdash;
              {{ formatDate(prescription.created_at) }}
            </p>
          </div>
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
        </div>

        <!-- Notes -->
        <p
          v-if="prescription.notes"
          class="text-sm text-gray-500 mb-4 bg-gray-50 rounded-xl px-4 py-2"
        >
          {{ prescription.notes }}
        </p>

        <!-- Medicines -->
        <div v-if="prescription.items && prescription.items.length">
          <p
            class="text-xs font-semibold text-gray-400 uppercase tracking-widest mb-2"
          >
            Medicines
          </p>
          <div class="space-y-2">
            <div
              v-for="item in prescription.items"
              :key="item.id"
              class="flex items-center justify-between rounded-xl px-4 py-3"
              style="background-color: #e8f0fe"
            >
              <div>
                <p class="text-sm font-medium text-gray-900">
                  {{ getMedicineName(item.medicine_id) }}
                </p>
                <p class="text-xs text-gray-500 mt-0.5">
                  {{ item.dosage || "" }}
                  <span v-if="item.frequency"> · {{ item.frequency }}</span>
                  <span v-if="item.duration_days">
                    · {{ item.duration_days }} days</span
                  >
                </p>
              </div>
              <span class="text-xs font-semibold" style="color: #1e3a5f"
                >Qty: {{ item.quantity }}</span
              >
            </div>
          </div>
        </div>
        <p v-else class="text-sm text-gray-400">No medicines prescribed</p>
      </div>

      <div
        v-if="prescriptions.length === 0"
        class="bg-white rounded-2xl border border-gray-100 shadow-sm px-6 py-12 text-center text-sm text-gray-400"
      >
        No prescriptions yet
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useAuthStore } from "../../stores/auth";
import api from "../../api/axios";

const auth = useAuthStore();
const prescriptions = ref([]);
const medicines = ref([]);

onMounted(async () => {
  try {
    const [patientsRes, medRes] = await Promise.all([
      api.get("/patients"),
      api.get("/medicines"),
    ]);
    medicines.value = medRes.data.data || [];
    const myPatient = patientsRes.data.data.find(
      (p) => p.email === auth.user.email,
    );
    if (myPatient) {
      const res = await api.get(`/patients/${myPatient.id}/prescriptions`);
      prescriptions.value = res.data.data || [];
    }
  } catch (err) {
    console.error(err);
  }
});

function getMedicineName(medicineId) {
  const med = medicines.value.find((m) => m.id === medicineId);
  return med ? med.name : medicineId;
}

function formatDate(dateStr) {
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}
</script>
