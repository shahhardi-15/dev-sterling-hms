<template>
  <div>
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">My Prescriptions</h1>
      <p class="text-gray-400 text-sm mt-1">All prescriptions you've written</p>
    </div>

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
                Medicines
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
              v-for="p in prescriptions"
              :key="p.id"
              class="border-b border-gray-50 hover:bg-gray-50 transition"
            >
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                {{ p.patient_name || "—" }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ p.diagnosis }}</td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ formatDate(p.created_at) }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ (p.items || []).length }} item(s)
              </td>
              <td class="px-6 py-4">
                <span
                  class="px-2.5 py-1 rounded-full text-xs font-semibold"
                  :class="
                    p.dispensed
                      ? 'bg-green-50 text-green-600'
                      : 'bg-orange-50 text-orange-500'
                  "
                  >{{ p.dispensed ? "Dispensed" : "Pending" }}</span
                >
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
import { ref, onMounted } from "vue";
import { useAuthStore } from "../../stores/auth";
import api from "../../api/axios";

const auth = useAuthStore();
const prescriptions = ref([]);

onMounted(async () => {
  try {
    const doctorRes = await api.get(`/doctors/me/${auth.user.id}`);
    const doctorID = doctorRes.data?.data?.id;
    if (!doctorID) return;

    const patientsRes = await api.get("/patients");
    const allPrescriptions = [];

    for (const patient of patientsRes.data?.data || []) {
      const res = await api.get(`/patients/${patient.id}/prescriptions`);
      const patientPrescriptions = (res.data?.data || []).filter(
        (p) => p.doctor_id === doctorID,
      );
      allPrescriptions.push(
        ...patientPrescriptions.map((p) => ({
          ...p,
          patient_name: patient.full_name,
        })),
      );
    }

    prescriptions.value = allPrescriptions.sort(
      (a, b) => new Date(b.created_at) - new Date(a.created_at),
    );
  } catch (err) {
    console.error(err);
    prescriptions.value = [];
  }
});

function formatDate(dateStr) {
  if (!dateStr) return "—";
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}
</script>
