<template>
  <div>
    <div class="mb-8 flex items-center justify-between">
      <div>
        <h1 class="text-2xl sm:text-3xl font-bold text-gray-900">Billing</h1>
        <p class="text-gray-400 text-sm mt-1">
          Manage patient billing and payments
        </p>
      </div>
    </div>

    <!-- Stat Cards -->
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 sm:gap-5 mb-8">
      <div class="bg-white rounded-2xl p-5 border border-gray-100 shadow-sm">
        <p class="text-sm text-gray-400 mb-1">Total Bills</p>
        <p class="text-3xl font-bold text-gray-900">{{ bills.length }}</p>
      </div>
      <div class="bg-white rounded-2xl p-5 border border-gray-100 shadow-sm">
        <p class="text-sm text-gray-400 mb-1">Pending</p>
        <p class="text-3xl font-bold text-orange-500">
          {{ pendingBills.length }}
        </p>
      </div>
      <div class="bg-white rounded-2xl p-5 border border-gray-100 shadow-sm">
        <p class="text-sm text-gray-400 mb-1">Paid</p>
        <p class="text-3xl font-bold text-green-600">{{ paidBills.length }}</p>
      </div>
    </div>

    <!-- Bills Table -->
    <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-gray-100">
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Bill ID
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Patient
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Amount
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
              v-for="bill in bills"
              :key="bill.id"
              class="border-b border-gray-50 hover:bg-gray-50 transition"
            >
              <td class="px-6 py-4 text-sm font-mono text-gray-400">
                {{ bill.id.slice(0, 8) }}...
              </td>
              <td class="px-6 py-4 text-sm font-medium text-gray-900">
                {{ bill.patient_name || "—" }}
              </td>
              <td class="px-6 py-4 text-sm font-semibold text-gray-900">
                ₹{{ bill.amount }}
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ formatDate(bill.created_at) }}
              </td>
              <td class="px-6 py-4">
                <span
                  class="px-2.5 py-1 rounded-full text-xs font-semibold"
                  :class="
                    bill.paid
                      ? 'bg-green-50 text-green-600'
                      : 'bg-orange-50 text-orange-500'
                  "
                  >{{ bill.paid ? "Paid" : "Pending" }}</span
                >
              </td>
              <td class="px-6 py-4">
                <div class="flex gap-3">
                  <button
                    v-if="!bill.paid"
                    @click="markAsPaid(bill.id)"
                    class="text-xs text-white px-3 py-1.5 rounded-lg font-medium transition"
                    style="background-color: #1e3a5f"
                  >
                    Mark Paid
                  </button>
                  <button
                    @click="downloadReceipt(bill)"
                    class="text-xs text-blue-600 hover:underline font-medium"
                  >
                    Download Receipt
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="bills.length === 0">
              <td
                colspan="6"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                No bills found
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
import jsPDF from "jspdf";

const bills = ref([]);

const pendingBills = computed(() => bills.value.filter((b) => !b.paid));
const paidBills = computed(() => bills.value.filter((b) => b.paid));

onMounted(async () => {
  try {
    const res = await api.get("/billing");
    bills.value = res.data?.data || [];
  } catch (err) {
    console.error(err);
    bills.value = [];
  }
});

async function markAsPaid(id) {
  try {
    await api.patch(`/billing/${id}/pay`, { payment_method: "cash" });
    bills.value = bills.value.map((b) =>
      b.id === id ? { ...b, paid: true, status: "paid" } : b,
    );
  } catch (err) {
    console.error(err);
  }
}

function downloadReceipt(bill) {
  const doc = new jsPDF();

  doc.setFontSize(18);
  doc.setFont(undefined, "bold");
  doc.text("STERLING HOSPITAL MANAGEMENT SYSTEM", 20, 20);

  doc.setFontSize(12);
  doc.setFont(undefined, "normal");
  doc.text("Payment Receipt", 20, 30);
  doc.line(20, 34, 190, 34);

  doc.setFontSize(11);
  doc.text(`Bill ID: ${bill.id}`, 20, 45);
  doc.text(`Patient: ${bill.patient_name || "—"}`, 20, 53);
  doc.text(`Amount: Rs. ${bill.amount}`, 20, 61);
  doc.text(`Status: ${bill.status === "paid" ? "PAID" : "PENDING"}`, 20, 69);
  doc.text(`Date: ${formatDate(bill.created_at)}`, 20, 77);

  doc.line(20, 85, 190, 85);
  doc.setFontSize(10);
  doc.text("Thank you for choosing Sterling Hospital", 20, 95);

  doc.save(`receipt-${bill.id.slice(0, 8)}.pdf`);
}

function formatDate(dateStr) {
  if (!dateStr) return "—";
  return new Date(dateStr).toLocaleDateString("en-IN", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}
</script>
